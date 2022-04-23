FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/User

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/User

# get all the dependencies
RUN go get ./... 

# build
RUN go build -o /UserService

EXPOSE 8001

CMD [ "/UserService" ]