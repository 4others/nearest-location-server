# nearest-location-server

Nearest-location-server is a simple web service that takes the source and a list of destinations
and returns a list of routes between source and each destination. 


## Cloning the repository

This repository should be cloned to ``$GOPATH/src/github.com`` directory


## Compiling the app and running on the host

To run this app, user is required to enter the following commands
```
go get -d -v ./...
go install github.com/4others/nearest-location-server
/go/bin/nearest-location-server
``` 


## Running in Docker

```
docker build -t awesome-location-app .

docker run -p 8080:8080 awesome-location-app
```


## Using the application

This service provides single GET endpoint accessible by following call
``
http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219
``
Both src and dst are defined as a pair of latitude and longitude.
Src is required as well as at least one dst parameter.

As a return, user receives list of sorted routes. These are listed by driving time and distance.