FROM golang:1.22.4

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o bin ./cmd/main

EXPOSE 8080

ENTRYPOINT ["/app/bin"]