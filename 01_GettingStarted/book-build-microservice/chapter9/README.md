# Introduction
This tutorial will help you understand Event-Driven Architecture.

# Table of Contents
1. Firstly, we need to understand the difference between synchronous and asynchronous processing. [here](#difference-between-synchronous-and-asynchronous-processing)
2. Secondly, we need to understand how many types of asynchronous messages? [here](#types-of-asynchronous-messages)
3. Thirdly, implement pull pattern in asynchronous. [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter9/queue)
4. Implement push pattern in asynchronous. [here](https://github.com/huavanthong/build-microservice-golang/tree/master/01_GettingStarted/book-build-microservice/chapter9/push)


### Synchronous
* With synchronous processing, all the communication to a downstream application happens in the process. 
* A request is sent, and you wait for a reply using the same network connection and not using any callbacks.
* Synchronous processing is the simplest method of communication; while you are waiting for an answer the downstream service is processing the request. 
* You have to manage the retry logic yourself, and it is typically best used only when you need an immediate reply. 
Let's take a look at the following diagram that depicts synchronous processing

![ảnh](https://user-images.githubusercontent.com/50081052/169202069-d5694256-34bb-46db-bc18-d6414568d28b.png)

### Asynchronous
* With asynchronous processing, all the communication to the downstream application happens out of process leveraging a queue or a message broker as an intermediary. 
* Rather than communicating directly with the downstream service, messages dispatch to a queue such as **AWS SQS/SNS**, **Google Cloud Pub/Sub**, or **NATS.io**. 
* Because there is no processing performed at this layer the only delay is the time it takes to deliver the message, which is very fast, also due to the design of these systems, acceptance, or not of a message is the only situation you must implement. 
* Retry and connection handling logic is delegated to either the message broker or the downstream system as it is the storage of messages for archive or replay:

![ảnh](https://user-images.githubusercontent.com/50081052/169202244-c8efec5e-64a2-4c23-8b09-9466f49d752b.png)
### Difference between synchronous and asynchronous processing
If there is a choice between processing a message synchronously or asynchronously, then I would always choose synchronous
as it always makes the application simpler with fewer components parts, the code is easier to understand, tests easier to write,
and the system easier to debug.

![ảnh](https://user-images.githubusercontent.com/50081052/169202289-977e0305-bb79-4481-b15a-bca279c4ac2d.png)


## Types of asynchronous messages
Asynchronous processing often comes in two different forms, such as push and pull. The strategy that you implement is dependent upon your requirements, and often a single system implements both patterns. Let's take a look at the two different approaches.'

