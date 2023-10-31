FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app.exe
EXPOSE 8080
CMD ["/app.exe", "-p", "8080"]
