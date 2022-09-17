APIKEY=

.ONESHELL: server run
.PHONY: server run

run:
	@go run cmd/app/*.go

server:
	@export APIKEY=$(APIKEY)
	@go run cmd/app/*.go

