# octopus
A distributed crawler collection program implemented in Go language



>For learning only, commercial use is strictly prohibited

>The idea comes from the ccmouse course
---

## Features
* Simple: microservice, gRPC request
* Expand: etcd cluster registration and monitoring service
* Distributed: Build multiple containers that include main,worker,itemsaver
* Data: ES storage, Kibana show, RabbitMQ message broker
* Deploy: Docker compose manage container
* Rule: bloom filter, Concurrent scheduling
* Supply: go mod, debug in container

## Getting started
#### Go build
- Make build main, worker and itemsave:
```
cd octopus && make go-build
cd octopus/worker/server && make go-build
cd octopus/persist/server && make go-build
```
    
#### Docker running
- Docker compose on Mac
```
cd octopus && docker-compose up -d
```
    
#### Kibana show
- Request the following URL on the browser
```
http://localhost:5601/app/kibana
```

#### debug supply
- Debug with dlv in the container
```
#use Dockerfile.debug in docker-compose.yml
```
