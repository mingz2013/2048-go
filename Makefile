help:
	@echo "Makefile help"


clean:
	rm game-2048-go -f


game-2048-go: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: game-2048-go
	docker build -t mingz2013/game-2048-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/game-2048-go


remove-container:
	docker rm game-2048-go


run: remove-container
	docker run -d --link redis-mq:redis-mq --name game-2048-go -it mingz2013/game-2048-go:latest


logs:
	docker logs game-2048-go


.PYONY: help, commit-docker, docker-image, game-2048-go, run, remove-container, logs

