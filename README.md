# wlog
This is a simple Go program that logs incoming HTTP requests and sends a response with the request body.

The program will start an HTTP server listening on port 8108.

```go
go run main.go
```

The docker run command:

```shell
docker run -it --rm  -p 8108:8108 zhaojiew/wlog:latest
```

Send a request to the server using a tool like cURL or a web browser. For example:

```shell
curl -d '@log.txt' localhost:8108
```
The program will log the request body and send the same body as the response.

Check the console output and the response received from the server to see the logged request body.