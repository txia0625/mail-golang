# Mail App
This is a number of small, self-contained, loosely coupled microservices that will will communicate with one another and a simple front-end application with a REST API, with RPC, over gRPC, and by sending and consuming messages using AMQP, the Advanced Message Queuing Protocol. The microservices we build will include the following functionality:

- A Front End service, that just displays web pages;

- An Authentication service, with a Postgres database;

- A Logging service, with a MongoDB database;

- A Listener service, which receives messages from RabbitMQ and acts upon them;

- A Broker service, which is an optional single point of entry into the microservice cluster;

- A Mail service, which takes a JSON payload, converts into a formatted email, and send it out.

All of these services will be written in Go, commonly referred to as Golang, a language which is particularly well suited to building distributed web applications.

## Time Line
1/6/2024 Created Project

1/7/2024 Implemented the Broker Service 