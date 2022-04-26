# Order Microservie

### Steps to setup locally

#### Using Docker compose
```
docker-compose up
```

#### Without Docker
```
# Download golang dependencies
> go mod download

# Start the CMD server
> go run main.go
```

#### Swagger Docs
```
# Open the below url
http://localhost:8007/swagger/index.html#/

# Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

![order](https://user-images.githubusercontent.com/39910073/165333266-82cf2838-42e9-4ce2-8138-1d04c459ae2d.svg)


### Routes
| Method | Route                           |   Description                                               |
|  ---   | ---                             | ---                                                         |
| POST   |   /order                        |   Route for Admin to Create Order                           |
| GET    |   /orders                       |   Route for Admin to get all Orders                         |
| GET    |   /order/:id                    |   Route for Admin to get order info given order id          |
| GET    |   /order/user/:userid           |   Route for Admin/User to get order info given user id      |
| GET    |   /order/status/:status         |   Route for Admin to get order info given status            |
| DELETE |   /order/:id                    |   Route for Admin/User to cancel/delete order               |
| POST   |   /confirm/:user_id             |   Route for User to confirm Order                           |
| GET    |   /order/invoice/:order_id      |   Route for User to get Invoice of a order                  |


