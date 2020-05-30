FROM golang:1.14.2-alpine

WORKDIR /app

COPY ./ .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o bin/main cmd/main.go" --command=./bin/main
