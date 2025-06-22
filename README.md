# Gym Partner API

---
* [Project's architecture](#architecture)
* [Run API with Docker](#run-api-with-docker)

---

## Architecture

For this project, I've chosen the **Clean Architecture**. It's the best way for create software with good scalability.
```
|-- core
|-- database
|-- docs
|-- env
|-- interfaces
    |-- controller
        |-- test
            |-- example.controller_test.go
        |-- example.controller.go
    |-- repository
        |-- test
            |-- example.repository_test.go
        |-- example.repository.go
|-- k8s
|-- middleware
|-- mock
|-- model
|-- router
|-- usecases
    |-- interactor
        |-- test
            |-- example.interactor_test.go
        |-- example.interactor.go
    |-- repository
        |-- example.repository.go
|-- utils
|-- main.go

```

---
## Run API with Docker

For running API with Docker, you need Docker, so install them and go to the next stage.  
After this, you have to open the terminal -> and run this command:  
`docker-compose up -d`  

While Gym Partner API it's running ! By the way, you have a new folder in your PC:  
**~/Documents/gym-partner-docker-volumn** Inside this you have the api's logs and the postgres data

## Way to using this API

For better practice, one swagger was implement. For the case was using the API in local, you have  
this host: **http://localhost:4200/swagger/index.html**  
Command for init swagger: ``swag init``

Then you can use the swagger for all routes, request and response.