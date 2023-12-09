# RoyServer blog项目后台


## 项目运行
```text
// 同步数据库
go run main.go -db

go run main.go
```


## 项目部署
```text
// 交叉编译 window编译linux平台
set GOARCH=amd62

set GOOS=linux

go build -o main

set GOOS=windows

mysqldump -uroot -proot blogger > blogger.sql

docs
uploads
main
application.yaml
blogger.sql

```