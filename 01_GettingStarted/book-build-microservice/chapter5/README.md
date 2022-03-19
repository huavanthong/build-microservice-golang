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

## About Back off Pattern

## About Circuite Breaking Pattern
* [What is Circuit Pattern?](#circuite-breaking-pattern)
* [Why do we need Circuit Pattern in Microservice?](#circuite-breaking-pattern)
* [Do you understand about work flow on Circuit Pattern?](#work-flow-circuite-pattern)
* [Getting Started with Circuit Pattern in Golang?](#getting-started-with-circuit-pattern)
* [What is the Hystix library from Netflix?](#hystix-library-from-netflix)

## About Heath checks Pattern

## About Throttling Pattern

## About Service Discovery Pattern

## About Load Balancing Pattern

## About Caching Pattern

--------------------------------------------------------------------------------------------------------------------------------
## Event Processing Pattern


## Timeout Pattern

## Back off Pattern

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
> https://github.com/Netflix/Hystrix

## Heath checks Pattern

## Throttling Pattern

## Service Discovery Pattern

## Load Balancing Pattern

## Caching Pattern