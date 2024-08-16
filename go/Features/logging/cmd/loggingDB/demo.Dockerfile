FROM golang:1.22

WORKDIR /go/src/app
COPY . .
RUN go mod download && go mod verify
## そのまま起動させておく
ENTRYPOINT [ "tail", "-f", "/dev/null" ]
