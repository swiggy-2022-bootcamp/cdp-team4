# User Service

- HTTP PORT = 8002
- GRPC PORT = 7002

![user service](https://user-images.githubusercontent.com/53436195/165159945-f9100c43-497e-465f-a34a-250f4ea43397.png)


### Routes
| Method | Route                           |  Description                                               |
|  ---   | ---                             |  ---                                                       |
| POST   |   /user                         |  Route for User Creation                                   |
| GET    |   /users                        |  Route for getting all the users                           |
| GET    |   /user/:id                     |  Route for getting user by given id                        | 
| PATCH  |   /user/:id                     |  Route for updating user by given id                       |
| DELETE |   /user/:id                     |  Route for deleting user by given id                       | 
