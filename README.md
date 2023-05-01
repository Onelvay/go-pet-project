# Pet Project

This project was created in order to show my hard skills.

## Installing

To start the project, follow these steps:

1. Clone the repository: `https://github.com/Onelvay/go-pet-project.git `
2. Go to the project directory: `cd go-pet-project`
3. Launch the application: `make run`

## Technologies

* Docker
* Postgres, Mongo, Redis
* JWT, REST API
* Simulation of payment via fondy

## Project structure
<p>
pet-project/ <br>
├── config/<br>
│   └── handler/<br>
├── db/<br>
│   ├── mongoDB/<br>
│   └── postgres/<br>
├── payment/<br>
│   ├── Request/<br>
│   │   └── request.go<br>
│   └── client/<br>
│       └── client.go<br>
├── pkg/<br>
│   ├── controller/<br>
│   │   ├── handlers.go<br>
│   │   ├── redis.go<br>
│   │   ├── token.go<br>
│   │   ├── user.go<br>
│   │   └── postgres.go<br>
│   ├── domain/<br>
│   │   └── models.go<br>
│   ├── handlers/<br>
│   │   ├── handlers.go<br>
│   │   └── middleware.go<br>
│   ├── routes/<br>
│   │   └── routes.go<br>
│   └── service/<br>
│       ├── interfaces.go<br>
│       └── hash.go<br>
└── redis/<br>
    └── redis.go<br>
</p>
