# Introduction
In this article, we will create a simple web service with gRPC in Go. We won’t use any third-party tools and only create a single endpoint to interact with the service to keep things simple.

# Reference
Refer: [here](https://sahansera.dev/building-grpc-server-go/)

# Install package
* To install protobuf. Refer: [here](https://github.com/huavanthong/microservice-golang/tree/master/currency#install-protos)
* To install grpcurl
```
go get github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
More details: [here](https://github.com/fullstorydev/grpcurl)
# Getting Started
To get all of books in shop
```
grpcurl -plaintext localhost:8080 Inventory/GetBookList
```
**Note:** 
* gRPC defaults to TLS for transport. 
* However, to keep things simple, I will be using the -plaintext flag with grpcurl so that we can see a human-readable response.