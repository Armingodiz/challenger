# golang-code-challenge


## Dependencies
You must have [go](https://golang.org/doc/install) and [Docker](https://www.docker.com/) installed on your machine
also you need to `go get` this Dependencies:
name     | repo
------------- | -------------
  gorilla/mux | https://github.com/gorilla/mux
  go redis    | https://github.com/go-redis/redis 
 
## Usage

Use `docker run --name redis-usdb -p "yourPort":6379 -d redis` to connect redis to port "yourPort".

("yourPort" is set to 8282 by default, but if you want to change it, change redisPort in config.json)

Broker will be use port 8080 by default, to change it go to file config.json.

build and run **main.go** file(`go run main.go`) to start the app.

Then we need to run publish.go. Go to publisher package and run `go run publish.go` (to start the publisher for broker).
