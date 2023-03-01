# Start with a Go image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port that the application listens on
EXPOSE 8081

# Start the application
CMD ["./app"]
