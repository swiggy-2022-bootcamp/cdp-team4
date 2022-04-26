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

![order](https://user-images.githubusercontent.com/39910073/165148554-acd92311-18c3-4852-9dc8-d5ba82e623ba.svg)


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


### Sonarqube Reports

![shipping_sonarqube](https://user-images.githubusercontent.com/39910073/165330756-26e97eba-bcc8-4c89-8632-b368f3274372.JPG)

