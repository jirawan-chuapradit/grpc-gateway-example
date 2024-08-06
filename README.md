# gRPC-Gateway Example with DataDog Tracing

This project demonstrates how to set up a gRPC service with a RESTful HTTP API using the gRPC-Gateway, integrated with DataDog tracing for observability.

## Overview

- The gRPC-Gateway translates RESTful HTTP API calls into gRPC calls, allowing you to expose your gRPC service to HTTP clients.
- DataDog tracing is implemented to monitor and analyze the performance of both the gRPC and HTTP services.

## Tracing Results

### Basic Tracing
![Basic Tracing Result](/img/basic_tracing_result.png)

### DataDog Tracing for HTTP Client
![DataDog Tracing Result](/img/http_tracing_result.png)

## Quick Start

Follow these steps to run the demo:

1. Start the gRPC-Gateway:
```
    $ make start-gateway
```

2. Start the gRPC server:
```
    $ make start-grpc
```

3. Run the client:
```
    $ make start-client
```

4. View the results in your DataDog dashboard

## Project Structure

- `cmd/gateway`: Contains the gRPC-Gateway implementation
- `cmd/grpc`: Houses the gRPC server code
- `cmd/client`: Includes the HTTP client with DataDog tracing

## Features

- gRPC service implementation
- RESTful HTTP API via gRPC-Gateway
- DataDog tracing integration for performance monitoring
- Makefile for easy project management

## Requirements

- Go 1.15+
- Protocol Buffers
- gRPC-Gateway
- DataDog agent (for tracing)


