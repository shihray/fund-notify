# Dockerfile References: https://docs.docker.com/engine/reference/builder/
FROM golang:1.17-alpine AS build-env

RUN apk add --update --no-cache ca-certificates git

# Add Maintainer Info
LABEL maintainer="Ray <diskahoy@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/stocker
ADD . /go/src/stocker

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o app .

# Run the executable
ENTRYPOINT ./app
