
# Challenge

A single facts rest endpoint secured by auth key  ready to be integrated with other services

## Table of Content
- [Features](#features)
- [Project Structure](#project-structure)
- [Environment Variables](#environment-variables)
- [Run Locally](#run-locally)
- [Running Tests](#running-tests)
- [API Reference](#api-reference)
  - [Get facts](#get-facts)
  - [Documentation](#documentation)

## Features

- Seed facts and keys collections by `db.json` and `keys.json` data during the start of MongoDB through docker-compose `seed.sh` file call
- Repository-Service pattern breaks up the business layer of the app into two layers to give more flexibility to any new change
- Start secured server only to make sure that all interactions guarded by SSL
- Key-based authentication to allow only authorized devices to call the API

## Project Structure
```
.
├── controllers
├── docker
│   ├── cert
│   └── mongo
├── integration
├── main.go
├── models
├── pkg
│   ├── config
│   └── db
├── repositories
├── server
└── services
```
  
## Environment Variables

To run this project, you will need to create `env.yaml` file and add the following environment variables to it

```yaml
app_port: 

db:
  host: ""
  port:
  database: ""
  username: ""
  password: ""

certificate:
  cert_file: ""
  key_file: ""

```

  
## Run Locally

Go to the project directory

```bash
  cd challenge
```

Create `env.yaml` file and add envs to it

```yaml
debug: true
app_port: 1323

db:
  host: "mongo"
  port: 27017
  database: "test"
  username: "root"
  password: "secret"

certificate:
  cert_file: "docker/cert/cert.pem"
  key_file: "docker/cert/key.pem"
```

Add `docker/cert/cert.pem` to root certificate or generate new certificate and use it instead <br><br>



Start MongoDB and Go server  using Docker Compose

```bash
  docker-compose up
```
  
## Running Tests

Go to the project directory

```bash
  cd challenge
```

Install dependencies

```bash
  go get
```

Start MongoDB using Docker Compose

```bash
  docker-compose up mongo
```

Create `env.test.yaml` file and add envs to it
```yaml
debug: true
app_port: 1323

db:
  host: "localhost"
  port: 27017
  database: "test"
  username: "root"
  password: "secret"

certificate:
  cert_file: "../docker/cert/cert.pem"
  key_file: "../docker/cert/key.pem"
```
Add `docker/cert/cert.pem` to root certificate or generate new certificate and use it instead <br> <br>

Run tests

```bash
  go test ./integration -v
```

## API Reference

### Get facts

Return facts about numbers

```http
  GET /api/v1/facts
```

| Parameter  | Type     | Description                                     |
| :--------- | :------- | :---------------------------------------------- |
| `s`        | `string` | **Optional**. Search word or sentence           |
| `page`     | `number` | **Optional**. Page number, default 1            |
| `per_page` | `number` | **Optional**. Returned items number, default 10 |

| Headers         | Type     | Description                |
| :-------------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**. Your API key |
  
### Documentation
 
[Facts API Documentation](https://documenter.getpostman.com/view/619668/Tzef9Nht) 

Please use one of the keys on the `docker/mongo/keys.josn` file to test API, for production-ready we could create Auth service which going to be responsible for generating Authentication key for each device call facts API
