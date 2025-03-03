.PHONY: build
build:
	docker build . -t go-cli

.PHONY: up
up:
	docker run -v ./:/app --name go-cli -itd go-cli

.PHONY: down
down:
	docker stop go-cli
	docker rm go-cli

.PHONY: exec
exec:
	docker exec go-cli go run main.go $(ARGS)
