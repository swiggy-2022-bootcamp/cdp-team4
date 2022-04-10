
##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o shipping-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/shipping-app /shipping-app

EXPOSE 7001

ENTRYPOINT ["/shipping-app"]