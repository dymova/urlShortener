# Use the official Go base image with the desired version
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o app

# Set the command to run when the container starts
CMD ["./app"]