# crayons

[![GoDoc](https://godoc.org/github.com/apriendeau/crayons?status.svg)](https://godoc.org/github.com/apriendeau/crayons)

There are alot of ANSI color libraries but none of them were useful for larger
impletations. It is heavily modeled after [faith/color](https://github.com/faith/color)
but exposes things slightly differently and adds a concept of grouping
different styles.



## Installing

```bash
go get github.com/apriendeau/crayons
```


## Basic Usage

```go
package main

import(
	"fmt"
	"github.com/apriendeau/crayons"
)

func main() {
	c := crayons.New(crayons.FgBrightCyan, crayons.BgBlack)

	c.Println("blizzard blue")

	str := c.Sprintf("blizzard %s", "blue")
	fmt.Println(str)
}

```

