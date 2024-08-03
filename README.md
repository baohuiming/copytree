# mvdirs

[![Go Report Card](https://goreportcard.com/badge/github.com/baohuiming/mvdirs)](https://goreportcard.com/report/github.com/baohuiming/mvdirs)

## Introduction

mvdirs is a Go package that provides a convenient way to move directories or files, with support for recursion and moving to different drives.

## Installation

To use mvdirs in your Go project, simply run the following command:

```shell
go get github.com/baohuiming/mvdirs
```

## Usage

Import the mvdirs package in your Go code:

```go
import "github.com/baohuiming/mvdirs"
```

Use the `MoveFiles` function to move directories or files:

```go
err := mvdirs.MoveFiles(sourcePath, destinationPath)
if err != nil {
    log.Fatal(err)
}
```

The `MoveFiles` function supports recursion, meaning it will move all the files and subdirectories within the source directory. It also allows you to move the files to a different drive by specifying the destination path.

## Example

Here's an example that demonstrates how to use mvdirs to move a directory:

```go
package main

import (
    "log"

    "github.com/baohuiming/mvdirs"
)

func main() {
    err := mvdirs.MoveFiles("/path/to/source", "/path/to/destination")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/baohuiming/mvdirs).

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).