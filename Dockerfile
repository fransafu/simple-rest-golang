FROM golang:1.14

LABEL maintainer="fransafu@gmail.com"

ENV PORT=8080

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]
