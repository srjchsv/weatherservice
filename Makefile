run:
	@go run cmd/app/*.go
runenv:
	@eval $(cat ./env | xargs -0) go run cmd/app/*.go
	