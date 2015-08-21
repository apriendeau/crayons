package crayons

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/shiena/ansicolor"
)

var (
	// Writer is a where crayons will draw too
	Writer = ansicolor.NewAnsiColorWriter(os.Stdout)
	// Monochrome checkts if it is tty
	Monochrome = !isatty.IsTerminal(os.Stdout.Fd())
)

// Style is alias type for int
type Style int

//Crayon is the structure for a crayon. It contains unexported fields
type Crayon struct {
	styles     []Style
	monochrome bool
}

const escape = "\x1b"

// Core Styles
const (
	Clear Style = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground Colors
const (
	FgBlack Style = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgBrightGrey
	_ // unused
	DefaultFg
)

// Bright Foreground Colors
const (
	FgBrightBlack = iota + 90
	FgBrightRed
	FgBrightGreen
	FgBrightYellow
	FgBrightBlue
	FgBrightMagenta
	FgBrightCyan
	FgWhite
)

// Background Colors
const (
	BgBlack Style = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgLightGrey
	_ // unused
	DefaultBg
)

// Bright Background Colors
const (
	BgBrightBlack Style = iota + 100
	BgBrightRed
	BgBrightGreen
	BgBrightYellow
	BgBrightBlue
	BgBrightMagenta
	BgBrightCyan
	BgWhite
)

// New returns a pointer to an instance of a crayon. You can add several styles
// and it will add them all.
func New(styles ...Style) *Crayon {
	c := &Crayon{
		styles:     make([]Style, 0),
		monochrome: Monochrome,
	}
	c.Append(styles...)
	return c
}

// Styles returns all the stored styles for a crayon
func (c *Crayon) Styles() []Style {
	return c.styles
}

// Append is in case you forgot to add a style
func (c *Crayon) Append(styles ...Style) *Crayon {
	c.styles = append(c.styles, styles...)
	return c
}

// Prepend is for when you want to add a style to the beginning
func (c *Crayon) Prepend(s Style) *Crayon {
	c.styles = append(c.styles, 0)
	copy(c.styles[1:], c.styles[0:])
	c.styles[0] = s
	return c
}

// Monochrome lets you set an individual crayon
func (c *Crayon) Monochrome(m bool) {
	if Monochrome {
		return
	}
	c.monochrome = m
}

// Apply is the manual way of enabling the style for your string
// but will remain in effect unless you call Reset.
func (c *Crayon) Apply() *Crayon {
	if !c.monochrome {
		fmt.Fprintf(Writer, c.Fmt())
	}
	return c
}

// Reset starts clears the styles that are enabled in the writer.
func (c *Crayon) Reset() *Crayon {
	if !c.monochrome {
		fmt.Fprintf(Writer, "%s[%dm", escape, Clear)
	}
	return c
}

// Print will print a string with the styles applied.
func (c *Crayon) Print(a ...interface{}) (int, error) {
	c.Apply()
	defer c.Reset()

	return fmt.Fprint(Writer, a...)
}

// Printf acts same as fmt.Printf but will apply the styles
func (c *Crayon) Printf(base string, a ...interface{}) (int, error) {
	c.Apply()
	defer c.Reset()
	return fmt.Fprintf(Writer, base, a...)

}

// Println acts the same fmt.Println but will apply the styles to each line
func (c *Crayon) Println(a ...interface{}) (int, error) {
	c.Apply()
	defer c.Reset()

	return fmt.Fprintln(Writer, a...)
}

// Sprint applies the styles and acts as fmt.Sprint
func (c *Crayon) Sprint(a ...interface{}) string {
	c.Apply()
	defer c.Reset()

	end := fmt.Sprint(a...)
	return c.wrap(end)
}

// Sprintf acts as fmt.Sprintf
func (c *Crayon) Sprintf(base string, a ...interface{}) string {
	c.Apply()
	defer c.Reset()
	end := fmt.Sprintf(base, a...)
	return c.wrap(end)
}

// Sprintln wraps color around a fmt.Sprintln.
func (c *Crayon) Sprintln(a ...interface{}) string {
	c.Apply()
	defer c.Reset()
	end := fmt.Sprintln(a...)
	return c.wrap(end)
}

func (c *Crayon) seq() string {
	format := make([]string, len(c.styles))
	for i, v := range c.styles {
		format[i] = strconv.Itoa(int(v))
	}
	return strings.Join(format, ";")
}

func (c *Crayon) wrap(s string) string {
	if c.monochrome {
		return s
	}
	return c.Fmt() + s + c.Unfmt()
}

// Fmt is the start of the ANSI color string
func (c *Crayon) Fmt() string {
	return fmt.Sprintf("%s[%sm", escape, c.seq())
}

// Unfmt is an end of the ANSI color string
func (c *Crayon) Unfmt() string {
	return fmt.Sprintf("%s[%dm", escape, Clear)
}

// Colorize is a shortcut for styling.
func Colorize(str string, styles ...Style) string {
	c := New(styles...)
	return c.wrap(str)
}
