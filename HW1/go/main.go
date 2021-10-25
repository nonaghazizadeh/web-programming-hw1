package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func redisHandler(action string, message string) string {
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
			return "something went wrong with database"
		}
		return encodedSHA
	case "Get":

	default:
	}
	return ""
	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
	// Output: key value
	// key2 does not exist
}

func main() {
	r := gin.Default()

	r.GET("/go/sha256/:req", func(c *gin.Context) {

	})

	r.POST("/go/sha256", func(c *gin.Context) {
		data := c.Query("data")
		if len([]rune(data)) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Your message lenght must be more than 8 characters!",
			})
		} else {
			mess := redisHandler("Set", data)
			if mess == "something went wrong with database" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": mess,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": mess,
				})
			}
		}
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
