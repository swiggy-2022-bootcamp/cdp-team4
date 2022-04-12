##
## Build
##

FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o reward-app main.go

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /app/reward-app /reward-app

COPY .env /.env

EXPOSE 8006

ENTRYPOINT ["/reward-app"]


# FROM golang:1.17-alpine

# RUN apk update && apk add --no-cache git build-base

# WORKDIR /go/github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin

# COPY . /go/github.com/swiggy-2022-bootcamp/cdp-team4/Product_Admin

# # get all the dependencies
# RUN go get ./... 

# # build
# RUN go build -o /ProductAdminService

# EXPOSE 8081

# CMD [ "/ProductAdminService" ]