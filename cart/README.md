# Reward Service 

With this service, the User can perform the following operations on the cart
   - Get all cart items and information for their own cart
   - Add products to the cart
   - Update the number of cart items
   - Delete a list of item from the cart
   - Delete the cart entirely

![Design Diagram](https://github.com/swiggy-2022-bootcamp/cdp-team4/blob/reward/reward/Cart-Design.png)

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
| Method | Route                         |   Description                                               |
|  ---   | ---                           | ---                                                         |
| POST   |   /cart                       |   Route to make a new cart                         |
| GET    |   /cart/:userId               |   Route to fetch info about their cart             |
| PUT    |   /cart/:userId               |   Route to add Items to cart or update quantity    |
| DELETE |   /cart/empty/:userId         |   Route to Delete their cart                       |
| DELETE |   /cart/:userId               |   Route to Delete a list of Items from the cart    |

