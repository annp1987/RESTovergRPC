# RESTovergRPC
Create a directory services with rest over grpc

# Testing local

Run postgres
============
```console
$ docker run --rm -d --name postgres -p 1234:1234 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=postgres postgres
```

Run server
==========

```
$ go run main.go
```

