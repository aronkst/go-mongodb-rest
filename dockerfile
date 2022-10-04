FROM golang:1.18-alpine AS build

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

FROM alpine:3.16

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/go-mongodb-rest go-mongodb-rest

CMD [ "./go-mongodb-rest" ]
