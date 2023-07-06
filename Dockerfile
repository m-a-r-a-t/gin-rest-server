FROM golang:1.20 AS build-stage

WORKDIR /app
COPY . .

RUN go mod download

WORKDIR /app/cmd/gin-rest-server

RUN CGO_ENABLED=0 GOOS=linux go build -o gin-rest-server

FROM alpine:latest AS build-release-stage
WORKDIR /

COPY --from=build-stage /app/cmd/gin-rest-server/gin-rest-server /app/config/config.yaml .
COPY --from=build-stage /app/static/ ./static

ENV CONFIG_PATH=/config.yaml
ENV GIN_MODE=release

EXPOSE 8080


CMD ["./gin-rest-server"]