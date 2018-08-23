# Postboard - the real-time example application

[![MIT Licence](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/qbeon/webwire-example-postboard)](https://goreportcard.com/report/github.com/qbeon/webwire-example-postboard)
[![Travis CI: build status](https://travis-ci.org/qbeon/webwire-example-postboard.svg?branch=master)](https://travis-ci.org/qbeon/webwire-example-postboard)
[![Coveralls: Test Coverage](https://coveralls.io/repos/github/qbeon/webwire-example-postboard/badge.svg?branch=master)](https://coveralls.io/github/qbeon/webwire-example-postboard?branch=master)

This example demonstrates a full-blown single-instance real-time API server
written in [Go](https://golang.org/) and powered by
the [WebWire](https://github.com/qbeon/webwire-go) websocket library.

It includes the following features:
- [A modular API-server architecture](#api-server-architecture)
- Optional TLS encryption
- Automated testing
- Metrics (real-time statistics)
- Logging (configurable)
- Request-Reply
- Server-side signals & subscriptions
- Authentication & Sessions
- Authorization & Permissions
- Password hashing
- Graceful shutdown
- Custom HTTP handlers alongside webwire
- [dep](https://golang.github.io/dep/) for dependency management

## API Server Architecture
The API server's architecture is based on a 3-stage request processing pipeline
where the last 2 pipeline segments implement an abstract interface.
<br>
<br>
![API server request processing pipeline stages](https://cdn.rawgit.com/qbeon/webwire-example-postboard/b693c352/doc/request-processing-pipeline.svg)

An incoming request is received by the network layer (implemented by
[webwire-go](https://github.com/qbeon/webwire-go)) where it's decoded,
unmarshalled, associated with a client session and finally passed over to
the stage 2 [resolver interface](https://github.com/qbeon/webwire-example-postboard/blob/master/server/apisrv/modules/resolver/interface.go).
The [resolver](https://github.com/qbeon/webwire-example-postboard/tree/master/server/apisrv/modules/resolver)
implementation validates the passed parameters and authorizes the request
by checking the provided session's permissions. If the request passes the
validation and authorization procedure the resolver resolves it into
one or more engine calls through the [stage 3 engine interface](https://github.com/qbeon/webwire-example-postboard/blob/master/server/apisrv/modules/engine/engine.go).

The resolver can be made responsible for optimizing engine calls (to reduce
database load for example) by using the [batching](https://en.wikipedia.org/wiki/Batch_processing),
[caching](https://en.wikipedia.org/wiki/Cache_(computing)) and [deduplication](https://en.wikipedia.org/wiki/Data_deduplication) techniques,
so that instead of directly performing an engine call
to retrieve the data of an *entity* to fulfill a *request* for example
the resolver could use a loader module and buffer multiple requests
of the same type while simultaneously deduplicating similar requests
to eventually perform a batched engine call when either
the loader's buffer limit is exceeded or the loader's interval tick is reached.

Finally the engine calls are executed by the [memeng](https://github.com/qbeon/webwire-example-postboard/tree/master/server/apisrv/modules/engine/memeng) in-memory engine implementation, which takes care of atomically mutating and/or reading
the data from it's in-memory store.
It is to be noted that engine calls are considered [atomic](https://en.wikipedia.org/wiki/Atomicity_(database_systems)) and are thus performed in an all-or-nothing fashion while
transactions spanning over multiple engine calls are not possible.

The goal of this architecture is to abstract away the actual implementation
of the individual system modules isolating responsibilities improving
both security and maintainability of the system. The engine can,
for example, be re-implemented using either a different store backed by a DBMS,
noSQL database, filesystem etc. or even [remote microservices](#microservices).
Theoretically multiple implementations of the engine can co-exist,
an in-memory engine could, for example, be used to quickly mockup
a prototype and develop API tests before
an actual, more complex, persistent database-based engine is implemented.


## Scalability
While this particular example demonstrates a single-instance service only,
it is possible to make it horizontally scalable using both the [clustering](https://en.wikipedia.org/wiki/Computer_cluster) and [microservice](https://en.wikipedia.org/wiki/Microservices)
techniques.

### Clustering
To improve performance and [fault-tolerance](https://en.wikipedia.org/wiki/Fault_tolerance)
this API service can be replicated onto multiple machines, which will require
a new module for the communication between the individual cluster nodes.
This module would be responsible for:
- synchronizing the nodes,
- broadcasting (and receiving) events for the individual nodes to be able
  to notify their connected clients,
- and invalidating local caches if necessary.

### Microservices
This architectural approach allows to ease the process of turning an initially
monolithic API service into a microservice-based distributed system by gradually
reimplementing the engine to offload certain functionalities and datasets
to individual microservices by redirecting the data mutation and retrieval logic
onto it.

Eventually the API service will turn into a [Microservice API Gateway](https://dzone.com/articles/microservice-pattern-api-gateway)
which is then only responsible for validating, authorizing and routing requests,
as well as acting as a bridge for signals notifying connected clients
about server-side events.
The implementation of the logic behind the engine calls in this case is
offloaded to the individual remote microservices.
