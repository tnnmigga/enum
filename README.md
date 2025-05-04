# Enum Library

This Go library provides a utility for creating and managing enums in a structured and type-safe manner. The library supports both string and integer enums, as well as nested enums.

## Features

- **String Enums**: Automatically initializes string fields with their field names.
- **Integer Enums**: Supports custom integer values using struct tags.
- **Nested Enums**: Allows defining enums with nested structures.

## Installation

To install the library, use:

```bash
go get github.com/tnnmigga/enum@latest
```

## Usage

### String Enums

```go
HttpStatus := New[struct {
	StatusOK                  string
	StatusNotFound            string
	StatusInternalServerError string
}]()

fmt.Println(HttpStatus.StatusOK) // Output: "StatusOK"
```

### Integer Enums

```go
HttpStatus := New[struct {
	StatusOK                  int `enum:"200"`
	StatusNotFound            int `enum:"404"`
	StatusInternalServerError int `enum:"500"`
}]()

fmt.Println(HttpStatus.StatusOK) // Output: 200
```

### Nested Enums

```go
HttpStatus := New[struct {
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
```

## Testing

The library includes comprehensive tests to ensure its functionality. Run the tests using:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.  