help:
	@echo "Makefile help"

game-2048-go:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: game-2048-go
	docker build -t mingz2013/game-2048-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/game-2048-go

run:
	docker run -d --link redis-mq:redis-mq --name game-2048-go -it mingz2013/game-2048-go:latest


.PYONY: help, commit-docker, docker-image, game-2048-go, run

