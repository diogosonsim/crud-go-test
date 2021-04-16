## CRUD-GO-TEST

- Clone this repository

- To create the database run:

  - `docker-compose -f "docker-compose.yml" up -d --build`

- You can build the project and run with `go build .` or `go run .`

- You can test if the application is up at `localhost:3000/api/public/healthcheck`

### Use API - Swagger

```
{url}/swagger/index.html
```

### API's

```

POST {url}/api/public/register - Public route to create users
{
    "name": "name",
    "age": 0,
    "email": "email",
    "password": "password",
    "confirm_password": "confirm_password"
}

POST {url}/api/public/login - Public route to create your session.
{
    "email": "email",
    "password": "password"
}

GET {url}/api/private/logged - Private route to retrieve logged data.

POST {url}/api/private/logout - Private route to logout your session.

GET {url}/api/private/users - Private route to retrieve all users.

GET {url}/api/private/users/:id - Private route to retrieve selected user.

POST {url}/api/private/users - Private route to create a new user.
{
    "name": "name",
    "age": 0,
    "email": "email",
    "address": "address",
    "password": "password",
    "confirm_password": "confirm_password"
}

PUT {url}/api/private/users - Private route to update selected user.
{
    "id": 0,
    "name": "name",
    "age": 0,
    "email": "email",
    "address": "address"
}

DELETE {url}/api/private/users/:id - Private route to delete selected user.

```
