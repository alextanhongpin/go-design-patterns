# Testing

## Run all test

```bash
$ go test -v ./...
```

## Run specific test with regex

```bash
$ go test -v -run=Builder ./...
```

## Check coverage

```bash
$ go test -v ./abstract-factory -cover
```
