# copytree

[![Go Report Card](https://goreportcard.com/badge/github.com/baohuiming/copytree)](https://goreportcard.com/report/github.com/baohuiming/copytree)

## Introduction

**copytree** is a Go package that provides a convenient way to move directories or files, with support for recursion and moving to different drives. It is similar to the `shutil.copytree` function in Python.

## Installation

To use copytree in your Go project, simply run the following command:

```shell
go get github.com/baohuiming/copytree
```

## Usage

Import the copytree package in your Go code:

```go
import "github.com/baohuiming/copytree"
```

Use the `MoveFiles` function to move directories or files:

```go
err := copytree.MoveFiles(sourcePath, destinationPath)
if err != nil {
    log.Fatal(err)
}
```

The `MoveFiles` function supports recursion, meaning it will move all the files and subdirectories within the source directory. It also allows you to move the files to a different drive by specifying the destination path.

If you only want to copy the files instead of moving them, you can use the `CopyFiles` function:

```go
err := copytree.CopyFiles(sourcePath, destinationPath)
if err != nil {
    log.Fatal(err)
}
```

## Example

Here's an example that demonstrates how to use copytree to move a directory:

```go
package main

import (
    "log"

    "github.com/baohuiming/copytree"
)

func main() {
    err := copytree.MoveFiles("/path/to/source", "/path/to/destination")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/baohuiming/copytree).

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).