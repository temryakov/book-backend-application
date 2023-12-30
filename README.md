# Book Backend Application

This project is a microservices-based application developed in Go. It essentially comprises three services: an Authentication Service, a Review Service, and a Book Service. Each microservice operates within its own container and has a dedicated PostgreSQL database. They interact through a REST API using the protobuf format. This project will evolve alongside my learning and professional growth. 

### User Service

User Service is a service which provides ability to register in application, sign in by retreiving access token and use it for following requests, get information about self profile and fetch information about users by their id for another users and services.

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `POST`   | `/api/signup`                            | Create new user,                         |
| `POST`   | `/api/login`                             | Log in by credentials.                   |
| `GET`    | `/api/user/:id`                          | Fetch user by id.                        |
| `GET`    | `/api/profile`                           | Fetch information about your profile.    |

### Book Service

Book Service is book-related service which provides users which have certain permissions to manage books data in system and other users and services to retrieve book information.

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/api/book/all`                          | Fetch all existing books.                |
| `GET`    | `/api/book/:id`                          | Fetch book by id.                        |
| `POST`   | `/api/book`                              | Add new book.                            |
| `PATCH`  | `/api/book/:id`                          | Update book by id.                       |
| `DELETE` | `/api/book/:id`                          | Delete book by id.                       |

### Review Service

Review Service focuses on user-generated reviews. It enables users to post and control their reviews. Additionally, this service integrates data from both user-service and book-service, ensuring that requests are verified and responses are accurately constructed based on relevant data.

| Method   | URL                                      | Description                              |
| -------- | ---------------------------------------- | ---------------------------------------- |
| `GET`    | `/api/review/:id`                        | Fetch review by id.                      |
| `POST`   | `/api/review`                            | Add new book.                            |





