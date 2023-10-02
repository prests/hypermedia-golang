FROM golang:1.20-alpine

ENV PORT=8080

WORKDIR /app
ADD . .

RUN apk --no-cache add curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.3/tailwindcss-linux-x64 \
  && chmod +x tailwindcss-linux-x64 \
  && mv tailwindcss-linux-x64 tailwindcss

RUN ./tailwindcss -i internal/input.css -o web/assets/output.css --minify

RUN go build -o ./build/app ./cmd/main.go

EXPOSE $PORT

CMD ./build/app $PORT
