# Reward Service 

With this service, the Admin can perform the following operations on the cart
   - Admin can Get Reward points for user By UserId
   - Admin can Add/SUbtract Reward Points for user

![Design Diagram](https://github.com/swiggy-2022-bootcamp/cdp-team4/blob/reward/reward/RewardService_Design.png)

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
| Method | Route                           |   Description                                               |
|  ---   | ---                             | ---                                                         |
| GET    |   /reward/:userId               |   Route for admin to get reward points of user              |
| PUT    |   /reward/:userId               |   Route for admin to add/subtract reward points to user     |
