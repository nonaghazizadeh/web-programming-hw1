# web-programming/HW1
![IMG_5983](https://user-images.githubusercontent.com/59199865/138674498-471b4d16-9e3d-47cc-98e9-43576fb9df7e.JPG)
## How to start
1. install nginx on your ubuntu server with the following command: </br>
```
sudo apt update
sudo apt install nginx
```
2. remove default file in `/etc/nginx/sites-enabled` with the following command: </br>
```
sudo rm /etc/nginx/sites-enabled/default
```
3. create file in `/etc/nginx/sites-enabled` and name it 'form' with the following command:</br>
```
sudo "${EDITOR:-vi}" form
```
4. write nginx.conf content that is in `/nginx/nginx.conf` directory, in form file and type `:wq` to write and quit  </br> 
5. go to `/var/www` directory with the following command: </br>
```
cd /var/www
```
6. clone the project with the following command: </br>
```
git clone https://github.com/nonaghazizadeh/web-programming.git
```
7. for other configurations do instructions in [go readme](https://github.com/nonaghazizadeh/web-programming/blob/master/HW1/go/README.md) and [node readme](https://github.com/nonaghazizadeh/web-programming/blob/master/HW1/node/README.md)
7. type your server ip with port 81 for example `http://0.0.0.0:81/` </br>
8. enjoy it!
