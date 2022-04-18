FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/Checkout

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/Checkout

# get all the dependencies
RUN go get ./... 

# build
RUN go build -o /CheckoutService

EXPOSE 8081

CMD [ "/CheckoutService" ]
