
##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o orders-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/orders-app /orders-app

EXPOSE 7000

ENTRYPOINT ["/category-app"]