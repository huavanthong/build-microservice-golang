# Chapter 4  
This chapter will help you understand about some patterns in Microservice? 
- It will help you get a overview about Microservice.
- Help you from basic design to the actual design in the reality.
- Beside that, we need to imagine your design patterns in Microservice to the real world, and thinking about how can it solve the problem?

# Introduction
I am not talking about software design patterns like factories or facades, but architectural designs like load balancing and service discovery.

# Table of contents
* [Event Processing Pattern](#event-processing-pattern)
* [Timeout Pattern](#timeout-pattern)
* [Back off Pattern](#back-off-pattern) 
* [Circuite Breaking Pattern](#circuite-breaking-pattern)
* [Heath checks Pattern](#heath-checks-pattern)
* [Throttling Pattern](#throttling-pattern)
* [Service Discovery Pattern](#service-discovery-pattern)
* [Load Balancing Pattern](#load-balancing-pattern)
* [Caching Pattern](#caching-pattern)

# Questions

## About Event Processing

## About Timeout Pattern
* [Why do we need Timeout Pattern in Microservice?](#timeout-pattern)
* [What is package for using timeout pattern? go-resiliency or ?](#package-for-time-out-pattern)
* [Could you use another package instead of go-resiliency](#)
## About Back off Pattern
* [What is Back off Pattern? When we need to use it in our system?](#back-off-pattern)
* [Do you understand about flooding the network?](#flooding-the-network)
* [The difference between flooding network and flooding server?](#flooding-the-server-with-requests)
* [Suppose you can a project, what steps you follow to implement backoff pattern for that project?](#design-for-back-off-pattern)
* [Beside using ConstantBackoff for Back off, do you know any else? What specification for this retrier package?](#specification-for-retrier-package)
* [How to run this project - backoff in this example](#to-run-back-off-pattern)
* [Could you implement backoff pattern on server project?](#implement-backoff-pattern-on-server)
* [Right now, we can brainstorm your mind to remember how many design did you learn before?](#implement-backoff-pattern-on-server)
## About Circuite Breaking Pattern
* [What is Circuit Pattern?](#circuite-breaking-pattern)
* [Why do we need Circuit Pattern in Microservice?](#circuite-breaking-pattern)
* [Do you understand about work flow on Circuit Pattern?](#work-flow-circuite-pattern)
* [Getting Started with Circuit Pattern in Golang?](#getting-started-with-circuit-pattern)
* [What is the Hystix library from Netflix?](#hystix-library-from-netflix)

## About Health checks Pattern
* [What is health check?](#heath-checks-pattern)
* [Do you know that what features we need to apply health check?](#usage-of-health-check)
* [Implement health check to measure time for any request](#getting-started-with-health-check)
## About Throttling Pattern
* [What is Throttling Pattern?](#throttling-pattern)

## About Service Discovery Pattern

## About Load Balancing Pattern

## About Caching Pattern

--------------------------------------------------------------------------------------------------------------------------------
## Event Processing Pattern


## Timeout Pattern
A timeout is an incredibly useful pattern while communicating with other services or data stores. The idea is that you set a limit on the response of a server and, if you do not receive a response in the given time, then you write a business logic to deal with this failure, such as retrying or sending a failure message back to the upstream service.  
More details explaination: [here](https://www.meisternote.com/app/note/0gdFcuDdHd3p/timeouts)
### Package for Time out pattern
To download package for timeout.
```
Clone github package:
    https://github.com/eapache/go-resiliency
    https://github.com/eapache/go-resiliency/tree/master/deadline
    
Download by Golang:
    go get github.com/eapache/go-resiliency/deadline
```

To know about version of go-resiliency
```
    https://pkg.go.dev/github.com/eapache/go-resiliency
```

To create go.mod for this project
```
module timeout

go 1.16

require github.com/eapache/go-resiliency v1.2.0
```

More details about project: [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter5/timeout)

### Design for Timeout Pattern
To create timeout for any service, we following steps below:
**Step 1:** Use go-resiliency package to create a instance deadline to set timeout
```
dl := deadline.New(1 * time.Second)
```
**Step 2:** Use instance deadline to run our service
```
err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction() <=============== This is our service
		return nil
	})
```
**Step 3:** Write a business logic to deal with this failure, such as retrying or sending a failure message back to the upstream service
```
	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
```
### To run timeout pattern
To run at slow case
```
go run .\main.go slow
```
Output
```
Loop:  0
Loop:  1
Loop:  2
Loop:  3
Loop:  4
exit status 2
```

To run at timeout case
```
go run .\main.go timeout
```
Output
```
Loop:  0
Timeout
``` 
### Context package
Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. 
More details: [here](https://pkg.go.dev/context)

### Implement timeout using Context

## Back off Pattern
Typically, once a connection has failed, you do not want to retry immediately to avoid flooding the network or the server with requests. To allow this, it's necessary to implement a back-off approach to your retry strategy. A back-off algorithm waits for a set period before retrying after the first failure, this then increments with subsequent failures up to a maximum duration
### Flooding the network 

More details: [here](https://www.meisternote.com/app/note/jx-Ok8sPiVr_/flooding-the-network)
### Flooding the server with requests
More details: [here](https://www.meisternote.com/app/note/wPg1NpQ4B-mP/flooding-the-server-and-request)

### Package for Back off Pattern
Back off is implemented in package go-resiliency package and the retrier package.  
To download it to your working project
```
    go get github.com/eapache/go-resiliency/retrier
```

To create go.mod for this project
```
module backoff

go 1.16

require github.com/eapache/go-resiliency v1.2.0
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter5/backoff)

### Design for Back off Pattern
To create a Back off Patter for your microservice, we follow steps:
**Step 1:** To create a new retrier, we use the New function which has the signature:
```
    r := retrier.New(retrier.ConstantBackoff(3, 1*time.Second), nil)
```
**Step 2:** Run our service with our backoff pattern.
```
    n := 0
    err := r.Run(func() error {
        fmt.Println("Attempt: ", n)
        n++
        return fmt.Errorf("Failed")
    })
```
### Specification for retrier package.
To create a new retrier, we use the New function which has the signature:
```
func New(backoff []time.Duration, class Classifier) *Retrier
```

The first parameter is an array of Duration. Rather than calculating this by hand, we can use the two built-in methods which
will generate this for us:
```
func ConstantBackoff(n int, amount time.Duration) []time.Duration
```

The ConstantBackoff function generates a simple back-off strategy of retrying n times and waiting for the given amount of
time between each retry:
```
func ExponentialBackoff(n int, initialAmount time.Duration) []time.Duration
```
The ExponentialBackoff function generates a simple back-off strategy of retrying n times doubling the time between each
retry.  

The second parameter is a Classifier. This allows us a nice amount of control over what error type is allowed to retry and
what will fail immediately.
```
type DefaultClassifier struct{}
```

The DefaultClassifier type is the simplest form: if there is no error returned then we succeed; if there is any error returned
then the retrier enters the retry state.
```
type BlacklistClassifier []error
```

The BlacklistClassifier type classifies errors based on a blacklist. If the error is in the given blacklist it immediately fails;
otherwise, it will retry.
```
type WhitelistClassifier []error
```
The WhitelistClassifier type is the opposite of the blacklist, and it will only retry when an error is in the given white list.
Any other errors will fail.

### To run Back off Pattern
To run this project
```
go run .\main.go
```
Output
```
Attempt:  0
Attempt:  1
Attempt:  2
Attempt:  3
Failed
```
### Implement Backoff pattern on Server
More details about source code: [here](https://github.com/huavanthong/build-microservice-golang/blob/feature/chapter5-BackOff/01_GettingStarted/book-build-microservice/chapter5/backoff/server.go)  
Right now, we will summary how many design exist in this server.
#### Implement function to get log

#### Implement function to validate a http request from client.

#### Using chain handler function in our server

#### Using back off pattern in our server.

### Using Marshalling and UnMarshalling for our server

## Circuite Breaking Pattern
Circuit breaking is all about failing fast, Michael Nygard in his book "Release It" says:
```
Circuit breakers are a way to automatically degrade functionality when the system is under stress.
```

### Why we need Circuite Breaking Pattern
In a simpel word, we can say that we use this pattern, to open connection again to avoid failing error related to timeout with a service which have a large time processing.
More details: [here](https://www.meisternote.com/app/note/DF_ZyKtkbCwT/circuit-breaking)
### Work flow Circuite Pattern
* **1.** Under normal operations, like a circuit breaker in your electricity switch box, the breaker is closed and **1.1** traffic flows normally.
* **1.2** However, once the pre-determined error threshold has been exceeded, the breaker enters the open state, and all requests **(include 2. -> 2.1, 3. ->3.2)** immediately fail without even being attempted. 
* After a period, a further request would be allowed and the circuit enters a halfopen state, in this state a failure immediately returns to the open state regardless of the errorThreshold.
* Once some requests have been processed without any error, then the circuit again returns to the closed state, and only if the number of failures
exceeded the error threshold would the circuit open again.
That gives us a little more context to why we need circuit breakers, but how can we implement them in Go?

![image](https://user-images.githubusercontent.com/50081052/159105466-b8903c3c-6ac9-4c1f-aa4b-4f0cd0f0dc8e.png)
### Getting Started with Circuit Pattern
Circuite Pattern can be found at **go-resilience package**. More details: [breaker](https://pkg.go.dev/github.com/eapache/go-resiliency/breaker) 


To create a circuit breaker 
```
func New(errorThreshold, successThreshold int, timeout time.Duration) *Breaker
```
We construct our circuit breaker with three parameters:
* The first errorThreshold, is the number of times a request can fail before the circuit opens
* The successThreshold, is the number of times that we need a successful request in the half-open state before we move back to open
* The timeout, is the time that the circuit will stay in the open state before changing to half-open

Design circuit pattern
```
	b := breaker.New(3, 1, 5*time.Second)

	for {
		result := b.Run(func() error {
			// Call some service
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Timeout")
		})

		switch result {
		case nil:
			// success!
		case breaker.ErrBreakerOpen:
			// our function wasn't run because the breaker was open
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
```

Could you explain this output:
```
Timeout
Timeout
Timeout
Breaker open
Breaker open
Breaker open
...
Breaker open
Breaker open
Timeout
Breaker open
Breaker open
```

More details: [here](https://github.com/huavanthong/build-microservice-golang/blob/master/01_GettingStarted/book-build-microservice/chapter5/circuit/main.go)

## Hystix library from Netflix
* One of the more modern implementations of circuit breaking and timeouts is the Hystix library from Netflix; Netflix is certainly renowned for producing some quality microservice architecture and the Hystrix client is something that has also been copied time and time again.  
* Hystrix is described as "a latency and fault tolerance library designed to isolate points of access to remote systems, services, and third-party libraries, stop cascading failure, and enable resilience in complex distributed systems where failure is inevitable."

### Hystrix from Github
More details explanation: [here](https://dzone.com/articles/go-microservices-part-11-hystrix-and-resilience)
> https://github.com/Netflix/Hystrix

## Health checks Pattern
* Health checks should be an essential part of your microservices setup. 
* Every service should expose a health check endpoint which can be accessed by the consul or another server monitor. Health checks are important as they allow the process responsible for running the application to restart or kill it when it starts to misbehave or fail. 
* Of course, you must be incredibly careful with this and not set this too aggressively.

### Usage of Health check
What you record in your health check is entirely your choice. However, I recommend you look at implementing these features:
* Data store connection status (general connection state, connection pool status)
* Current response time (rolling average)
* Current connections
* Bad requests (running average)

### Getting Started with Health check
We are defining two handlers one which deals with our main request at the path / and one used for checking the health at the path /health.
```
func main() {
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
```

Implement a handler function with a health check
```
func mainHandler(rw http.ResponseWriter, r *http.Request) {
	
	startTime := time.Now()

	if !isHealthy() {
		respondServiceUnhealthy(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Average request time: %f (ms)\n", ma.Value()/1000000)

	duration := time.Now().Sub(startTime)
	ma.Add(float64(duration))
}
```

Create a mutex Lock for setting global varaible.
```
func respondServiceUnhealthy(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)

	resetMutex.RLock()
	defer resetMutex.RUnlock()

	if !resetting {
		go sleepAndResetAverage()
	}
}
```

Sleep to wait timeout, and reset values
```
func sleepAndResetAverage() {

	resetMutex.Lock()
	resetting = true
	resetMutex.Unlock()

	time.Sleep(timeout)
	ma = ewma.NewMovingAverage()

	resetMutex.Lock()
	resetting = false
	resetMutex.Unlock()
}

```

## Throttling Pattern
Throttling is a pattern where you restrict the number of connections that a service can handle, returning an HTTP error code when this threshold has been exceeded. The full source code for this example can be found in the file throttling/limit_handler.go. The middleware pattern for Go is incredibly useful here: what we are going to do is to wrap the handler we would like to call, but before we call the handler itself, we are going to check to see if the server can honor the request. In this example, for simplicity, we are going only to limit the number of concurrent requests that the handler can serve, and we can do this with a simple buffered channel.
### Getting Started with Throttling 
Declare a LimitHandler handler. We have two private fields:
-  one holds the number of connections as a buffered channel.
-  the second is the handler we are going to call after we have checked that the
system is healthy

```
type LimitHandler struct {
	connections chan struct{}
	handler     http.Handler
}
```

To create an instance of this object we are going to use the **NewLimitHandler** function. This takes the parameters connection, which is the number of
connections we allow to process at any one time and the handler which would be called if successful:
```
func NewLimitHandler(connections int, next http.Handler) *LimitHandler {
	cons := make(chan struct{}, connections)
	for i := 0; i < connections; i++ {
		cons <- struct{}{}
	}

	return &LimitHandler{
		connections: cons,
		handler:     next,
	}
}
```

This is quite straightforward: we create a buffered channel with the size equal to the number of concurrent connections, and then we fill that ready for use:
```
func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r)
		l.connections <- struct{}{} // release the lock
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	}
}

Note: 
    If we look at the ServeHTTP method starting at line 29, we have a select statement. The beauty of channel is that we can write a statement like this: if we cannot retrieve an item from the channel then we should return a busy error message to the client.
```

### Testing 
Let's take a closer look at the flow through this test:
1. Block at line 109.2. Call handler.ServeHTTP twice concurrently.
3. One ServeHTTP method returns immediately with http.TooManyRequests and decrements the wait group.
4. Call cancel context allowing the one blocking ServeHTTP call to return and decrement the wait group.
5. Perform assertion.

### Running
To build this project.
```
go test
```

Output
```
PASS
ok      throttling      0.490s
```

## Service Discovery Pattern

## Load Balancing Pattern

## Caching Pattern