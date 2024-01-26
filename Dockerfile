# Use a Go base image
FROM alpine

# Set the working directory inside the container
WORKDIR /server

# Copy the local package files to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o server .

# Expose a port if your application listens on a specific port
EXPOSE 9000

# Specify the command to run on container start
CMD ["./server"]
