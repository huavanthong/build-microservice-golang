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

## About Health checks Pattern
* [What is health check?](#heath-checks-pattern)
* [Do you know that what features we need to apply health check?](#usage-of-health-check)
* [Implement health check to measure time for any request](#getting-started-with-health-check)
## About Throttling Pattern

## About Service Discovery Pattern

## About Load Balancing Pattern

## About Caching Pattern
###############################################################################################################
## Event Processing Pattern

## Timeout Pattern

## Back off Pattern

## Circuite Breaking Pattern

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

## Service Discovery Pattern

## Load Balancing Pattern

## Caching Pattern