FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/User

COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/User

# get all the dependencies
RUN go get ./... 

# build
RUN go build -o /userService

EXPOSE 8002
EXPOSE 7002

CMD [ "/userService" ]