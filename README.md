# Go Backend Book Application

This project is an application built on a microservices architecture using the Go language. Conceptually, the application consists of an Auth Service for handling authorization, authentication, and user storage, a Review Service, and a Book Service. These microservices communicate with each other through gRPC. This project will evolve alongside my learning and professional growth.

Please feel free to leave your comments if you notice any weak points or have questions about architectural or code design decisions. Thank you!

### User Service

User Service allows users to create accounts, log in by providing an access token, and fetch other users by id.

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `POST`   | `/api/signup`                            | Create new user,                         |
| `POST`   | `/api/login`                             | Log in by credentials.                   |
| `GET`    | `/api/user/:id`                          | Fetch user by id.                        |


#### POST /signup:

```
Request:

curl --location 'localhost:3002/api/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "temryakov@gmail.com",
    "name": "Maxim Temryakov",
    "password": "qwerty123"
}'

Response:

{
    "message": "You have sign up successfully! %)"
}
```

#### POST /login:

```
Request:

curl --location 'http://localhost:3002/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "temryakov@gmail.com",
    "password": "qwerty123"
}'

Response:
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiTWF4aW0gVGVtcnlha292IiwiaWQiOjIsImV4cCI6MTcwMTE4NzYyOH0.T3_CA-O53_Cnv1fBSQW2_8PHJo8GG7PyLNS8MmshmsE",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiZXhwIjoxNzAxMTg3NjI4fQ.cq1N8hmhiItsNHqQhjZzLTXswkJ26EYlW9e1umRMIsY"
}
```

#### GET /user/{id}:

```
Request:

curl --location 'localhost:3002/api/user/1'

Response:

{
    "email": "temryakov@gmail.com",
    "name": "Maxim Temryakov"
}
```



