package crayons_test

import (
	"testing"

	"github.com/apriendeau/crayons"
	"github.com/stretchr/testify/suite"
)

type BoxSuite struct {
	suite.Suite
	defaultCrayon *crayons.Crayon
	backupCrayon  *crayons.Crayon
}

func (s *BoxSuite) SetupSuite() {
	s.defaultCrayon = crayons.New(crayons.FgWhite)
	s.backupCrayon = crayons.New(crayons.FgBrightRed)
}

func TestBoxSuite(t *testing.T) {
	suite.Run(t, new(BoxSuite))
}

func (s *BoxSuite) TestNewBox() {
	box := crayons.NewBox(s.defaultCrayon)
	s.NotNil(box)

	names := box.Names()
	s.Len(names, 1)
	s.Equal("base", names[0])
}

func (s *BoxSuite) TestBoxStore() {
	box := crayons.NewBox(s.defaultCrayon)
	err := box.Store("tickle-me-pink", s.backupCrayon)
	s.NoError(err)

	names := box.Names()
	s.Len(names, 2)
	s.Contains(names, "tickle-me-pink")

	err = box.Store("tickle-me-pink", s.backupCrayon)
	s.Error(err)
}

func (s *BoxSuite) TestBoxPick() {
	box := crayons.NewBox(s.defaultCrayon)
	err := box.Store("tickle-me-pink", s.backupCrayon)
	s.NoError(err)

	crayon := box.Pick("tickle-me-pink")
	s.Equal(s.backupCrayon, crayon)
}

func (s *BoxSuite) TestBoxRemove() {
	box := crayons.NewBox(s.defaultCrayon)
	err := box.Store("tickle-me-pink", s.backupCrayon)
	s.NoError(err)

	err = box.Remove("tickle-me-pink")
	s.NoError(err)

	err = box.Remove("base")
	s.Error(err)

	s.Len(box.Names(), 1)
}
