build:
	protoc -I. --micro_out=${GOPATH}/src/github.com/threehook/shippy/vessel-service --go_out=${GOPATH}/src/github.com/threehook/shippy/vessel-service proto/vessel/vessel.proto
	docker build -t vessel-service .

all: build

run: all
	docker-compose up
