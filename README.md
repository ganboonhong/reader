# News Reader
A platform to read news from different sources using Go.

### deployment
1. install Go on Ubuntu // https://tecadmin.net/install-go-on-ubuntu/
2. git clone this repo // `git clone git@github.com:ganboonhong/reader.git`
3. check **update binary**
4. create a log directory
5. check **run on server (in background)**

### run on server (in background)
`nohup ./main > log/main.out &` // use `fuser -k 8080/tcp` to kill a process by port number

### run on local
`go run main.go`

### update binary (on different platform, ex: mac os, ubuntu)
`go build main.go` will generate a new `main` binary code

### unit test
`go test ./...`

### news sources
News API: https://newsapi.org/

### run on Docker
**How to Build Image**

```
$ cd ~/go/src/github.com/ganboonhong/reader  // change directory to where Dockerfile resides

$ docker build -t ganboonhong/reader . // -t: image tag
```



**Create and run a new container from an image**

```
docker run -d -p 80:8080 reader 
// -d: detached mode
// -p: port mapping,  Docker_host_port:container_port 
// reader: image name
```

**Frequent Use Commands** 

```
$ docker rm CONTAINER_ID // Remove a container
$ docker stop CONTAINER_ID // Stop a container
$ docker start CONTAINER_ID // Start a container
$ docker rm -f CONTAINER_ID // Force the removal of a running container: (stop and remove) 
$ docker ps // List active containers
$ docker ps -a // List all containers
$ docker rm $(docker ps -aq) // Remove all containers

$ docker images // List images
$ docker images -f dangling=true // Show unstagged images
$ docker rmi IMAGE_ID // Remove single image
$ docker rmi $(docker images -f dangling=true -q) // Remove all unstageged images

$ docker-compose up // create and start containers according to docker-compose.yml file
```
test commit directly to master
