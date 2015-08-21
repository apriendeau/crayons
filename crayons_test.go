package crayons_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/apriendeau/crayons"
	"github.com/shiena/ansicolor"
	"github.com/stretchr/testify/suite"
)

type CrayonSuite struct {
	suite.Suite
	oldWriter io.Writer
	file      *os.File
}

func read(f *os.File) (string, error) {

	f.Seek(0, 0)
	b, err := ioutil.ReadAll(f)
	return string(b), err
}

func (s *CrayonSuite) SetupTest() {
	var err error
	s.oldWriter = crayons.Writer
	s.file, err = ioutil.TempFile("/tmp", "capture-")
	s.NoError(err)
	crayons.Writer = ansicolor.NewAnsiColorWriter(s.file)
}

func (s *CrayonSuite) TearDownTest() {
	s.file.Close()
	os.Remove(s.file.Name())
	crayons.Writer = s.oldWriter
}

func TestCrayonSuite(t *testing.T) {
	suite.Run(t, new(CrayonSuite))
}

func (s *CrayonSuite) TestCrayonFmt() {
	c := crayons.New(crayons.FgBrightCyan, crayons.BgBlack)
	s.NotNil(c)
	s.Equal("\x1b[96;40m", c.Fmt())
}

func (s *CrayonSuite) TestCrayonUnfmt() {
	c := crayons.New(crayons.FgBrightCyan, crayons.BgBlack)
	base := fmt.Sprintf("%s[%dm", "\x1b", crayons.Clear)
	s.Equal(base, c.Unfmt())
}

func (s *CrayonSuite) TestCrayonSprint() {
	c := crayons.New(crayons.FgBrightCyan, crayons.BgBlack)
	out := c.Sprint("testing")
	base := c.Fmt() + "testing" + c.Unfmt()
	s.Equal(base, out)
}

func (s *CrayonSuite) TestCrayonSprintf() {
	c := crayons.New(crayons.FgBrightRed, crayons.BgWhite)
	out := c.Sprintf("%s %s!", "hello", "world")
	base := c.Fmt() + "hello world!" + c.Unfmt()
	s.Equal(base, out)
}

func (s *CrayonSuite) TestCrayonSprintln() {
	c := crayons.New(crayons.FgBlack, crayons.BgWhite, crayons.Underline)
	out := c.Sprintln("hello", "world", "foo")
	base := c.Fmt() + "hello world foo\n" + c.Unfmt()
	s.Equal(base, out)
}

func (s *CrayonSuite) TestCrayonPrint() {
	c := crayons.New(crayons.FgYellow, crayons.BgRed)
	i, err := c.Print("iron man", "tony stark")
	s.NoError(err)
	s.Len("iron mantony stark", i)
	out, err := read(s.file)
	s.NoError(err)
	s.Len(out, 30)
	exp := c.Fmt() + "iron mantony stark" + c.Unfmt()
	s.Equal(exp, out)
}

func (s *CrayonSuite) TestCrayonPrintf() {
	c := crayons.New(crayons.FgBrightGreen, crayons.BgGreen)
	i, err := c.Printf("%s %s", "the hulk", "bruce banner")
	s.NoError(err)
	s.Len("the hulk bruce banner", i)
	out, err := read(s.file)
	s.NoError(err)
	s.Len(out, 33)
	exp := c.Fmt() + "the hulk bruce banner" + c.Unfmt()
	s.Equal(exp, out)
}

func (s *CrayonSuite) TestCrayonPrintln() {
	c := crayons.New(crayons.FgWhite, crayons.BgBlue, crayons.Underline)
	i, err := c.Println("captain america", "Steve Rodgers")
	s.NoError(err)
	exp := c.Fmt() + "captain america Steve Rodgers\n" + c.Unfmt()
	s.Len("captain america Steve Rodgers\n", i)
	out, err := read(s.file)
	s.NoError(err)
	s.Len(out, 44)
	s.Equal(exp, out)
}

func (s *CrayonSuite) TestCrayonStyles() {
	c := crayons.New(crayons.FgWhite, crayons.BgBlue, crayons.Underline)
	styles := c.Styles()
	s.Len(styles, 3)
}

func (s *CrayonSuite) TestCrayonPrepend() {
	c := crayons.New(crayons.BgBlue, crayons.Underline)
	styles := c.Styles()
	s.Len(styles, 2)

	c.Prepend(crayons.FgWhite)
	styles = c.Styles()
	s.Len(styles, 3)
}

func (s *CrayonSuite) TestCrayonColorize() {
	str := crayons.Colorize("message", crayons.FgCyan)

	c := crayons.New(crayons.FgCyan)
	cstr := c.Sprint("message")
	s.Equal(cstr, str)
}

func (s *CrayonSuite) TestCrayonMonochrome() {
	old := crayons.Monochrome
	crayons.Monochrome = true

	c := crayons.New(crayons.FgCyan)
	c.Monochrome(false)
	s.Equal(crayons.Monochrome, true)
	crayons.Monochrome = false
	c.Monochrome(true)
	str := c.Sprint("message")
	s.Equal("message", str)

	crayons.Monochrome = old
}
