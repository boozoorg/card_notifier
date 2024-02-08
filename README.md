Notification service
====================

## Description
This service receives, store card activities and send a notify message to client (terminal), this service can receive a single message an array of messages. Those messages will be sent to client ones with those user session id.


## Develop
### Run command
```go
  go run cmd/main.go
```

### Build linux
```go
  env GOOS=linux go build cmd/main.go
```

### Build windows
```go
  env GOOS=windows go build cmd/main.go
```