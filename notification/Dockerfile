FROM golang:1.21

# Set the Current Working Directory inside the container
WORKDIR /app


#RUN go get -u http://git.gc.io/infrastructure/infrastructure.shared.go.logger.git
#RUN go mod download -x
COPY . .

# Build the Go app
RUN go build -mod vendor -o ./app/account ./cmd/main.go


# Run the binary program produced by `go install`
CMD ["./app/account"]