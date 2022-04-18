FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/payment

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/payment

# get all the dependencies
RUN go get ./... 

# build
RUN go build -o /paymentService

EXPOSE 8081

CMD [ "/PaymentService" ]
