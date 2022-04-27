FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/order

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/order

# get all the dependencies
RUN go get ./...

# build
RUN go build -o /orderService

EXPOSE 8007

CMD [ "/orderService" ] 