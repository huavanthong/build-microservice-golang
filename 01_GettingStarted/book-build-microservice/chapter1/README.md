This chapter 1 will help you understand some lesson below.  

In this example, we are going to create an HTTP server with a single endpoint that returns static text represented by the JSON  
standard, this will introduce the basic functions of the HTTP server and handlers.  
We will then modify this endpoint to accept a request that is encoded in JSON and using the encoding/json package return a response to the client.  
We will also examine how the routing works by adding a second endpoint that returns a simple image.  
By the end of this chapter, you will have a fundamental grasp of the basic packages and how you can use them to quickly and
efficiently build a simple microservice.  

## Table of contents
* [Building a simple web server with net/http](#Build-web-server)
* [Reading and writing json](#technologies)
* [Routing in net/http](#Routing)
* [Context](#Context)
* [RPC in the Go standard library](#RPC)

## Build-web-server
This example is putting in basic_http_server.
To run this example:  
```
$go run ./basic_http_example.go
```

You should now see the application output:  
```
2022/02/08 23:08:54 Server starting on port 8080
```

To check this instance running on process.  
```
$ps -aux | grep 'go run
```
## Json
Thanks to the encoding /json [package](#https://pkg.go.dev/encoding/json), which is built into the standard library encoding and decoding JSON to and from Go
types is both fast and easy.
It implements the simplistic Marshal and Unmarshal functions; however, if we need them, the package also provides Encoder and Decoder types that allow us greater control when reading and writing

### Marshalling Go structs to JSON
Json in Golang is powerful, we need to remember some feature for coding:
* Do not ouput this field in struct.
* Do not output the field if the value is empty.
* convert output to a string and rename "id".

### Unmarshalling JSON to Go structs

## Routing
Summary all projects reference from other resource on website.

## GolangProject
Begin to deploy the actual golang project. This project will implement Golang following MVC model. 
We will find the interesting feature, and integrate into this project.

## RPC
Begin to deploy the actual golang project. This project will implement Golang following MVC model. 
We will find the interesting feature, and integrate into this project.
