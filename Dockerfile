FROM golang:1.18 as builder
# We create an /go/src/github.com/ganboonhong/reader directory in which
# we'll put all of our project code
WORKDIR /go/src/github.com/ganboonhong/reader

# copy from local to image
ADD . .

# get Go dependencies recursively
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# the lightweight Alpine Linux image we'll
# run our application within
FROM alpine:latest AS production

# newsapi defaultBaseURL is "https://newsapi.org/v2/"
# use `docker run -it reader sh` to enter shell
# go to /etc/ssl/certs and list this directory, you will the cert files
RUN apk --no-cache add ca-certificates

# We have to copy the output from our
# builder stage to our production stage
WORKDIR /go/src/github.com/ganboonhong/reader

# set $GOPATH environment variables 
# (article_service.go::ArticlePageHandler needs to use absolute path, otherwise go test ./... will fail)
ENV GOPATH /go

# copy from builder stage to production stage
COPY --from=builder /go/src/github.com/ganboonhong/reader .

# start running app
CMD ["./main"]

# ref
# https://tutorialedge.net/golang/go-multi-stage-docker-tutorial/
# https://flaviocopes.com/golang-docker/
