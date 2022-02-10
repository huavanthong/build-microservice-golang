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

## Questions
* [Why we need to use Marshalling Go structs for json?](#Marshalling-Go-structs-to-JSON)
* [How to use Marshalling for response?](#Demo-Marshalling)
* [what if we need to read input before returning the output?](#Unmarshalling-JSON-to-Go-structs)
* [What is format data of HTML Request](https://github.com/huavanthong/MasterGolang/blob/main/01_GettingStarted/book-go-web-application/Chapter_4_Processing_Requests/README.md#HTML-Form)
* [How to decode JSON request from client?](#Demo-Unmarshalling)
* [Suppose you have a lot of images-css-js, how do you access it on golang?](#Static-file-handler)
* [How to design a chain handler? Think about purpose for this design](#Creating-handlers)
* [We get any trouble with HTTP Request? Why we need Context?](#Context)
* [What is RPC?](#RPC)
* [What is the default protocol in your DefaultServeMux ? Could you choose another protocol?](#Protocol)
* [Demo a example about RPC to understand the work flow](#Simple-RPC-example)
* [What problem if multiple client access to server?](#Server)
* [How client make message to server without use HTTP protocol (meaning that don't use Brower)?](#Client)


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
Thanks to the encoding /json [package](https://pkg.go.dev/encoding/json), which is built into the standard library encoding and decoding JSON to and from Go types is both fast and easy.
It implements the simplistic Marshal and Unmarshal functions; however, if we need them, the package also provides Encoder and Decoder types that allow us greater control when reading and writing

### Marshalling-Go-structs-to-JSON 
Json in Golang is powerful, we need to remember some feature for coding:
* Change the output field to be "message"
```
    Message string `json:"message"`
```
* Do not ouput this field in struct.
```
    Author string `json:"-"`
```
* Do not output the field if the value is empty.
```
    Date string `json:",omitempty"`
```
* convert output to a string and rename "id".
```
    Id int `json:"id, string"`
```
#### Demo-Marshalling
Using Marshalling
```
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	// Create message response following member in struct.
	response := helloWorldResponse{Message: "HelloWorld"}

	// Convert to json by using Marshal
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}
	// Write(p []byte) (n int, err error) only accept bytes
	// So we use Fprint to convert json to bytes.
	fmt.Fprint(w, string(data))
}
```
##### Is there any better way to send our data to the output stream without marshalling to a temporary object before we return it?  
The encoding/json package has a function called NewEncoder this returns us an Encoder object that can be used to write JSON straight to an open writer and guess what?  
Using Encode  
```
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Create message response following member in struct.
	response := helloWorldResponse{Message: "HelloWorld"}
	// Create object encoder.
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
}
```
### Unmarshalling-JSON-to-Go-structs
> func Unmarshal(data []byte, v interface{}) error

This function will allocate maps, slices, and pointers as required  

#### HTML-request-format
More detail at [here](https://github.com/huavanthong/MasterGolang/blob/main/01_GettingStarted/book-go-web-application/Chapter_4_Processing_Requests/README.md#HTML-Form)
* io.ReadCloser as a stream and does not return a []byte or a string
```
type Requests struct {
    ...
    // Method specifies the HTTP method (GET, POST, PUT, etc.).
    Method string
    // Header contains the request header fields received by the server. The type Header is a link to map[string] []string.
    Header Header
    // Body is the request's body.
    Body io.ReadCloser
    ...
}
```

#### Demo-Unmarshalling
```
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	// Body request as a stream.
	//  If we need the data contained in the body, we can simply read it into a byte array
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Convert json body request to local request.
	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// Take info request from client, and response to client.
	response := helloWorldResponse{Message: "Hello " + request.Name}

	// Create object encoder
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
}
```

## Routing
Even a simple microservice will need the capability to route requests to different handlers dependent on the requested path or
method.
- In Go this is handled by the DefaultServeMux method which is an instance of ServerMux.
- When we call the http.HandleFunc("/helloworld", helloWorldHandler) package function we are actually just indirectly calling http.DefaultServerMux.HandleFunc(…).

There are two functions to adding handlers to a ServerMux handler:
- Function handler
> func HandlerFunc(pattern string, handler func(ResponseWriter, *Request))
- Handler
> func Handle(pattern string, handler Handler)

**More details:** [here](https://www.meisternote.com/app/note/tIRDBorhiJSC/3-3-handlers-and-handler-functions)

### Paths
ServeMux is responsible for routing inbound requests to the registered handlers.  
Refer: Architecture for ServeMux in Go-web-application [book](https://www.meisternote.com/app/note/B6NG-U69TSGK/3-2-serving-go).
```
http.Handle("/images/", newFooHandler())
http.Handle("/images/persian/", newBarHandler())
http.Handle("/images", newBuzzHandler())
```
### Convenience handlers
```
?
```
### FileServer
To map the contents of the file system path ./images to the server route /images, Dir implements a file system which is restricted to a specific directory tree, the FileServer method uses this to be able to serve the assets.
```
http.Handle("/images", http.FileServer(http.Dir("./images")))
```

### NotFoundHandler
The NotFoundHandler function returns a simple request handler that replies to each request with a 404 page not found reply:
```
func NotFoundHandler() Handler
```

### RedirectHandler
To redirect to another handler
```
func RedirectHandler(url string, code int) Handler
```

### StripPrefix
The StripPrefix function returns a handler that serves HTTP requests by removing the given prefix from the request URL's path and then invoking h handler. If a path does not exist, then StripPrefix will reply with an HTTP 404 not found error:
```
func StripPrefix(prefix string, h Handler) Handler
```

### TimeoutHandler
The TimeoutHandler function returns a Handler interface that runs h with the given time limit. When we investigate common patterns in Chapter 6, Microservice Frameworks, we will see just how useful this can be for avoiding cascading failures in your service:
```
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

### Static-file-handler
Try to run **reading_writing_json_5** to see problem:
```
> http://localhost:8080/cat/cat.jpg.
```
Now we run **reading_writing_json_6** to see solution:
```
> http://localhost:8080/cat/cat.jpg.
```
Success due to some code below:
```
cathandler := http.FileServer(http.Dir("./images"))
http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))
```
### Creating-handlers
How to create handler. More detail refer: [here](https://github.com/huavanthong/MasterGolang/tree/main/01_GettingStarted/book-go-web-application/Chapter_3_Handling_Requests/handler)
Please distinct the handler and function handler.
* How to implement handler by pointer?
* How to implement handler by value?
* How to use chain handler?
Refer, and read carefully for more details: **reading_writing_json_7**

## Context
The problem with the previous pattern is that there is no way that you can pass the validated request from one handler to the next without breaking the http.Handler interface.  
This is a reason, Golang provide Context for our problem.  
The Context type implements a safe method for accessing request-scoped data that is safe to use simultaneously by multiple Go routines.  
### Backgroud
The Background method returns an empty context that has no values; it is typically used by the main function and as the toplevel Context.
```
func Background() Context
```

### WithCancel
The WithCancel method returns a copy of the parent context with a cancel function, calling the cancel function releases resources associated with the context and should be called as soon as operations running in the Context type are complete:  
```
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

### WithDeadline
The WithDeadline method returns a copy of the parent context that expires after the current time is greater than deadline. At this point, the context's Done channel is closed and the resources associated are released. It also passes back a CancelFunc method that allows manual cancellation of the context
```
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

### WithTimeout
The WithTimeout method is similar to WithDeadline except you pass it a duration for which the Context type should exist.  
Once this duration has elapsed, the Done channel is closed and the resources associated with the context are released:
```
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

### WithValue
The WithValue method returns a copy of the parent Context in which the val value is associated with the key.  
The Context values are perfect to be used for request-scoped data:
```
func WithValue(parent Context, key interface{}, val interface{}) Context
```

### Using-contexts
Refer: reading_writing_json_8

## RPC
Remote Procedure Call(RPC) in Operating System is a powerful technique for constructing distributed, client-server based applications.  
More details: [here](https://www.geeksforgeeks.org/remote-procedure-call-rpc-in-operating-system/)  

### Protocol
Firstly, we need to know about architecture of protocol: [here](#https://www.digitalocean.com/community/tutorials/http-1-1-vs-http-2-what-s-the-difference)  
And you know that the data will transfer to internet through 4 layer:  
* Application Layer (HTTP)
* Transport Layer (TCP)
* Network Layer (IP)
* Data Link Layer
Until the current ponint, we don't know that what protocol does DefaultServeMux use?  
However, we know DefaultServeMux use HTTP protocol.  
  
When you use RPC standard:
* you can select your protocol such as: tcp, tcp4, tcp6, unix, or unixpacket.
* you also using a given protocol and binding it to IP as same as DefaultServeMux [demo](#https://github.com/huavanthong/build-microservice-golang/blob/feature/chapter1/01_GettingStarted/book-build-microservice/chapter1/basic_http_server/basic_http_server.go)

### Simple-RPC-example
#### Server
To register a handler into server in RPC API.
```
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
```

To make server listen client.
```
l, err := net.Listen("("tcp",", fmt.Sprintf(":%(":%v",", port))
```

To accept an connection between client and server, and block to wait client complete
```
	for {
		conn, _ := l.Accept()
			go rpc.ServeConn(conn)
		}
	}
```

#### Client
How client can connect to server without HTTP protocol?
```
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
```

to make an request from Client
```
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
```

### RPC over HTTP
As you know, Simple-RPC-example is a example about communication between client and server without HTTP Protocol.
Right now, how we can implement application using HTTP by RPC.

#### Server
To make RPC over HTTP
```
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()
```

And then, we need HTTP serve our server
```
	l, err := net.Listen("("tcp",", fmt.Sprintf(":%(":%v",", port))
	http.Serve(l, nil)
```

#### Client
To make connection through HTTP.
```
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
```

### JSON-RPC over HTTP
Have you ever put a question that how we can communicate by JSON?  

#### Server
To create json on handler
```
	serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
```

#### Client
To create an request by json
```
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
```
