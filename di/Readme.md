# Container Package

`container` is a lightweight DI (Dependency Injection) container for Go, allowing you to register and retrieve type instances by their types. The package is useful for managing dependencies and centralized object initialization.

---

## Table of Contents

- [Types and Constants](#types-and-constants)
- [Variables](#variables)
- [Functions](#functions)
- [Usage Examples](#usage-examples)
- [Features](#features)

---

## Types and Constants

### `RegisterStatus`

`RegisterStatus` describes the registration status of an object in the container.

| Constant             | Value | Description |
|---------------------|-------|-------------|
| `RegisterUnknown`    | 0     | Unknown registration status |
| `RegisterCreate`     | 1     | Object created automatically by the container |
| `RegisterCreateInit` | 2     | Object created and initialized with a provided value |
| `RegisterExists`     | 3     | Object is already registered in the container |

```go
type RegisterStatus int
```

## Variables

### `providers`

A storage for all registered providers.

- **Key**: String representation of the object type.
- **Value**: Instance of the object.

```go
var providers map[any]any
```

## Functions

### `RegisterRef[T any](initType ...*T) (RegisterStatus, error)`

Registers a type `T` in the container.

**Behavior:**

- If the container is not initialized — returns `RegisterUnknown` and an error.
- If the object is already registered — returns `RegisterExists`.
- If an instance of the object is provided — registers it and returns `RegisterCreateInit`.
- If no instance is provided — creates a new object using `new(T)` and returns `RegisterCreate`.

```go
func RegisterRef[T any](initType ...*T) (RegisterStatus, error)
```

### `AssignRef[T any]() (*T, error)`

Returns a reference to a registered object of type `T`.

**Errors:**

- If the object is not found — returns an error.
- If the object is found but its type does not match `T` — returns an error.

```go
func AssignRef[T any]() (*T, error)
```

## Usage Examples

```go
package main

import (
    "fmt"
    "log"
    "github.com/SeRj-ThuramS/go-container/di"
)

type MyService struct {
    Name string
}

func main() {
    // Register without passing an instance
    status, err := di.RegisterRef[MyService]()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Status:", status) // Output: 1 (RegisterCreate)

    // Register with an instance
    svc := &MyService{Name: "Service1"}
    status, err = di.RegisterRef(svc)
    fmt.Println("Status:", status) // Output: 2 (RegisterCreateInit)

    // Retrieve the object
    service, err := di.AssignRef[MyService]()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Service Name:", service.Name) // Output: Service1
}
```

## Features

- **Type-Safe Access**: Uses generics to ensure type-safe registration and retrieval.
- **Automatic Object Creation**: Automatically creates objects if no instance is provided.
- **Lightweight and Simple**: Suitable for projects and microservices that require a simple DI container.
- **Unique Type Registration**: Registers and retrieves types by their string representation, avoiding naming conflicts.
