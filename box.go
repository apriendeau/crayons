package crayons

import (
	"errors"
	"fmt"
)

var (
	ErrRemoveBase = errors.New("Cannot remove base crayon")
	ErrNilCrayon  = errors.New("Crayon cannot be a nil reference")
)

type box struct {
	Crayons map[string]*Crayon
}

// NewBox creates a box of crayons and applys a default setting. If Base is
// nil, it will default to white text on a black background.
func NewBox(base *Crayon) *box {
	c := make(map[string]*Crayon)
	if base == nil {
		base = New(DefaultFg, DefaultBg)
	}
	c["base"] = base
	return &box{Crayons: c}
}

// Store adds a crayon to the box for later retrieval.
func (b *box) Store(name string, crayon *Crayon) error {
	if crayon == nil {
		return ErrNilCrayon
	}
	if _, ok := b.Crayons[name]; ok {
		msg := "%s is already in the box. Please remove the crayon first."
		return fmt.Errorf(msg, name)
	}
	b.Crayons[name] = crayon
	return nil
}

// Pick retrives a crayon for your coloring pleasure.
func (b *box) Pick(name string) *Crayon {
	c, ok := b.Crayons[name]
	if !ok {
		return b.Crayons["base"]
	}
	return c
}

// Remove destroys a crayon from your box.
func (b *box) Remove(name string) error {
	if name == "base" {
		return ErrRemoveBase
	}
	delete(b.Crayons, name)
	return nil
}

// Names returns the names of all the stored crayons
func (b *box) Names() []string {
	keys := make([]string, 0, len(b.Crayons))
	for k := range b.Crayons {
		keys = append(keys, k)
	}
	return keys
}
