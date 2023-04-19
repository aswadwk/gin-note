FROM golang:latest

ENV APP_NAME gin-todo
ENV APP_PATH /go/src/github.com/aswadwk/gin-todo

WORKDIR $APP_PATH

COPY . $APP_PATH

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3030

CMD ["gin-todo"]
