# Go Load Balancer

## Overview

This project implements a simple load balancer in Go. It supports multiple backend servers and uses a round-robin algorithm to distribute incoming requests among them.

## Features

- Load balancing across multiple backend servers.
- Configurable server settings via a YAML configuration file.
- Health checks for backend servers.
- Simple HTTP server to handle incoming requests.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-load-balancer.git
   cd go-load-balancer
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

The load balancer configuration is defined in `configs/config.yaml`. You can specify the following settings:

```yaml
server:
  port: 8080
  read_timeout: 10s
  write_timeout: 10s

backends:
  - url: http://servers:9080
  - url: http://servers:9001
  - url: http://servers:9002

load_balancer:
  algorithm: "round-robin"
```

## Running the Application

To start the load balancer, run the following command:

```bash
go run cmd/main.go
```

To start the backend servers separately, run:

```bash
go run servers/servers.go
```

## Using Docker Compose

Build and start all services:

```bash
docker compose up
```
