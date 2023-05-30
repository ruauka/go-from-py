FROM golang:1.19.2
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -buildmode=c-shared -o go_app_linux.so main.go