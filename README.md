# Demo Distributed Wallet

Demo of a distributed wallet service, implemented in Golang using Protocol Buffers, gRPC and PostreSQL.

- horizontally scalable
- double spend protection
- overdraft protection

## Run in Docker

This will spin up the database, build everything, execute all the tests and expose the service on `$RPC_PORT` (`8080`).

```
$ go install github.com/go-task/task/v3/cmd/task@latest
$ task docker-start
```

One may want to `task clean` afterwards to remove persitent cache volumes.

## Run locally

```
$ go install github.com/go-task/task/v3/cmd/task@latest
$ task docker-start-db
$ task test
$ task start
```

## Taskfile

```
$ task --list-all
task: Available tasks for this project:
* build:
* build-protobufs:
* clean:
* demo:
* docker-rebuild-start:
* docker-start:
* docker-start-db:
* docker-stop:
* install-deps:
* kill:
* start:
* test:
```

## Raw notes

```
- golang wallet microservice
    - create wallet
    - alter balance
    - get balance
- grpc for communication
- postgres for persistence and locks
    - balances
    - ledger
- docker compose for the env
- unit tests
- acceptance tests for concurrency
- bugs
    - SIGTERM in docker compose
- not covered
    - timeouts, timezones, currencies
    - auth, validation
    - dependency injection
    - load balancer
    - empty test DB
    - multiple wallet instances (k8s)
    - postgres cluster
    - metrics
    - proper RPC errors
    - more tests
    - benchmarks
    - code comments ;]
- scaling up
    - request / response queues
    - locks via a key-value store (eg redis sentinel)
    - ledger via a timeseries db
    - balance recalculation only when subtracting
    - manual rollbacks
```
