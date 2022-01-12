FROM golang:1.16-alpine

RUN mkdir -p /app/src

WORKDIR /app/src

COPY . .

CMD ["go", "run", "main.go"]