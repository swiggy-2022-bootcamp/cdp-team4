
FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/auth

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/auth

# get all the dependencies
RUN go get ./...

# build
RUN go build -o /authService

EXPOSE 8001

CMD [ "/authService" ]