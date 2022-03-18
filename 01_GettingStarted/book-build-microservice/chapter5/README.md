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
* [What is package for using timeout pattern? go-resiliency or ?](#package-for-time-out-pattern)
* [Could you use another package instead of go-resiliency](#)
* [What is the Hystix library from Netflix?]

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

### Getting Started with Circuit Pattern


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