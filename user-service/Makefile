build:
	protoc -I. --micro_out=${GOPATH}/src/github.com/threehook/shippy/user-service --go_out=${GOPATH}/src/github.com/threehook/shippy/user-service proto/user/user.proto
	docker build -t user-service .

run:
	docker run -p 5053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service
