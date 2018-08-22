# Example: WebWire Messenger

[![MIT Licence](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbeon/webwire-example-postboard)](https://goreportcard.com/report/github.com/qbeon/webwire-example-postboard)
[![Travis CI: build status](https://travis-ci.org/qbeon/webwire-example-postboard.svg?branch=master)](https://travis-ci.org/qbeon/webwire-example-postboard)
[![Coveralls: Test Coverage](https://coveralls.io/repos/github/qbeon/webwire-example-postboard/badge.svg?branch=master)](https://coveralls.io/github/qbeon/webwire-example-postboard?branch=master)

This example demonstrates a full-blown single-instance messenger API server
written in [Go](https://golang.org/) and powered by
the [WebWire](https://github.com/qbeon/webwire-go) websocket library.

It includes the following features:
- Modular API-server architecture
- Optional TLS encryption
- Automated testing
- Metrics (real-time statistics)
- Logging (configurable)
- Request-Reply
- Server-side signals (**TODO**)
- Authentication & Sessions
- Authorization & Permissions
- Password hashing
- Graceful shutdown
- Custom HTTP handlers alongside webwire
- [dep](https://golang.github.io/dep/) for dependency management
