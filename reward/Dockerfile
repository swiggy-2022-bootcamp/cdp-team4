##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o reward-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/reward-app /reward-app

COPY .env /.env

EXPOSE 8010

ENTRYPOINT ["/reward-app"]
