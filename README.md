# Gonca (GraphQL API Boilerplate)

![Docker](https://github.com/erdaltsksn/gonca/workflows/Docker/badge.svg)
[![Postman](https://img.shields.io/badge/Postman-reference-orange)](https://documenter.getpostman.com/view/5671920/TVRg8Vb7)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/erdaltsksn/gonca)](https://pkg.go.dev/github.com/erdaltsksn/gonca)
![Go](https://github.com/erdaltsksn/gonca/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/gonca)](https://goreportcard.com/report/github.com/erdaltsksn/gonca)
![CodeQL](https://github.com/erdaltsksn/gonca/workflows/CodeQL/badge.svg)

Gonca is an extendible GraphQL API boilerplate aiming to follow idiomatic go and
best practice.

## Features

- [Docker](https://www.docker.com) and [Docker Compose](https://github.com/docker/compose)
  based devops environment
- GraphQL API with [99designs/gqlgen](https://github.com/99designs/gqlgen)

## Requirements

- [Docker](https://www.docker.com)
- [Docker Compose](https://github.com/docker/compose)

## Getting Started

```sh
# Obtain the code
git clone https://github.com/erdaltsksn/gonca.git

# Get in the directory
cd gonca

# Run the application
docker-compose up --build
```

You can visit the GraphQL Playground at http://localhost:4000/

## Installation

TODO: Need documentation

## Updating / Upgrading

TODO: Need documentation

## Usage

You can use any method describe below after running the application via
`docker-compose up`.

### Via `Curl`

```sh
curl --request POST \
  --url 'http://localhost:4000/query?=' \
  --header 'content-type: application/json' \
  --data '{"query":"{\n  ping {\n    message\n  }\n}"}'
```

### Via `Axios`

```javascript
import axios from "axios";

const options = {
  method: "POST",
  url: "http://localhost:4000/query",
  params: { "": "" },
  headers: { "content-type": "application/json" },
  data: '{"query":"{\n  ping {\n    message\n  }\n}"}',
};

axios
  .request(options)
  .then(function (response) {
    console.log(response.data);
  })
  .catch(function (error) {
    console.error(error);
  });
```

### Via `Python`

```python
import http.client

conn = http.client.HTTPConnection("localhost:4000")

payload = "{\"query\":\"{\\n  ping {\\n    message\\n  }\\n}\"}"

headers = { 'content-type': "application/json" }

conn.request("POST", "/query?=", payload, headers)

res = conn.getresponse()
data = res.read()

print(data.decode("utf-8"))
```

## Getting Help

You can visit [Gonca's Postman Collection](https://documenter.getpostman.com/view/5671920/Szf3Z9Zy)
for API refence and examples.

```sh
# Show available `make` commands.
make help
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](.github/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](.github/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](.github/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
