# go-communication

communication

## Development

### set up dependency services

1. familia

``` 
docker run -d \
    --name familia \
    -e MODEL_NAME=news \
    -p 5000:5000 \
    orctom/familia
```

### set up golang project

1. go mod

### run cmd

``` 
go run ./main.go
```
