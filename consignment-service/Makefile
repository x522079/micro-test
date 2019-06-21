build:
		protoc -I. --go_out=plugins=grpc:/Users/xin/data1/projects/go/shippy/consignment-service proto/consignment/consignment.proto
		GOOS=linux GOARCH=amd64 ENABLECGO=false go build
		docker build -t consignment-service .
run:
		docker run --name consignment-service -p 50051:50051 consignment-service