# Transaction Service 

With this service, the Admin can perform the following operations on the cart
   - Admin can Get Transaction points for user By UserId
   - Admin can Add/SUbtract Transaction Points for user

![Design Diagram](https://github.com/swiggy-2022-bootcamp/cdp-team4/blob/transaction/transaction/TransactionService_Design.png)

## Steps to run the service locally 
#### Using Docker
```
docker-compose up
```
#### Without Docker
```
# Download golang dependencies
> go mod download

# Start the app by running main.go
> go run main.go
```
## Swagger documentation - 

```
# Open the below URL
http://localhost:8010/swagger/index.html#/

# Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

### Routes
| Method | Route                           |   Description                                                 |
|  ---   | ---                             |   ---                                                         |
| GET    |   /transaction/:userId          |   Route for admin to get transaction points of user           |
| PUT    |   /transaction/:userId          |   Route for admin to add/subtract transaction points to user  |
