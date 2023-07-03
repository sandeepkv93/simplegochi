# Official Go image to build the binary as base image for build stage.
FROM golang:1.16 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

# Use a distroless static image for the final stage of the build.
FROM gcr.io/distroless/static

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage.
COPY --from=builder /app/main /app/main

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the binary
CMD ["/app/main"]