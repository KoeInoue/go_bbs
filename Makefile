all: test build
setup:
	go get
test:
	go test ./...
build:
	go build -o app
up:
	docker-compose up -d
down:
	docker-compose down --remove-orphans
web: 
	docker container exec -it go_server ash
db: 
	docker container exec -it mariadb bash
logs:
	docker-compose logs -f web
bindata:
	docker container exec go_server go-bindata -debug -o ./bindata.go assets/...