# Use the official Golang image as the base image
FROM golang:1.22.5

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire project directory to the container
COPY . .

# Run the Go application
CMD ["go", "run", "."]

