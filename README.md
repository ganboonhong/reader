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

### run on local (with docker command)
**Build**  
Build image:
`cd ~/go/src/github.com/ganboonhong/reader`  // cd to reader root directory where Dockerfile resides
`docker build -t reader .`  
(-t: image tag)

**Run**  
Run Docker image to get Docker container:  
`docker run -d -p 80:8080 reader`  
(-d: detached mode, -p: port mapping,  Docker_host_port:container_port, reader: image name)

**Container**  
Remove a container: `docker rm CONTAINER_ID`  
Stop a container: `docker stop CONTAINER_ID`  
Start a container: `docker start CONTAINER_ID`  
Force the removal of a running container: `docker rm -f CONTAINER_ID` (stop and remove)  
List active containers: `docker ps`  
List all containers: `docker ps -a`  
Remove all containers: `docker rm $(docker ps -aq)`  

**Image**  
List images: `docker images`
Show unstagged images: `docker images -f dangling=true` (-f: filter)  
Remove single image: `docker rmi IMAGE_ID`  
Remove all unstageged images: `docker rmi $(docker images -f dangling=true -q)`


### run on local (with docker-compose cmd)
Build image and start container:  
`cd ~/go/src/github.com/ganboonhong/reader`  // cd to reader root directory where docker-compose.yml resides  
`docker-compose up`  
