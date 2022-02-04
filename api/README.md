## API Service

### Instructions

#### Proto Generation
Execute the following command in the api directory to ensure the service is up to date with the most recent changes:

```
go generate ./...
```

#### Building
To build and start the api service execute the following command:

```
go build && ./api
```

It should be noted that this will start the API server on port 8000. If this port is already in use, you can configure a different port number in main.go.
