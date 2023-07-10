# Create build stage based on buster image
FROM golang:alpine
RUN apk add --update redis
CMD ["redis-server"]
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY go.* ./
# Install any required modules
RUN go mod download
# Copy over Go source code
COPY *.go ./
# Run the Go build and output binary under hello_go_http
RUN go build -o /hello_go_http
# Make sure to expose the port the HTTP server is using
EXPOSE 8080
# Run the app binary when we run the container
ENTRYPOINT ["/hello_go_http"]
