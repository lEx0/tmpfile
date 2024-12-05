Here is a `README.md` file for your GitHub repository:

# tmpfile

`tmpfile` is a Go package that provides a simple wrapper around temporary files. It allows you to create, read, write, seek, and automatically delete temporary files.

## Installation

To install the package, use the following command:

```sh
go get github.com/lEx0/tmpfile
```

## Usage

### Creating a Temporary File from a Reader

You can create a temporary file from an `io.Reader` using the `NewFromReader` function.

```go
package main

import (
    "strings"
    "github.com/lEx0/tmpfile"
)

func main() {
    f, err := tmpfile.NewFromReader(strings.NewReader("test"))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Use the temporary file
}
```

### Creating an Empty Temporary File

You can create an empty temporary file using the `New` function.

```go
package main

import (
    "github.com/lEx0/tmpfile"
)

func main() {
    f, err := tmpfile.New()
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Use the temporary file
}
```

### Reading from a Temporary File

You can read from a temporary file using the `Read` method.

```go
package main

import (
    "fmt"
    "io"
    "strings"
    "github.com/lEx0/tmpfile"
)

func main() {
    f, err := tmpfile.NewFromReader(strings.NewReader("test"))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    body, err := io.ReadAll(f)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(body)) // Output: test
}
```

### Writing to a Temporary File

You can write to a temporary file using the `Write` method.

```go
package main

import (
    "strings"
    "github.com/lEx0/tmpfile"
)

func main() {
    f, err := tmpfile.New()
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = f.Write([]byte("test"))
    if err != nil {
        panic(err)
    }
}
```

### Seeking in a Temporary File

You can seek to a specific position in a temporary file using the `Seek` method.

```go
package main

import (
    "fmt"
    "io"
    "strings"
    "github.com/lEx0/tmpfile"
)

func main() {
    f, err := tmpfile.NewFromReader(strings.NewReader("hello world!"))
    if err != nil {
        panic(err)
    }
    defer f.Close()

    _, err = f.Seek(6, io.SeekStart)
    if err != nil {
        panic(err)
    }

    buf := make([]byte, 6)
    _, err = f.Read(buf)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(buf)) // Output: world!
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
