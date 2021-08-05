# go bbs
go web api server for bbs.  
<img src="https://stickershop.line-scdn.net/stickershop/v1/product/1349132/LINEStorePC/main.png;compress=true" width="100"><img src="https://future-architect.github.io/images/20210121/Moby-logo.png" width="100">  

## Features
- Clean Architecture
- golang:1.15.13
- [Air](https://github.com/cosmtrek/air)
- gin
- gorm
- goose
- mysql:5.7
- phpmyadmin

## Quick Start
1. copy env  
$ `cd app/`  
$ `cp .env.example .env`  
$ `cd ..`  
2. change db name on .env and docker-compose.yml  
3. build docker image  
$ `docker-compose build`  
4. make container of web(go), mysql, phpmyadmin  
$ `docker-compose up -d`  
5. check server logs  
$ `make logs`  
6. enter the web container   
$ `make web`  
7. migration  
\# `goose up`

8. access web  
http://localhost:8080/  
access phpmyadmin  
http://localhost:8081/

## Tips
- You don't need to run `go run main.go` everytime you changed go file. [Air](https://github.com/cosmtrek/air) is running in go_web container
