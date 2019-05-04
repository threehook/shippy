build:
	protoc -I. --micro_out=${GOPATH}/src/github.com/threehook/shippy/consignment-service --go_out=${GOPATH}/src/github.com/threehook/shippy/consignment-service proto/consignment/consignment.proto
	docker build -t consignment-service .

run:
	docker run -p 5051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
