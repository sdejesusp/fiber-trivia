## Simple REST API using Fiber (Golang)

## Overview
This is a simple REST API using the Golang framework Fiber. 

User can do the following operations:

  - Create a new records with a questions and an answers - Facts
  - List all records in the DB
  - List an individual record by ID
  - Update a record by ID
  - Delete a record by ID
---
This project uses the following technologies:

   - Fiber: Golang framework
   - PostgreSQL
   - Docker
   - Swagger middleware
   - Air - Live reload for Go apps
---
## How to Run
Before to run the project, the following variables must to be in your .env file:
```
- DB_USER
- DB_PASSWORD
- DB_NAME

Example: 

DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
```

To run the application, execute the following:
```
❯ docker-compose up

It will look like this:
❯ docker-compose ps

NAME                COMMAND                  SERVICE             STATUS              PORTS
db                  "docker-entrypoint.s…"   db                  running             0.0.0.0:5434->5432/tcp
web                 "air cmd/main.go -b …"   web                 running             0.0.0.0:3000->3000/tcp
```

---
## API Swagger documentation

![1- Fiber API Docs](https://github.com/sdejesusp/fiber-trivia/assets/36713176/bfec2744-4a26-4a94-b50c-bc5d0549f49e)

