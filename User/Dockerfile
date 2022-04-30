##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o user-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/user-app /user-app

COPY .env .env

EXPOSE 8002
EXPOSE 7002

ENTRYPOINT ["/user-app"]
