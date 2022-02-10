This example helps you understand the important keyword in Docker.
## Table of contents
* [Build Docker](#Build-Docker) 
* [Run Docker](#Run-Docker) 
* [Docker volumnes](#Docker-volumnes)
* [Docker network](#Docker-network)

## Questions
* [Why this example, don't build docker image before to run?](#Build-web-server) 
* [In Docker networking, how many kind of network does we need?](#Docker-network)
* [How do you connect to the custom network on Docker](#Custom-network-driver)


## Build-Docker
Normally, we will build image with Docker, so to build it:
```
$ docker build -t hello-world .
```
However, if you don't build, and want to run directly, it connects to the default registry, in this case, https://hub.docker.com/ to retrieve it. 
Because there are many images with open-source on Docker-Hub, to see it: [here](https://hub.docker.com/search?type=image)  
That is a reason why we get the below message:
```
$ docker run --rm hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
2db29710123e: Pull complete
Digest: sha256:97a379f4f88575512824f3b352bc03cd75e239179eea0fecc38e597b2209f49a
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.
```

## Run-Docker
To run image on Docker from **Docker Hub**.  
**Note:**  
- Please sign in Docker Hub from brower and execute cmd below.
- The below command is executed on image hell-world with open-source. We can refer [here](https://hub.docker.com/_/hello-world) to see it
```
$ docker run --rm hello-world
```
To access the terminal on image Docker
* -it: interactive terminal - it maps the standard in from your terminal to the input of the running container
* -rm: To remove the container and delete any resources such as volumes it was using on exit.
```
$ docker run -it --rm alpine:latest sh
Unable to find image 'alpine:latest' locally
latest: Pulling from library/alpine
59bf1c3509f3: Already exists
Digest: sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300
Status: Downloaded newer image for alpine:latest
/ #
```

To check Filesystem on image Docker
```
/ # ls
bin    dev    etc    home   lib    media  mnt    opt    proc   root   run    sbin   srv    sys    tmp    usr    var
/ #
```
## Docker-volumnes
We have seen how Docker containers are immutable; however, there are some instances when you may wish to write some files to a disk or when you want 
to read data from a disk such as in a development setup. Docker has the concept of volumes, which
can be mounted either from the host running the Docker machine or from another Docker container.  

### Union-filesystem

### Mounting-volumes
The -v, or --volume parameter allows you to specify a pair of values corresponding to the file system you wish to mount on the host and the path where you would like to mount the volume inside the container.
* -v hostfolder:destinationfolder: volumne hostfolder to destination folder by syntax
```
$ docker run -it -v host alpine:latest /bin/sh
/ #
```
**Note:**  
  - Please avoid use cmd: "rm -rf *", because this means that any changes you have made to the volumewill be lost.
 
 ## Docker-network
 Docker networking is an interesting topic, and by default, Docker supports the following network modes:  
 * [bridge](#Bride-networking)
 * [host](#Host-networking)
 * [none](#No-networking)
 * [overlay](#Overlay-networking)

To check how many networks on Docker
```
$ docker network ls
NETWORK ID     NAME                                  DRIVER    SCOPE
988362ff75ec   07_docker-compose-go-nodejs_default   bridge    local
10b2e3d8d242   bridge                                bridge    local
5cc6efd534d1   host                                  host      local
7d732ae5c70b   none                                  null      local
```
 
### Bride-networking
 
To create a bridge network
```
$ docker network create testnetwork
```
### Host-networking
 
 
### No-networking

### Overlay-networking
 
### Custom-network-driver
To create a bridge network:
```
$ docker network create testnetwork
```
To connect a container to a customer network, with example at chapter 1:
```
$ docker run -it --rm -v $(pwd):/src -w /src --name server --network=testnetwork golang:alpine go run main.go
```
To run server on Docker
```
$ docker run --rm appropriate/curl:latest curl -i -XPOST server:8080/helloworld -d '{"name":"Nic"}'
If error
$ docker run --rm --network=testnetwork appropriate/curl:latest curl -i -XPOST server:8080/helloworld -d '{"name":"Nic"}'
```

 
 
