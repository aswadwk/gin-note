# Stage 1: Build the application
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /app/app

# Stage 2: Create a minimal image with the compiled binary and .env file
FROM golang:latest AS runner
WORKDIR /app
COPY --from=builder /app/app /app
COPY .env /app/.env
ARG MYSQL_HOST
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_DBNAME
ENV MYSQL_HOST=$MYSQL_HOST MYSQL_USER=$MYSQL_USER MYSQL_PASSWORD=$MYSQL_PASSWORD MYSQL_DBNAME=$MYSQL_DBNAME
EXPOSE 3030
ENTRYPOINT ["/app/app"]


