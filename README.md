# tsr-bootcamp-proxy

Simple proxy server for providing data to [TSR's angular page(s)](https://github.com/cbdr/tsr_bootcamp)

## Getting Started

1. Clone the repo
2. Start the server: `go run main.go`

## Organization

Currently the application is split into several packages. Here's an overview of their responsibilities.

- `package main`: sets up routing, configuration, and starts the server.
- `package handler`: `http.Handler`s that handle requests, strip query params, and call the appropriate controller package.
- `package *`: all other packages simply provide basic functions for retrieving the data necessary. They **should be http independent** and simply return the data (or an error...)

E.g: To add a `/user` endpoint you would do the following:

- add the file `user.go` to  `/handler` folder. This file must export a `http.Handler` that will handle requests to the endpoint. That handler should then delegate any of the "heavy lifting" to a new `user` package.
- add the `user` package (in a `/user` folder). This package is responsible for retrieving the user data from its source.
- add the `/user` route in `main.go`
