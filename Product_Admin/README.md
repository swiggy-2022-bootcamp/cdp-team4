# Product Admin Service
This service is responsible for CRUD operations on products which are written for admin.
### Steps to setup locally
##### Using Docker compose

```
> docker-compose up
```

##### Using Docker
```
# Build docker Image
> docker build -t cdp-team4/product-admin .
# Run docker image
> docker run -p 8004:8004 -d --name product-admin cdp-team4/product-admin
```

##### Without docker 

```
# Download golang dependencies
> go mod download
# Start the CMD server
> go run cmd/main.go
```
##### Run kafka as subscriber with following topic - 
```
cart
```
##### Swagger docs - 
```
Swagger documentation can be found at the following url - 
[http://localhost:8004/swagger/index.html#/]
#Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

### Routes supported
| Method | Route | Description | 
| ------ | ------ | ------ | 
| POST | /products/ | Route for Admin to Create Product | 
| GET | /products/ | Route for Admin to Get All Products | 
| GET | /products/:id | Route for Admin to Get Product details given Product ID | 
| PUT | /products/:id | Route for Admin to Update Product Details given Product ID | 
| DELETE | /products/:id | Route for Admin to Delete a Product given Product ID | 
| GET | /products/search/:searchstring | Route for Admin to Search a Product given a search string | 

###### Reference- 
    https://weihungchin.medium.com/how-to-set-up-sonarqube-in-windows-mac-and-linux-using-docker-3959c5a95eb2