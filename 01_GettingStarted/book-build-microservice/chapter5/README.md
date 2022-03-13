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
* [What is Back off Pattern? When we need to use it in our system?](#back-off-pattern)
* [Do you understand about flooding the network?](#flooding-the-network)
* [The difference between flooding network and flooding server?](#flooding-the-server-with-requests)
* [Suppose you can a project, what steps you follow to implement backoff pattern for that project?](#design-for-back-off-pattern)
* [Beside using ConstantBackoff for Back off, do you know any else? What specification for this retrier package?](#specification-for-retrier-package)
* [How to run this project - backoff in this example](#to-run-back-off-pattern)
* [Could you implement backoff pattern on server project?](#implement-backoff-pattern-on-server)
## About Circuite Breaking Pattern

## About Heath checks Pattern

## About Throttling Pattern

## About Service Discovery Pattern

## About Load Balancing Pattern

## About Caching Pattern
###############################################################################################################
## Event Processing Pattern

## Timeout Pattern

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
More details: [here](#)
## Circuite Breaking Pattern

## Heath checks Pattern

## Throttling Pattern

## Service Discovery Pattern

## Load Balancing Pattern

## Caching Pattern