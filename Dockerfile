FROM golang:latest

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY app/ /go/src/app
COPY handlers/ /go/src/app
COPY static/ /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 8080
EXPOSE 8080
RUN ["echo", "$PATH"]
CMD ["go", "run", "app.go"]