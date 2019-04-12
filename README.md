# News Reader
A platform to read news from different sources

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
