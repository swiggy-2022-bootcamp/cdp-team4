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

EXPOSE 8006

ENTRYPOINT ["/cart-app"]


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