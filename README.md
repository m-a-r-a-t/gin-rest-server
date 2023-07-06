# gin-rest-server

REST сервер на Go

## Сборка и запуск в Docker

Собираем образ

```
docker build -t gin-rest-server:latest .
```

Запуск

```
docker run  -p 8080:8080 gin-rest-server:latest
```

## Локальный запуск

```
go mod download
```

```
go run cmd/gin-rest-server/main.go
```

## Прочее

Коллекция запросов в Postman https://www.postman.com/planetary-eclipse-277425/workspace/public-ws/request/14409173-98dff8f3-14d3-4988-8f61-0d5a89408cbd
