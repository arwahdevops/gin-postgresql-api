# Sample Gin Gonic api with PostgreSQL database
Coding for fun 🌱
## Run
`go run main.go`
## Endpoint
### - Register
POST `http://localhost:8080/register` \
```
{
    "name": "Riki Permana", 
    "email": "arwahdevops@gmail.com", 
    "password": "1234567890"
}
```
### - Login
POST `http://localhost:8080/login` \
```
{
    "email": "arwahdevops@gmail.com", 
    "password": "1234567890"
}
```
### - Update
PUT `http://localhost:8080/user/:id` \
```
{
    "name": "Riki Permana", 
    "email": "arwahdevops@gmail.com", 
    "password": "qwebnm123"
}
```
### - Delete
DELETE `http://localhost:8080/user/:id`

