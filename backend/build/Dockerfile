FROM golang:1.11

WORKDIR /go/src/github.com/ja1984/cogCMS/backend

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN go get -u github.com/golang/protobuf/proto
ENV PORT=5000
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"