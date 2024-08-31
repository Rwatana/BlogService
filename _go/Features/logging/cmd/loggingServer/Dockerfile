FROM golang:1.22

WORKDIR /go/src/app
COPY . .
RUN go mod download && go mod verify
ENTRYPOINT [ "go", "run", "server.go" ]