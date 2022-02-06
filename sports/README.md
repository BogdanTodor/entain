## Sports Service

### Instructions

#### Proto Generation
Execute the following command in the sport directory to ensure the service is up to date with the most recent changes:

```
go generate ./...
```

#### Building
To build and start the sport service execute the following command:

```
go build && ./sports
```

It should be noted that this will start the gRPC server on port 9000. If this port is already in use, you can configure a different port number in main.go.
