# Go Load Balancer

## Overview

This project implements a simple load balancer in Go. It supports multiple backend servers and uses a round-robin algorithm to distribute incoming requests among them.

## Features

- Load balancing across multiple backend servers.
- Configurable server settings via a YAML configuration file.
- Health checks for backend servers.
- Simple HTTP server to handle incoming requests.

## Directory Structure

go-load-balancer/
├── cmd/ # Command-line applications
│ ├── main.go # Main application entry point
│ └── servers.go # Separate server functionality
├── configs/ # Configuration files
│ └── config.yaml # YAML configuration for the load balancer
├── internal/ # Internal packages
│ ├── algorithms/ # Load balancing algorithms
│ ├── backend/ # Backend server management
│ ├── health/ # Health check functionality
│ └── pool/ # Connection pool management
├── types/ # Type definitions and interfaces
├── go.mod # Go module file
└── go.sum # Go module dependencies

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
  port: 8084
  read_timeout: 10s
  write_timeout: 10s

backends:
  - url: http://localhost:9080
  - url: http://localhost:9001
  - url: http://localhost:9002

load_balancer:
  algorithm: "round-robin"
```

## Running the Application

To start the load balancer, run the following command:

```bash
go run ./cmd/main.go
```

To start the backend servers separately, run:

```bash
go run ./servers.go
```
