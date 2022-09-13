# kitly

## Requirements

- Go (1.18+)
- Docker

## Setup

### Build, Run & Test

```bash
# From root of the file

make build

make run

make test
```

## Docker images [Link](https://hub.docker.com/r/aniruddhabasak/kitly)

```bash
# Pull image

docker pull aniruddhabasak/kitly:1.1.0
```

## Test

```bash
$ curl -X POST http://localhost:8989/short?url=github.com

"http://localhost:8989/Z2l0aHViLmNvbQ=="
```
