# Introduction
This tutorial will help you implement a simple RPC with Hello World message.  
To understand rpc: [here](https://www.meisternote.com/app/note/pOnavUy3wJ8e/simple-rpc-example)
## Knowledge
### Server
To implement a handler for RPC. Handler must have 2 arguement to apdapt rpc package.
* args *contract.HelloWorldRequest: For message request
* reply *contract.HelloWorldResponseL For message response
```go 
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
```
### Client
If a client want to perform a request to server, they must follow rules to reply messsage
* client *rpc.Client: input parameter is a rpc.Client
* Call(): must specific handler message "HelloWorldHandler.HelloWorld"
```go
func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	// After connect to server successfully, tức có nghĩa, khi mà ta sử dụng client.Call(), thì nó đã auto chạy được Dial() hay gì rồi?
	// Sau đó, ta có thể make a request đến Server bằng việc input các parameter vào client.Call().
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
```
# Getting Started
To build rpc example 
```go
go build
```

To run rpc example
```
./rpc
```

Output
```
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
Hello World
```
