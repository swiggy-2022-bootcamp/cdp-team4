# Shipping Microservie

### Steps to setup locally

#### Using Docker compose
```
docker-compose up
```

#### Without Docker
```
# Download golang dependencies
> go mod download

# Start the server
> go run main.go
```

#### Swagger Docs
```
# Open the below url
http://localhost:8012/swagger/index.html#/

# Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

PORT=8012

![shipping](https://user-images.githubusercontent.com/39910073/165150262-7c406914-d13c-4d47-b33e-ed820897e0bb.svg)

### Routes
| Method | Route                           |  Description                                               |
|  ---   | ---                             |  ---                                                       |
| POST   |   /shippingaddress              |  Route for User to Create Shipping Address                 |
| GET    |   /shippingaddress/:id          |  Route for User to get shipping address given id           |
| PUT    |   /shippingaddress/:id          |  Route for User to update shipping address given id        | 
| DELTE  |   /shippingaddress/:id          |  Route for User to delet shipping address given id         |
| POST   |   /shippingcost                 |  Route for Admin to Create Shipping Cost                   | 
| GET    |   /shippingcost/:id             |  Route for Admin to get shipping cost given city           |
| PUT    |   /shippingcost/:id             |  Route for Admin to update shipping cost given city        |
| DELTE  |   /shippingcost/:id             |  Route for Admin to delte shipping cost given city         |


