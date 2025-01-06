# Archivarius

**Archivarius** is a high-performance gRPC-based service written in Go, designed for handling and retrieving statistical data with extensive filtering, ordering, and grouping capabilities. Whether you're dealing with large datasets or require precise data manipulation, Archivarius provides a robust and scalable solution to meet your needs.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Generating gRPC Code](#generating-grpc-code)
- [Usage](#usage)
  - [Server](#server)
  - [Client](#client)
- [API Reference](#api-reference)
  - [Protocol Buffers](#protocol-buffers)
  - [Client Methods](#client-methods)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Advanced Filtering**: Apply complex filter conditions on various keys to retrieve precise data.
- **Flexible Ordering**: Order results based on multiple criteria in ascending or descending order.
- **Grouping & Pagination**: Group data by specified keys and handle large datasets with pagination support.
- **Extensible Schema**: Easily extendable Protocol Buffers schema to accommodate additional requirements.
- **Secure & Efficient**: Leverages gRPC for high-performance communication with optional secure connections.

## Architecture

Archivarius is built using Go and gRPC, following a client-server architecture. The server exposes a `Statistic` RPC method that accepts a `StatisticRequest` and returns a `StatisticResponse`. The client interacts with the server using the generated gRPC client code.

## Getting Started

### Prerequisites

- **Go**: Version 1.18 or higher. [Download Go](https://golang.org/dl/)
- **Protocol Buffers Compiler (`protoc`)**: Version 3.11 or higher. [Installation Guide](https://grpc.io/docs/protoc-installation/)
- **Go Modules**: Ensure Go modules are enabled.

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/archivarius.git
   cd archivarius
   ```

2. **Install Dependencies**

   ```bash
   go mod download
   ```

### Generating gRPC Code

The Protocol Buffers definitions are located in the `proto/` directory. To generate the Go code for gRPC:

```bash
protoc --go_out=. --go-grpc_out=. proto/archivarius.proto
```

This will generate the necessary `.pb.go` files in the `grpc/` directory as specified by the `go_package` option.

## Usage

### Server

To run the Archivarius gRPC server:

1. **Navigate to the Server Directory**

   ```bash
   cd server
   ```

2. **Build and Run the Server**

   ```bash
   go build -o archivarius-server
   ./archivarius-server
   ```

   The server will start listening on the default gRPC port (e.g., `:50051`).

### Client

The client provides a convenient way to interact with the Archivarius service.

1. **Import the Client Package**

   ```go
   import "github.com/yourusername/archivarius/client"
   ```

2. **Connect to the Server**

   ```go
   ctx := context.Background()
   apiClient, err := client.ConnectAPI(ctx, "localhost:50051")
   if err != nil {
       log.Fatalf("Failed to connect: %v", err)
   }
   defer apiClient.Close()
   ```

3. **Make a `Statistic` Request**

   ```go
   req := &client.StatisticRequest{
       // Populate request fields
   }
   resp, err := apiClient.Statistic(ctx, req)
   if err != nil {
       log.Fatalf("Statistic RPC failed: %v", err)
   }
   fmt.Printf("Total Count: %d\n", resp.TotalCount)
   for _, item := range resp.Items {
       // Process each item
   }
   ```

## API Reference

### Protocol Buffers

The core definitions are in `proto/archivarius.proto`. Below are key components:

- **Enums**
  - `OrderingKey`: Defines keys for ordering results.
  - `Key`: Defines various keys used in filtering and grouping.
  - `Condition`: Specifies filter operations like EQ, NE, GT, etc.

- **Messages**
  - `ItemKey`: Represents a key-value pair.
  - `Value`: Holds different types of values (string, int, float, timestamp, IP).
  - `Item`: Contains statistical metrics and associated keys.
  - `FilterCondition`: Defines a single filter condition.
  - `Filter`: Aggregates multiple filter conditions with a date range.
  - `Order`: Specifies ordering preferences.
  - `StatisticRequest`: Request structure for the `Statistic` RPC.
  - `StatisticResponse`: Response structure containing items and total count.

- **Service**
  - `Archivarius`: Defines the `Statistic` RPC method.

### Client Methods

The client package provides the following key methods:

- **ConnectAPI**

  ```go
  func ConnectAPI(ctx context.Context, connect string, opts ...grpc.DialOption) (*APIClient, error)
  ```

  Establishes a connection to the Archivarius server.

- **Statistic**

  ```go
  func (c *APIClient) Statistic(ctx context.Context, in *StatisticRequest, opts ...grpc.CallOption) (*StatisticResponse, error)
  ```

  Sends a `StatisticRequest` to the server and returns a `StatisticResponse`.

## Examples

### Example: Fetching Statistics with Filters and Ordering

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/yourusername/archivarius/client"
    "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
    ctx := context.Background()
    apiClient, err := client.ConnectAPI(ctx, "localhost:50051")
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer apiClient.Close()

    // Define filter conditions
    filter := &client.Filter{
        Conditions: []*client.FilterCondition{
            {
                Key: client.KEY_COUNTRY,
                Op:  client.Condition_EQ,
                Value: []*client.Value{
                    {Value: client.MakeExtValue("US")},
                },
            },
            {
                Key: client.KEY_DATE_MARK,
                Op:  client.Condition_GE,
                Value: []*client.Value{
                    {Value: client.MakeExtValue(time.Now())},
                },
            },
        },
        StartDate: timestamppb.New(time.Now().AddDate(0, -1, 0)),
        EndDate:   timestamppb.Now(),
    }

    // Define ordering
    order := []*client.Order{
        {
            Key: client.ORDERING_KEY_SPENT,
            Asc: false,
        },
    }

    // Create request
    req := &client.StatisticRequest{
        Filter:     filter,
        Order:      order,
        Group:      []client.Key{client.KEY_COUNTRY},
        PageOffset: 0,
        PageLimit:  10,
    }

    // Make RPC call
    resp, err := apiClient.Statistic(ctx, req)
    if err != nil {
        log.Fatalf("Statistic RPC failed: %v", err)
    }

    // Process response
    fmt.Printf("Total Count: %d\n", resp.TotalCount)
    for _, item := range resp.Items {
        fmt.Printf("Spent: %.2f, Profit: %.2f\n", item.Spent, item.Profit)
        for _, key := range item.Keys {
            fmt.Printf("Key: %s, Value: %v\n", key.Key, key.Value)
        }
    }
}
```

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork the Repository**

2. **Create a Feature Branch**

   ```bash
   git checkout -b feature/YourFeature
   ```

3. **Commit Your Changes**

   ```bash
   git commit -m "Add your feature"
   ```

4. **Push to the Branch**

   ```bash
   git push origin feature/YourFeature
   ```

5. **Open a Pull Request**

Please ensure your code adheres to the project's coding standards and includes appropriate tests.

## License

This project is licensed under the [Custom License](LICENSE.md).
