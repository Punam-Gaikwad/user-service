build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/Punam-Gaikwad/user-service \
      proto/user/user.proto