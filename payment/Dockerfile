##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o payment-service main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/payment-service /payment-service

COPY .env /.env

EXPOSE 8008

ENTRYPOINT ["/payment-service"]
