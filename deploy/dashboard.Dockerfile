FROM golang:1.24 AS build-stage

# Create a separate folder for building the application
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy across the source
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./cmd/web ./cmd/web

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /output/web cmd/web/main.go

# Run the application tests
# TODO: When we write some tests we need to reenable this bit
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:3.21 AS build-release-stage

# Copy across the built binary from the build stage
WORKDIR /
COPY --from=build-stage /output /usr/bin

# Define the entry point
ENTRYPOINT ["/usr/bin/web", "-address", "0.0.0.0", "-port", "1234"]