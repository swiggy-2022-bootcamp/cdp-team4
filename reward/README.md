# Reward Service 

With this service, the User can perform the following operations on the cart
   - Get all cart items and information for their own cart
   - Add products to the cart
   - Update the number of cart items
   - Delete item from the cart
   - Delete the cart entirely

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
