# web-programming/HW1
![IMG_5979](https://user-images.githubusercontent.com/59199865/138659512-b6aead5f-2f6e-43d6-885c-66b9fc4174fe.JPG)
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
7. type your server ip with port 81 for example `http://0.0.0.0:81/` </br>
8. enjoy it!
