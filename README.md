# mev-boost-monitoring

It's a middleware between execution provider and mev-boost client.

## How to use the template

1.  Clone repository
2.  cd root repository
3.  make tools
4.  docker-composer up -d
5.  make migrate
6.  make build
7.  Run service ./bin/service

## How to create migrations?

./bin/migrate create -ext=sql -dir=db/migrations <your table name>

## Dependencies
1. go GRPC. You have to install protobuf.
   1. Mac OS - brew install protobuf
   2. Ubuntu - sudo apt install protobuf-compiler
2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest