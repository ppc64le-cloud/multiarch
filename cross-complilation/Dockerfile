# Dockerfile

# Use the official golang base image
FROM --platform=${BUILDPLATFORM} golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the source code into the container
COPY main.go main.go

ARG TARGETOS
ARG TARGETARCH
# Build the Go application
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o multiarch-app .

# Final stage to create a lightweight production image
FROM gcr.io/distroless/static:nonroot
# Copy the binary from the build stage into the final image
COPY --from=build /app/multiarch-app /multiarch-app

# Set the command to run the executable
CMD ["/multiarch-app"]
