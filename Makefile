help:
	echo "help"

game-2048-go:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

docker-image: game-2048-go
	docker build -t mingz2013/game-2048-go .


commit-docker:docker-image
	docker login
	docker push mingz2013/game-2048-go


run:
	docker run --net="host" -it mingz2013/game-2048-go


.PYONY: commit-docker, docker-image, game-2048-go, help

