FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir mongodb
RUN mkdir web

COPY mongodb/*.go mongodb
COPY web/*.go web

COPY main.go .

RUN go build -o go-mongodb-rest main.go

CMD [ "./go-mongodb-rest" ]
