FROM golang:1.12.2 as builder

WORKDIR /go/src/github.com/threehook/shippy/vessel-service

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the go mod download
# command will _only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

# Here we copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/threehook/shippy/vessel-service/vessel-service .

CMD ["./vessel-service"]
