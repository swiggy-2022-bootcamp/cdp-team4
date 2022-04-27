# Category Service
This service is responsible for CRUD operations on categories which are written for admin.
![design diagram](https://github.com/swiggy-2022-bootcamp/cdp-team4/blob/category/Category/DesignDiagram.png)
### Steps to setup locally
##### Using Docker compose

```
> docker-compose up
```

##### Using Docker
```
# Build docker Image
> docker build -t cdp-team4/category .
# Run docker image
> docker run -p 8004:8004 -d --name category cdp-team4/category
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
[http://localhost:8005/swagger/index.html#/]
#Command to regenerate swagger docs
> swag init -g main.go --output docs/
```

### Routes supported
| Method | Route | Description | 
| ------ | ------ | ------ | 
| GET | / | Route Category serive health check | 
| POST | /categories/ | Route for Admin to Add a new category  | 
| GET | /categories/ | Route for Admin to Get All Categories | 
| GET | /categories/:id | Route for Admin to Get Category details given Category ID | 
| PUT | /categories/:id | Route for Admin to Update Category Details given Category ID | 
| DELETE | /categories/:id | Route for Admin to Delete a Category given Category ID | 
| DELTE | /categories/ | Route for Admin to Delete a list of Categories, provided category ids | 

