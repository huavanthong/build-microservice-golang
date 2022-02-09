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
> Is there any better way to send our data to the output stream without marshalling to a temporary object before we return it?

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
Summary all projects reference from other resource on website.

## GolangProject
Begin to deploy the actual golang project. This project will implement Golang following MVC model. 
We will find the interesting feature, and integrate into this project.

## RPC
Begin to deploy the actual golang project. This project will implement Golang following MVC model. 
We will find the interesting feature, and integrate into this project.
