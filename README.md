# Web Broker
  FIFO web broker with file storage

## Usage:

```
PUT http://localhost:8000/queue?v=task
PUT http://localhost:8000/color?v=red
```

```
GET http://localhost:8000/queue
```

## Run app

```
make build && make run
```

## Test app

```
make test
```
