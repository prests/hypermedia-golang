FROM golang:1.20-alpine

ENV PORT=8080

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app
COPY ./ ./

EXPOSE $PORT

ENTRYPOINT CompileDaemon --build="go build -o ./build/app ./cmd/main.go" -command="./build/app $PORT" -build-dir=/app -polling -polling-interval=10
