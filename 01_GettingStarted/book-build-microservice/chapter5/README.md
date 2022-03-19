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

## About Heath checks Pattern

## About Throttling Pattern

## About Service Discovery Pattern

## About Load Balancing Pattern
* [What is load balancing?](#load-balancing-pattern)
* [Do you remember load balancing in ZeroMQ?](https://zguide.zeromq.org/docs/chapter4/)
* [Could you compare the distinct load balancing between ZeroMQ and Golang?](https://www.meisternote.com/app/note/p14YA6V5DCNN/distinct-to-zeromq)
## About Caching Pattern
###############################################################################################################
## Event Processing Pattern

## Timeout Pattern

## Back off Pattern

## Circuite Breaking Pattern

## Heath checks Pattern

## Throttling Pattern

## Service Discovery Pattern

## Load Balancing Pattern

### Getting Started with Load Balancing
I have created a simple implementation of a load balancer. We create it by calling **NewLoadBalancer** which has the following signature:
- a strategy, an interface that contains the selection logic for the endpoints
- and a list of endpoints.
```
func NewLoadBalancer(strategy Strategy, endpoints []url.URL) *LoadBalancer
```

To be able to implement multiple strategies for the load balancer, such as round-robin, random, or more sophisticated strategies
like distributed statistics, across multiple instances you can define your own strategy which has the following interface:
```
// Strategy is an interface to be implemented by loadbalancing
// strategies like round robin or random.
type Strategy interface {
    NextEndpoint() url.URL
    SetEndpoints([]url.URL)
}

NextEndpoint() url.URL
```
This is the method which will return a particular endpoint for the strategy. It is not called directly, but it is called internally by
the **LoadBalancer package** when you call the GetEndpoint method. This has to be a public method to allow for strategies to be
included in packages outside of the LoadBalancer package:


```
SetEndpoints([]url.URL)
```
This method will update the Strategy type with a list of the currently available endpoints. Again, this is not called directly but
is called internally by the LoadBalancer package when you call the UpdateEndpoints method.

To use the LoadBalancer package, you just initialize it with your chosen strategy and a list of endpoints, then by calling
GetEndpoint, you will receive the next endpoint in the list:
```
func main() {
	endpoints := []url.URL{
		url.URL{Host: "www.google.com"},
		url.URL{Host: "www.google.co.uk"},
	}

	lb := NewLoadBalancer(&RandomStrategy{}, endpoints)

	fmt.Println(lb.GetEndpoint())
}

```
### Run it
To run it
```
go run main.go
```

Could you explain this output
```
{   www.google.com   false   }
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter5/loadbalancing)

## Caching Pattern