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
* [What is package for using timeout pattern?]
## About Back off Pattern

## About Circuite Breaking Pattern

## About Heath checks Pattern

## About Throttling Pattern

## About Service Discovery Pattern

## About Load Balancing Pattern

## About Caching Pattern
###############################################################################################################
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

## Back off Pattern

## Circuite Breaking Pattern

## Heath checks Pattern

## Throttling Pattern

## Service Discovery Pattern

## Load Balancing Pattern

## Caching Pattern