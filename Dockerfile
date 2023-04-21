# Gunakan golang:latest sebagai base image
FROM golang:latest

# Tetapkan direktori kerja default
WORKDIR /app

# Salin go.mod dan go.sum untuk mendownload dependensi
COPY go.mod .
COPY go.sum .
RUN go mod download

# Salin kode sumber aplikasi
COPY . .

# Kompilasi aplikasi
RUN go build -o app .

# Gunakan argumen untuk variabel lingkungan
ARG MYSQL_HOST
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_DBNAME
ENV MYSQL_HOST=$MYSQL_HOST MYSQL_USER=$MYSQL_USER MYSQL_PASSWORD=$MYSQL_PASSWORD MYSQL_DBNAME=$MYSQL_DBNAME

# Tetapkan port default
EXPOSE 3030

# Tetapkan perintah default
CMD ["./app"]
