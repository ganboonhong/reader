FROM golang:1.11.1 as builder
# We create an /app directory in which
# we'll put all of our project code
RUN mkdir $GOPATH/src/github.com/ganboonhong/reader
ADD . $GOPATH/src/github.com/ganboonhong/reader
WORKDIR $GOPATH/src/github.com/ganboonhong/reader
# We want to build our application's binary executable
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# the lightweight scratch image we'll
# run our application within
FROM alpine:latest AS production
EXPOSE 8080
# We have to copy the output from our
# builder stage to our production stage
COPY --from=builder $GOPATH/src/github.com/ganboonhong/reader .
# we can then kick off our newly compiled
# binary exectuable!!
CMD ["./main"]