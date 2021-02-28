.PHONY: build
build:
	yarn --cwd ./web build
	go build -v ./cmd/omo

.PHONY: build start
start: build
	./omo

.PHONY: makemig
makemig:
	migrate create -ext sql -dir ./db/migrations -seq $(seq)

.PHONY: migup
migup:
	migrate -database postgresql://postgres:postgres@localhost:5432/omo_dev -path ./db/migrations/ up

.DEFAULT_GOAL := build