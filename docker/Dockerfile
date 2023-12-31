# syntax=docker/dockerfile:1

FROM golang:1.20 as builder

# Set destination for COPY
WORKDIR /build

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
ADD internal ./internal

# Download Go modules
COPY go.mod ./
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 80

# Build the production image
FROM alpine:latest as app

LABEL maintainer="ivan.kuchin@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/IvanKuchin/${BUILD_REPO}"

COPY --from=builder /app /app
EXPOSE 1025
CMD ["/app"]
