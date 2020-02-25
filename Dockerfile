FROM golang:1.13.4-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build cmd/recipes/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ./main
