# Govdev

A highly configurable and scalable backend for managing user services.
Originally built to serve UCLA DevX's application portal needs. This document
serves as the technical specification of how the software is designed and how it
is supposed to work.
- [Govdev](#govdev)
  * [Languages and Frameworks](#languages-and-frameworks)
    + [Go](#go)
    + [Echo](#echo)
  * [External Dependencies](#external-dependencies)
    + [Database](#database)
    + [Cache](#cache)
    + [Message Broker](#message-broker)
    + [File/Object Storage](#file-object-storage)
    + [Proxy Server](#proxy-server)
    + [Containers and Orchestration](#containers-and-orchestration)
  * [Dependency Graph](#dependency-graph)
    + [Stores](#stores)
    + [Services](#services)
    + [Adapters](#adapters)
  * [Entities and Relationships](#entities-and-relationships)
    + [Users](#users)
    + [Roles](#roles)
    + [Applications](#applications)
    + [Scores](#scores)
    + [Teams](#teams)
  * [Relationships](#relationships)
    + [Users and Applications](#users-and-applications)
    + [Users and Roles](#users-and-roles)
    + [Applications and Scores](#applications-and-scores)
    + [Users and teams](#users-and-teams)
  * [Authentication and Roles](#authentication-and-roles)
    + [Sessions](#sessions)
    + [Permissions](#permissions)
  * [External HTTP API](#external-http-api)
  * [Misc](#misc)
    + [Error Handling](#error-handling)
    + [Logging](#logging)

## Languages and Frameworks

### Go

Go is an awesome language and everyone should use Go for their projects.

### Echo
`github.com/labstack/echo` is our framework of choice to serve requests. It is
lightweight and fast, and offers middleware and logging options that are
conviennt.

## External Dependencies

### Database

This pertains to relational databases, though through abstractions, we can
technically support NoSQL databases, but it would be quite a hassle.

* PosgreSQL (selected)
* Oracle (who wants to pay Oracle millions?)
* MySQL

### Cache

A cache service using a key-value system, with fast retrievel. While this cache
should operate mostly in memory, in can be dumped and reread later if needed.
Several options include:

* Redis (selected)
* memcached
* MongoDB
* freecache

### Message Broker

Protocol that translates a message to a formal messaging protocol of the sender,
used to process many messages and offload messaging to another service.

* Redis (selected because cache was already Redis)
* NATS
* Kafka
* RAbbitMQ

### File/Object Storage

Service that stores files and objects, and can be retrieved by a unique URL.

* S3 and cloud equivalents
* Minio (S3 compatible but self hosted, selected for this reason)

### Proxy Server

* Caddy
* NginX
* Apache
* Traefik

### Containers and Orchestration

While not directly related to our services, our services will run in containers
and so affects our runtime. Our services will run in Docker containers, ideally
under the Alpine distribution (or scratch). Docker-compose will string our
services together through networks. It should be noted that the above services
will also run in containers.

* Docker and Docker-compose (selected)
* Kubernetes (Awesome tech, but totally unnecessary)

## Dependency Graph

Everything in our application dependes on interfaces to interact. For example, a
HTTP handler will not make direct calls to a database; instead, the database is
abstracted away with a stores object, a services object handles all the business
logic, and the HTTP layer handles inputs and outputs. In this fashion, if we
were to change an underlying implementation of a store or service, higher up
layers are not affected.

### Stores

Stores wrap an interface around an external dependency, such as an interface
around Redis, so that even if we change our cache service, we do not have to
change every reference to the cache. We can swap out Redis for say memcached or
something and implement the interface, and the application should work. If we
wanted to use a hosted PostgreSQL or similiar database, we can simply swap out
the link and the rest of our codebase would not be affected by the swap.

### Services

Services define a group of related actions. Services take in stores via
interfaces, and use these stores to do business logic. For example, we might
have a users store that deals with all actions a user might need to do. Services
use stores to persist data or make calls to stores to retrieve data. 

### Adapters

Adapters do not handle business logic, but instead handle incoming data from a
endpoint, decodes the inputs, calls on a service or services to compute some
actions, then sends back the outputs. An adapter should not deal with any
business logic, only with I/O.

The reason that adapters are implemented separately from services is that one
day, we may need to support GraphQL, XML, gRPC and Protobuffers, or some other 
communication protocol that is not standard HTTP. In that fashion, we can simply
reconstruct the code needed to decode inputs, pass it into services, and return
the outputs, without affecting other parts of our codebase.

## Entities and Relationships

### Users

### Roles

### Applications

### Scores

### Teams

## Relationships

### Users and Applications

### Users and Roles

### Applications and Scores

### Users and teams

## Authentication and Roles

### Sessions

### Permissions 

## External HTTP API

This section defines the external facing API.

## Misc

### Error Handling

### Logging
