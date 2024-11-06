
# Go Rest API Documentation

## Table Of Content
- [Prerequisite](#prerequisites)
  - [Structure](#structure)
- [How To](#how-to)
- [References](#references)
  - [Architecture](#architecture)


## Prerequisites
What things you need to setup the application:
- [Docker](https://www.docker.com/)
- [Golang](https://go.dev/doc/install)
- Makefile
- MySQL

### Structure
```
.
└── cmd
|   └── main.go
├── pkg
├── internal
|   ├── database
|   |	├── migrations
|	|		└── *.go
|   ├── domain
|   │   └── entities
|   │       └── *.go
| 	| 	└── services
|   │       └── *.go
|   ├── repository
|   │   └── *.go
|   └── usecase
|       └── *.go
|   ├── interfaces
|   │   └── *.go
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## How To
### Running The App (local)
- First get the dependencies with this command:
```shell
make mod
```

- and for running the application can use this command:
```shell
make run
```

## References
### GIT Style
For commit message style or git style guide, use this doc
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)


### Architecture
Architecture reference on go-clean-arch, follow this doc repo
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
