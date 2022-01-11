# ErrOverflow

The **real** way to handle errors.

## Usage

```go
package main

import (
	"strconv"

	"github.com/vidhanio/erroverflow"
)

func main() {
	_, err := strconv.Atoi("")
	if err != nil {
		erroverflow.Handle(err)
	}
}
```
