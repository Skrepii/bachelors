FROM golang:latest

RUN mkdir -p /go/src/bachelors
WORKDIR /go/src/bachelors

ADD app/ /go/src/bachelors/app
ADD handlers/ /go/src/bachelors/handlers
ADD static/ /go/src/bachelors/static

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 80
EXPOSE 80
RUN ["echo", "$PATH"]
CMD ["go", "run", "app/app.go"]