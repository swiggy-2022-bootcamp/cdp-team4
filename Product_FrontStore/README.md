# Product Admin Service
This service is responsible for fetcing product details for the front store/ user. 

![design diagram](https://github.com/swiggy-2022-bootcamp/cdp-team4/blob/product_front_store/Product_FrontStore/DesignDiagram.png)
### Steps to setup locally
##### Using Docker compose

```
> docker-compose up
```

##### Using Docker
```
# Build docker Image
> docker build -t cdp-team4/product-front-store .
# Run docker image
> docker run -p 8003:8003 -d --name product-front-store cdp-team4/product-front-store
```

##### Without docker 

```
# Download golang dependencies
> go mod download
# Start the CMD server
> go run cmd/main.go
```

##### Swagger docs - 
```
Swagger documentation can be found at the following url - 
[http://localhost:8003/swagger/index.html#/]
#Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

### Routes supported
| Method | Route | Description | 
| ------ | ------ | ------ | 
| GET | / | Route for Product service health check | 
| GET | /products/ | Route for Admin to Get All Products | 
| GET | /products/:id | Route for Admin to Get Product details given Product ID | 
| PUT | /products/category/:id | Route for Admin to Get Product Details given Category ID | 

