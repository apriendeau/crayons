![crayons](https://cloud.githubusercontent.com/assets/700344/9230514/d3e3b2f2-40df-11e5-8f22-cf70a69a08ed.png)

[![GoDoc](http://img.shields.io/badge/go-docs-blue.svg?style=flat-square)](https://godoc.org/github.com/apriendeau/crayons)
[![Build Status](https://img.shields.io/travis/apriendeau/crayons/master.svg?style=flat-square)](https://travis-ci.org/apriendeau/crayons)
[![Coverage](https://img.shields.io/coveralls/apriendeau/crayons/master.svg?style=flat-square)](https://coveralls.io/github/apriendeau/crayons?branch=master)
[![MIT License](https://img.shields.io/github/license/apriendeau/crayons.svg?style=flat-square)](https://github.com/apriendeau/crayons/blob/master/LICENSE)

There are alot of ANSI color libraries but none of them were useful for larger
implementations. It is heavily modeled after [fatih/color](https://github.com/fatih/color)
but exposes things slightly differently and adds a concept of grouping
different styles.

**This is very conceptual right now and in alpha so note, there might be
some changes and is not stable just yet and is totally just a POC repo for me at this point.**

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

## Boxes for Branding

Who wants just one crayon? I want a whole freaking box!

So as I was building a CLI tool, I always had to write another package just to
have to use for styles for certain things such as errors, headers and just
regular text. I found this annoying. As developer, we tend to forget about
branding and UX so I wanted a way to define a style that would be constantly
used throughout my CLI.

The value is consistency.

```go
package main

import(
	"fmt"
	"log"

	"github.com/apriendeau/crayons"
)

func main() {
	redCrayon := crayons.New(crayons.FgBrightRed)
	greenCrayon := crayons.New(crayons.FgBrightGreen)

	box := crayons.NewBox(nil) // nil will default to the defaultFg, defaultBg
	err := box.Store("error", redCrayon)
	if err != nil {
		redCrayon.Fatal(err)
	}
	err = box.Store("success", greenCrayon)
	if err != nil {
		redCrayon.Fatal(err)
	}

	c := box.Pick("success")
	c.Println("Yay, I worked this time")

	c = box.Pick("error")
	c.Println("There was an error! Oh snap!")
	os.Exit(1)
}
```
