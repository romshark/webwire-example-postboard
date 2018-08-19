# Example: WebWire Messenger

This example demonstrates a full-blown single-instance messenger API server
written in [Go](https://golang.org/) and powered by
the [WebWire](https://github.com/qbeon/webwire-go) websocket library.

It includes the following features:
- Modular API-server architecture
- Optional TLS encryption
- Automated testing (**TODO**)
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