FROM golang:1.22-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/rinha

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-rinha ./cmd/main.go

# Start fresh from a smaller image
FROM alpine:3.19

COPY --from=build_base /tmp/rinha/out/go-rinha /app/go-rinha

EXPOSE 5000

# Run the binary program produced by `go install`
CMD ["/app/go-rinha"]