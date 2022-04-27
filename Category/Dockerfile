FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/Category

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/Category

# get all the dependencies
RUN go get ./... 

# build
RUN go build -o /CategoryService

EXPOSE 8005

CMD [ "/CategoryService" ]
