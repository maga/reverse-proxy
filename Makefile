run:
	env `cat .env` go run cmd/recipes/main.go

build:
	export GO111MODULE=on
	mkdir -p bin/recipes
	env GOOS=linux go build -ldflags="-s -w" -o bin/services/resipes cmd/recipes/main.go

docker-build:
	docker-compose run app make build

clean:
	rm -rf ./bin ./vendor

test:
	env `cat .env` go test -v internal/domains/*

