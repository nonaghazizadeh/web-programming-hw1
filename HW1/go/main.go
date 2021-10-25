package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type postJSON struct {
	Data string
}

var ctx = context.Background()

func redisHandler(action string, message string, c *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:7001",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	switch action {

	case "Set":
		h := sha256.New()
		h.Write([]byte(message))
		encodedSHA := hex.EncodeToString(h.Sum(nil))
		err := rdb.Set(ctx, encodedSHA, message, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong with database",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": encodedSHA,
		})

	case "Get":
		data, err := rdb.Get(ctx, message).Result()
		if err == redis.Nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Data doesn't exist",
			})
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong with database",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": data,
			})
		}

	default:
	}
}

func main() {
	r := gin.Default()

	r.GET("/go/sha256/:data", func(c *gin.Context) {
		data := c.Param("data")
		redisHandler("Get", data, c)
	})

	r.POST("/go/sha256/:data", func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong with database",
			})
		}
		var data postJSON
		json.Unmarshal(jsonData, &data)
		fmt.Println(data.Data)
		if len([]rune(data.Data)) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Your message lenght must be more than 8 characters!",
			})
		} else {
			redisHandler("Set", data.Data, c)
		}
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
