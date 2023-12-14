# 导航后台管理

## 如何使用
### 打包编译运行
```
git clone https://github.com/playonefor/FinderDaohang.git
make init module=FinderDaohang
make install
make serve
```
### docker
```
docker build -t finderdaohang:v1 .
docker run -d --name finder -p 9115:9115 finderdaohang:v1
```


## 访问导航页面
```
http://0.0.0.0:9115
```

## 访问后台管理
```
http://0.0.0.0:9115/admin
username: admin
password: admin
```
