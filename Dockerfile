FROM golang:latest
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
CMD ["go", "run", "main.go", "-start=/app/env"]