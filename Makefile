build:
	protoc -I. \
	--go_out=plugins=grpc:. ./proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
	docker run -p 50053:50053 -e MICRO_SERVER_ADDRESS=:50053 -e MICRO_REGISTRY=mdns user-service