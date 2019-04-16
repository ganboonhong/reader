FROM golang:1.11.1 as builder
# We create an /go/src/github.com/ganboonhong/reader directory in which
# we'll put all of our project code
WORKDIR /go/src/github.com/ganboonhong/reader
ADD . .
# We want to build our application's binary executable
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# the lightweight scratch image we'll
# run our application within
FROM alpine:latest AS production

# newsapi defaultBaseURL is "https://newsapi.org/v2/"
# use `docker run -it reader sh` to enter shell
# go to /etc/ssl/certs and list this directory, you will the cert files
RUN apk --no-cache add ca-certificates
# We have to copy the output from our
# builder stage to our production stage
WORKDIR /go/src/github.com/ganboonhong/reader
ENV GOPATH /go
COPY --from=builder /go /go
EXPOSE 8080
# we can then kick off our newly compiled
# binary exectuable!!
CMD ["./main"]