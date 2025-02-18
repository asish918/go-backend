# Go Backend Clean Architecture

A Go (Golang) Backend Clean Architecture project with Gin, MongoDB, JWT Authentication Middleware, Test, and Docker.

**You can use this project as a template to build your Backend project in the Go language on top of this project.**

Thanks to all those projects, I learned a lot from all of those. As I keep saying:

> The best way to learn to code is to code. But, to write good code, you will also have to read good code. Make a habit of reading good code. You can find many open-source projects on GitHub and start reading.

Then for the implementation part, I combined all of my ideas, experiences, and learnings from those projects to create this project.

And as always I would love to get feedback on my project. This helps everyone and most importantly me.

## Architecture Layers of the project

- Router
- Controller
- Usecase
- Repository
- Domain

#### Run without Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `MongoDB` if not installed on your machine.
- Important: Change the `DB_HOST` to `localhost` (`DB_HOST=localhost`) in `.env` configuration file. `DB_HOST=mongodb` is needed only when you run with Docker.
- Run `go run cmd/main.go`.
- Access API using `http://localhost:8080`

#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:8080`

### How to run the test?

```bash
# Run all tests
go test ./...
```

### How to generate the mock code?

In this project, to test, we need to generate mock code for the use-case, repository, and database.

```bash
# Generate mock code for the usecase and repository
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

# Generate mock code for the database
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all
```

Whenever you make changes in the interfaces of these use-cases, repositories, or databases, you need to run the corresponding command to regenerate the mock code for testing.

### The Complete Project Folder Structure

```
.
├── Dockerfile
├── api
│   ├── controller
│   │   ├── login_controller.go
│   │   ├── profile_controller.go
│   │   ├── profile_controller_test.go
│   │   ├── refresh_token_controller.go
│   │   ├── signup_controller.go
│   │   └── task_controller.go
│   ├── middleware
│   │   └── jwt_auth_middleware.go
│   └── route
│       ├── login_route.go
│       ├── profile_route.go
│       ├── refresh_token_route.go
│       ├── route.go
│       ├── signup_route.go
│       └── task_route.go
├── bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── domain
│   ├── error_response.go
│   ├── jwt_custom.go
│   ├── login.go
│   ├── profile.go
│   ├── refresh_token.go
│   ├── signup.go
│   ├── success_response.go
│   ├── task.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   └── tokenutil
│       └── tokenutil.go
├── mongo
│   └── mongo.go
├── repository
│   ├── task_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── usecase
    ├── login_usecase.go
    ├── profile_usecase.go
    ├── refresh_token_usecase.go
    ├── signup_usecase.go
    ├── task_usecase.go
    └── task_usecase_test.go
```

### Example API Request and Response

- signup

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/signup' \
  --data-urlencode 'email=test@gmail.com' \
  --data-urlencode 'password=test' \
  --data-urlencode 'name=Test Name'
  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- login

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/login' \
  --data-urlencode 'email=test@gmail.com' \
  --data-urlencode 'password=test'
  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```

- profile

  - Request

  ```
  curl --location --request GET 'http://localhost:8080/profile' \
  --header 'Authorization: Bearer access_token'
  ```

  - Response

  ```json
  {
    "name": "Test Name",
    "email": "test@gmail.com"
  }
  ```

- task create

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/task' \
  --header 'Authorization: Bearer access_token' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'title=Test Task'
  ```

  - Response

  ```json
  {
    "message": "Task created successfully"
  }
  ```

- task fetch

  - Request

  ```
  curl --location --request GET 'http://localhost:8080/task' \
  --header 'Authorization: Bearer access_token'
  ```

  - Response

  ```json
  [
    {
      "title": "Test Task"
    },
    {
      "title": "Test Another Task"
    }
  ]
  ```

- refresh token

  - Request

  ```
  curl --location --request POST 'http://localhost:8080/refresh' \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'refreshToken=refresh_token'
  ```

  - Response

  ```json
  {
    "accessToken": "access_token",
    "refreshToken": "refresh_token"
  }
  ```
