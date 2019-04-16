FROM golang:1.11.1 as builder
# We create an /go/src/github.com/ganboonhong/reader directory in which
# we'll put all of our project code
RUN mkdir -p /go/src/github.com/ganboonhong/reader
ADD . /go/src/github.com/ganboonhong/reader
WORKDIR /go/src/github.com/ganboonhong/reader
# We want to build our application's binary executable
RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# the lightweight scratch image we'll
# run our application within
FROM alpine:latest AS production
# We have to copy the output from our
# builder stage to our production stage
RUN mkdir -p /go/src/github.com/ganboonhong/reader
ENV GOPATH /go
COPY --from=builder /go /go
# we can then kick off our newly compiled
# binary exectuable!!
CMD ["/go/src/github.com/ganboonhong/reader/main"]