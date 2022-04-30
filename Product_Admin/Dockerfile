##
## Build
##

FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o product-admin-service main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/product-admin-service /product-admin-service

COPY .env .env

EXPOSE 8004
EXPOSE 7003

ENTRYPOINT ["/product-admin-service"]