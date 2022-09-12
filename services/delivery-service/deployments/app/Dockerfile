# Get Go image from DockerHub.
FROM golang:1.17.8 AS builder

LABEL maintainer="Luis Alfredo Torres Cabrera <luis.torres.net@gmail.com>"

# Set working directory.
WORKDIR /builder

# Copy dependency locks so we can cache.
COPY ${SERVICE_APP_PATH}/go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY ${SERVICE_APP_PATH}/.. ./

# Set necessary environment variables needed for our image and build the API server.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o delivery-service ./cmd/delivery-service/main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=builder /builder/delivery-service .

# Run application and expose port 8080.
EXPOSE 5000
CMD ["./delivery-service"]
