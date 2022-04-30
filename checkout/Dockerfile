##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o checkout-service main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/checkout-service /checkout-service

COPY .env /.env

EXPOSE 8011
EXPOSE 7011

ENTRYPOINT ["/checkout-service"]

