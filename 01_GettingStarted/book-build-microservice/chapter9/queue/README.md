# Introduction
This tutorial will help you understand how do you implement a queue using Redis with Docker.
### Redis open source.
Redis is an open source (BSD licensed), in-memory data structure store, used as a database, cache, and message broker. This example use Redis for storing the messages with some benefits:
* For the fast data store.
* Leverage a cloud provides queue rather managing our infrastucture.
* However, even if we are using cloud providers queue the pattern we are about to look at is easily replaceable with a different data store client.

# Table of Contents
1. Think about message and feature for using a Redis queue. [here](#design-message-for-queue)
2. With Redis package, how to use implement a function to use that package. [here](#design-to-access-redis-software)
3. How do you initialize a queue using Redis. [here](#initialize-a-queue-with-redis)
4. How do you implement a server to receive request and write data into the Redis queue. [here](#implement-server-for-writing-to-the-queue)
5. How do you implement a service to read data in the Redis queue. [here](#implement-service-to-read-the-queue)


### Design message for queue
Design our message and some actions on messages.
```go
// Message represents messages stored on the queue
type Message struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Payload string `json:"payload"`
}
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/blob/master/01_GettingStarted/book-build-microservice/chapter9/queue/queue.go)

### Design to function access Redis software
When you have messages, think about actions to handle our message. And we will define in the same location with message type.
```go
// Queue defines the interface for a message queue
type Queue interface {
	Add(messageName string, payload []byte) error
	AddMessage(message Message) error
	StartConsuming(size int, pollInterval time.Duration, callback func(Message) error)
}
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/blob/master/01_GettingStarted/book-build-microservice/chapter9/queue/queue.go)

### Initialize a queue with Redis
To init a queue with Redis, we need to embedded queue into Redis structure, and implement init() fucntion for it.
```go
// RedisQueue implements the Queue interface for a Redis based message queue
type RedisQueue struct {
	Queue    rmq.Queue
	name     string
	callback func(Message) error
}

var serialNumberLimit *big.Int

// init() initialize our Redis
func init() {
	serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)
}
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/blob/master/01_GettingStarted/book-build-microservice/chapter9/queue/redis_queue.go)

### Implement server for writing to the queue
Right now, we can create a queue Redis, and use it at anywhere. We need to think about how an server receive request and write message into queue. 
1. **Step 1:** Init a queue from Redis.
```go
func main() {
	q, err := queue.NewRedisQueue("redis:6379", "test_queue")
	if err != nil {
		log.Fatal(err)
	}

    ........
}
```
2. **Step 2:** Implement handler for our server
```go
func main() {

    ..........

    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		err := q.Add("new.product", data)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter9/queue/writer)
### Implement service to read the queue
Implement service to read message.
```go
func main() {
	log.Println("Starting worker")

    // Create a new Redis Queue at the same port with writer handler.
	q, err := queue.NewRedisQueue("redis:6379", "test_queue")
	if err != nil {
		log.Fatal(err)
	}

	q.StartConsuming(10, 100*time.Millisecond, func(message queue.Message) error {
		log.Printf("Received message: %v, %v, %v\n", message.ID, message.Name, message.Payload)

		return nil // successfully processed message
	})

	runtime.Goexit() // avoid main from exiting untill all other go routines have completed
}
```
More details: [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter9/queue/reader)