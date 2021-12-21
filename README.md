# Bank Rest API
This codebase was created to demonstrate a backend application built with Golang including CRUD operations, routing, pagination, and more.

# Directory Structure
```markdown
┌── api
│   └── swagger.yml         // API Spec
├── cmd     
│   └── server              
│       └── main.go         // Server Start  
├── internal                // Project internal files
│   └── app     
│       ├──  api.go         // Business Logic & Router
│       └──  api_test.go    // Business Logic Tests
└── Dockerfile              // Dockerfile for building an image
```
# Getting started

## Install Golang

Make sure you have Go 1.17 or higher installed.

https://golang.org/doc/install

## Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).


## Install Dependencies
From the project root, run:
```
go build ./...
go test ./...
go mod tidy
```

## Testing
From the project root, run:
```
go test ./...
```
or
```
go test ./... -cover
```
or
```
go test -v ./... -cover
```
depending on whether you want to see test coverage and how verbose the output you want.

## Run
```
go run ./...
```

## APIs

