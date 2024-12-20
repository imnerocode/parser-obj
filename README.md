# OBJ Parser

This package provides a function to parse OBJ files and convert them into a model structure using the `vo-structures` package.

## Installation

To install this package, use `go get`:

```sh
go get github.com/imnerocode/parser-obj
```

## Usage

Below is an example of how to use this package to parse an OBJ file:

```go
package main

import (
    "fmt"
    "log"

    "github.com/imnerocode/parser-obj"
)

func main() {
    model, err := parser_obj.ParseOBJ("path/to/your/file.obj")
    if err != nil {
        log.Fatalf("Error parsing OBJ file: %v", err)
    }

    fmt.Printf("Parsed model: %+v\n", model)
}
```

## Functions

### ParseOBJ

```go
func ParseOBJ(filePath string) (*vo.Model, error)
```

Parses an OBJ file and converts it into a `vo.Model` structure.

- **filePath**: The path to the OBJ file.
- Returns a pointer to `vo.Model` and an error if any issue occurs.

## Structures

This package uses the following structures from the `vo-structures` package:

- `vo.Model`
- `vo.Vertex`
- `vo.Face`

You need to import these structures from the `vo-structures` package, which is a separate package available at: [https://github.com/imnerocode/vo-structures](https://github.com/imnerocode/vo-structures)

## License

This project is licensed under the terms of the MIT license.