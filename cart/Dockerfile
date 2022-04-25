##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o cart-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/cart-app /cart-app

COPY .env /cart-app

EXPOSE 8005

ENTRYPOINT ["/cart-app"]
