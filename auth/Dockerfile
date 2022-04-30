
#FROM golang:1.18-alpine
#
#RUN apk update && apk add --no-cache git build-base
#
#WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/auth
#
#COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/auth
#
## get all the dependencies
#RUN go get ./...
#
## build
#RUN go build -o /authService
#
#EXPOSE 8001
#
#CMD [ "/authService" ]

##
## Build
##

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o auth-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/auth-app /auth-app

COPY .env .env

EXPOSE 8001

ENTRYPOINT ["/auth-app"]