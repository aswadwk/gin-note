FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-note

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v
# Build the Go app
RUN go build -o ./out/go-note .

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

WORKDIR /app

COPY --from=build_base /tmp/go-note/out/go-note /app/go-note
COPY --from=build_base /tmp/go-note/.env /app/.env

# This container exposes port 8080 to the outside world
EXPOSE 3030

# Run the binary program produced by `go install`
CMD ["/app/go-note"]
