This is helloworld example for testing grpc backward compatibility in golang. 

Changes:

- response changed to stream
- response message type changed to oneof 

Compile the protocol buffer definition

```
protoc --go_out=plugins=grpc:. proto/v1/helloworld.proto
protoc --go_out=plugins=grpc:. proto/v2/helloworld.proto
```

Run servers

```
go run server/v1/main.go
go run server/v2/main.go
```

Run client
```
go run client/main.go
```


Client output

```
Process client v1 & server v1:
Greeting: Hello world
Process client v1 & server v2:
Greeting: Hello world
Process client v2 & server v1:
Greeting: Hello world
Process client v2 & server v2:
Description: Example description
Greeting: Hello world
```
