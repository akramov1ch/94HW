run:
    go run main.go

docker-build:
    docker build -t rest-api-metrics .

docker-run:
    docker-compose up --build
