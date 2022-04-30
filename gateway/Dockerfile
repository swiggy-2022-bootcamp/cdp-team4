##
## Build
##

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o gateway-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/gateway-app /gateway-app

COPY .env .env

EXPOSE 8000

ENTRYPOINT ["/gateway-app"]