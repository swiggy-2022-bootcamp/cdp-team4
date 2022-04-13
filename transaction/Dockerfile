##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o transaction-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/transaction-app /transaction-app

COPY .env /.env

EXPOSE 8007

ENTRYPOINT ["/transaction-app"]

