package utils

import (
	"fmt"

	fcolor "github.com/fatih/color"
)

type focus int

const (
	foreground focus = iota + 1
	background
)

type colorizer struct {
	text  string
	focus focus
	fg    fcolor.Attribute
	bg    fcolor.Attribute
	attrs []fcolor.Attribute
}

const (
	black   = fcolor.FgBlack
	red     = fcolor.FgRed
	green   = fcolor.FgGreen
	yellow  = fcolor.FgYellow
	blue    = fcolor.FgBlue
	magenta = fcolor.FgMagenta
	cyan    = fcolor.FgCyan
	white   = fcolor.FgWhite

	// ansi bg colors are +10 fg color
	bgOffset = 10
	// ansi hi intensity colors are +60
	hiOffset = 60
)

func (c *colorizer) Black(s ...any) (*colorizer, error) {
	return c.setColor(black).apply(s)
}

func (c *colorizer) HiBlack(s ...any) (*colorizer, error) {
	return c.setColor(black + hiOffset).apply(s)
}

func (c *colorizer) Red(s ...any) (*colorizer, error) {
	return c.setColor(red).apply(s)
}

func (c *colorizer) HiRed(s ...any) (*colorizer, error) {
	return c.setColor(red + hiOffset).apply(s)
}

func (c *colorizer) Green(s ...any) (*colorizer, error) {
	return c.setColor(green).apply(s)
}

func (c *colorizer) HiGreen(s ...any) (*colorizer, error) {
	return c.setColor(green + hiOffset).apply(s)
}

func (c *colorizer) Yellow(s ...any) (*colorizer, error) {
	return c.setColor(yellow).apply(s)
}

func (c *colorizer) HiYellow(s ...any) (*colorizer, error) {
	return c.setColor(yellow + hiOffset).apply(s)
}

func (c *colorizer) Blue(s ...any) (*colorizer, error) {
	return c.setColor(blue).apply(s)
}

func (c *colorizer) HiBlue(s ...any) (*colorizer, error) {
	return c.setColor(blue + hiOffset).apply(s)
}

func (c *colorizer) Magenta(s ...any) (*colorizer, error) {
	return c.setColor(magenta).apply(s)
}

func (c *colorizer) HiMagenta(s ...any) (*colorizer, error) {
	return c.setColor(magenta + hiOffset).apply(s)
}

func (c *colorizer) Cyan(s ...any) (*colorizer, error) {
	return c.setColor(cyan).apply(s)
}

func (c *colorizer) HiCyan(s ...any) (*colorizer, error) {
	return c.setColor(cyan + hiOffset).apply(s)
}

func (c *colorizer) White(s ...any) (*colorizer, error) {
	return c.setColor(white).apply(s)
}

func (c *colorizer) HiWhite(s ...any) (*colorizer, error) {
	return c.setColor(white + hiOffset).apply(s)
}

func (c *colorizer) setColor(col fcolor.Attribute) *colorizer {
	if c.focus == foreground {
		c.fg = col
		return c
	}
	c.bg = col + bgOffset
	return c
}

func (c *colorizer) apply(s []any) (*colorizer, error) {
	if len(s) > 1 {
		return nil, fmt.Errorf("expected 0 or 1 argument, got %d", len(s))
	}

	if len(s) > 0 {
		switch v := s[0].(type) {
		case *colorizer:
			c.text = v.text
			c.focus = v.focus
			if v.bg != 0 {
				c.bg = v.bg
			}
			if v.fg != 0 {
				c.fg = v.fg
			}
			c.attrs = append(c.attrs, v.attrs...)
		case *stylizer:
			c.text = v.text
			c.attrs = append(c.attrs, v.attrs...)
		case string:
			c.text = v
		default:
			return nil, fmt.Errorf("unexpected argument type %T", s[0])
		}
	}

	// fmt.Fprintf(os.Stderr, "Colorizer: %#v\n", c)

	return c, nil
}

func (c colorizer) String() string {
	if fcolor.NoColor {
		return c.text
	}
	attrs := []fcolor.Attribute{}
	if c.fg != 0 {
		attrs = append(attrs, c.fg)
	}
	if c.bg != 0 {
		attrs = append(attrs, c.bg)
	}
	attrs = append(attrs, c.attrs...)
	return fcolor.New(attrs...).Sprint(c.text)
}

type stylizer struct {
	text  string
	attrs []fcolor.Attribute
}

func (s *stylizer) apply(args []any) (any, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("expected 0 or 1 argument, got %d", len(args))
	}

	if len(args) > 0 {
		switch v := args[0].(type) {
		case *colorizer:
			// if we are merging styles with colors then
			// upgrade the style to a color, eg:
			//	{{"hi" | bold | fg.Red}}
			c := &colorizer{
				text:  s.text,
				attrs: s.attrs,
			}
			return c.apply(args)
		case *stylizer:
			s.text = v.text
			s.attrs = append(s.attrs, v.attrs...)
		case string:
			s.text = v
		default:
			return nil, fmt.Errorf("unexpected argument type %T", args[0])
		}
	}

	return s, nil
}

func (s stylizer) String() string {
	if fcolor.NoColor {
		return s.text
	}
	attrs := []fcolor.Attribute{}
	attrs = append(attrs, s.attrs...)
	return fcolor.New(attrs...).Sprint(s.text)
}

func ColorFuncMap() map[string]any {
	return map[string]any{
		"fg": func(s ...any) (*colorizer, error) {
			return (&colorizer{focus: foreground}).apply(s)
		},
		"bg": func(s ...any) (*colorizer, error) {
			return (&colorizer{focus: background}).apply(s)
		},
		"bold": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.Bold}}).apply(s)
		},
		"dim": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.Faint}}).apply(s)
		},
		"italic": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.Italic}}).apply(s)
		},
		"underline": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.Underline}}).apply(s)
		},
		"slowBlink": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.BlinkSlow}}).apply(s)
		},
		"rapidBlink": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.BlinkRapid}}).apply(s)
		},
		"invert": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.ReverseVideo}}).apply(s)
		},
		"hide": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.Concealed}}).apply(s)
		},
		"strike": func(s ...any) (any, error) {
			return (&stylizer{attrs: []fcolor.Attribute{fcolor.CrossedOut}}).apply(s)
		},
	}
}
