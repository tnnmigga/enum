# Enum Library

This Go library provides a utility for creating and managing enums in a structured and type-safe manner. The library supports both string and integer enums, as well as nested enums, with additional functionality to query enum fields and values.

## Features

- **String Enums**: Automatically initializes string fields with their field names.
- **Integer Enums**: Supports custom integer values using struct tags.
- **Nested Enums**: Allows defining enums with nested structures.
- **Field Checking**: Check if a top-level field exists with a specific value using `Contains`.
- **Field Listing**: Retrieve names of top-level fields using `Keys`.
- **Value Listing**: Retrieve values of top-level fields (string or integer) using `Values`.

## Installation

To install the library, use:

```bash
go get github.com/tnnmigga/enum@latest
```

## Usage

### String Enums

```go
var HttpStatus = New[struct {
    StatusOK                  string
    StatusNotFound            string
    StatusInternalServerError string
}]()

fmt.Println(HttpStatus.StatusOK) // Output: "StatusOK"
fmt.Println(enum.Contains(HttpStatus, "StatusOK")) // Output: true
fmt.Println(enum.Keys(HttpStatus)) // Output: [StatusOK StatusNotFound StatusInternalServerError]
fmt.Println(enum.Values[string](HttpStatus)) // Output: [StatusOK StatusNotFound StatusInternalServerError]
```

### Integer Enums

```go
var HttpStatus = New[struct {
    StatusOK                  int `enum:"200"`
    StatusNotFound            int `enum:"404"`
    StatusInternalServerError int `enum:"500"`
}]()

fmt.Println(HttpStatus.StatusOK) // Output: 200
fmt.Println(enum.Contains(HttpStatus, 200)) // Output: true
fmt.Println(enum.Keys(HttpStatus)) // Output: [StatusOK StatusNotFound StatusInternalServerError]
fmt.Println(enum.Values[int](HttpStatus)) // Output: [200 404 500]
```

### Nested Enums

```go
var HttpStatus = New[struct {
    Code struct {
        StatusOK                  int `enum:"200"`
        StatusNotFound            int `enum:"404"`
        StatusInternalServerError int `enum:"500"`
    }
    Type struct {
        StatusOK                  string
        StatusNotFound            string
        StatusInternalServerError string
    }
}]()

fmt.Println(HttpStatus.Code.StatusOK) // Output: 200
fmt.Println(HttpStatus.Type.StatusOK) // Output: "StatusOK"
fmt.Println(enum.Contains(HttpStatus, HttpStatus.Code)) // Output: true
fmt.Println(enum.Keys(HttpStatus)) // Output: [Code Type]
fmt.Println(enum.Values[string](HttpStatus)) // Output: []
```

## Testing

The library includes comprehensive tests for initializing enums and verifying the `Contains`, `Keys`, and `Values` functions. Run the tests using:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.