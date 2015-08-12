package crayons_test

import (
	"github.com/apriendeau/crayons"
)

func Example_Basic() {
	c := crayons.New(crayons.FgCyan)
	c.Println("hello world!")
}
