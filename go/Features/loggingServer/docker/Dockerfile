FROM golang:1.22

WORKDIR /go/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build ./cmd/server
ENTRYPOINT [ "./server" ]
