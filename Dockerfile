# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS builder
LABEL maintainer="manoj_men@hotmail.com"

# Copy the source code
RUN mkdir /janes_book_store
WORKDIR /janes_book_store
COPY . .

# Download dependencies and build the executable
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go get -d -v ./...
RUN go build -ldflags "-s" -a -installsuffix cgo -o janes_book_store .

# Copy binary over to alpine image
FROM golang:1.18-alpine
WORKDIR /root
COPY --from=builder ./janes_book_store .
RUN apk add --no-cache postgresql-client
# Run the executable
EXPOSE 8080
CMD ["/root/janes_book_store"]