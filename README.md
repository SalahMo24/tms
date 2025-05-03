# Project tms

This project is a basic transaction management system built by using the blueprint provided by
https://github.com/Melkeydev/go-blueprint
the app uses an in-memory locking system and queue to handle transactions.
In a real world app a locking system and a queue that supports a distributed architecture will be used such as redis but for ease of use in-memory was used

## Getting Started

clone the project and make sure you have Postgres installed. After that create the TMS database and use this command:
`migrate -path db/migrations -database "postgresql://postgres:[PASSWORD]@[HOST]:[PORT]/[DATABASE_NAME]]?sslmode=disable" -verbose up`
Before that you will need to install `golang-migrate` to run the command if it is not already installed

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
