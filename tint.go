package gotint

type Tint struct {
	// the style to be applied
	style []TintStyle

	// the foreground color to be applied
	fgColor TintFGColor

	// the background color to be applied
	bgColor TintBGColor
}

// NewTint creates a new base tint
func NewTint() *Tint {
	return &Tint{
		style:   []TintStyle{_NORMAL_ST},
		fgColor: _NORMAL_FG,
		bgColor: _NORMAL_BG,
	}
}

func (t *Tint) String() string {
	result := ""
	for _, style := range t.style {
		result += string(style)
	}
	return result + string(t.fgColor) + string(t.bgColor)
}

// WithStyle returns a new tint with the given style applied in place of the existing styles
func (t *Tint) WithStyle(style TintStyle) *Tint {
	t.style = []TintStyle{style}
	return t
}

// AddStyle returns a new tint with the given style applied in addition to the existing styles
func (t *Tint) AddStyle(style TintStyle) *Tint {
	t.style = append(t.style, style)
	return t
}

// WithFGColor returns a new tint with the given foreground color applied
func (t *Tint) WithFGColor(fgColor TintFGColor) *Tint {
	t.fgColor = fgColor
	return t
}

// WithBGColor returns a new tint with the given background color applied
func (t *Tint) WithBGColor(bgColor TintBGColor) *Tint {
	t.bgColor = bgColor
	return t
}

// NoStyle returns a new tint with no style applied
func (t *Tint) NoStyle() *Tint {
	return t.WithStyle(_NORMAL_ST)
}

// Bold returns a new tint with the bold style applied
func (t *Tint) Bold() *Tint {
	return t.WithStyle(BOLD)
}

// AddBold returns a new tint with the bold style applied in addition to the existing styles
func (t *Tint) AddBold() *Tint {
	return t.AddStyle(BOLD)
}

// Dim returns a new tint with the dim style applied
func (t *Tint) Dim() *Tint {
	return t.WithStyle(DIM)
}

// AddDim returns a new tint with the dim style applied in addition to the existing styles
func (t *Tint) AddDim() *Tint {
	return t.AddStyle(DIM)
}

// Italic returns a new tint with the italic style applied
func (t *Tint) Italic() *Tint {
	return t.WithStyle(ITALIC)
}

// AddItalic returns a new tint with the italic style applied in addition to the existing styles
func (t *Tint) AddItalic() *Tint {
	return t.AddStyle(ITALIC)
}

// Underline returns a new tint with the underline style applied
func (t *Tint) Underline() *Tint {
	return t.WithStyle(UNDERLINE)
}

// AddUnderline returns a new tint with the underline style applied in addition to the existing styles
func (t *Tint) AddUnderline() *Tint {
	return t.AddStyle(UNDERLINE)
}

// Blink returns a new tint with the blink style applied
func (t *Tint) Blink() *Tint {
	return t.WithStyle(BLINK)
}

// AddBlink returns a new tint with the blink style applied in addition to the existing styles
func (t *Tint) AddBlink() *Tint {
	return t.AddStyle(BLINK)
}

// Reverse returns a new tint with the reverse style applied
func (t *Tint) Reverse() *Tint {
	return t.WithStyle(REVERSE)
}

// AddReverse returns a new tint with the reverse style applied in addition to the existing styles
func (t *Tint) AddReverse() *Tint {
	return t.AddStyle(REVERSE)
}

// Hidden returns a new tint with the hidden style applied
func (t *Tint) Hidden() *Tint {
	return t.WithStyle(HIDDEN)
}

// AddHidden returns a new tint with the hidden style applied in addition to the existing styles
func (t *Tint) AddHidden() *Tint {
	return t.AddStyle(HIDDEN)
}

// NormalFG returns a new tint with the base foreground color applied
func (t *Tint) NormalFG() *Tint {
	return t.WithFGColor(_NORMAL_FG)
}

// Black returns a new tint with the black foreground color applied
func (t *Tint) Black() *Tint {
	return t.WithFGColor(FG_BLACK)
}

// Red returns a new tint with the red foreground color applied
func (t *Tint) Red() *Tint {
	return t.WithFGColor(FG_RED)
}

// Green returns a new tint with the green foreground color applied
func (t *Tint) Green() *Tint {
	return t.WithFGColor(FG_GREEN)
}

// Yellow returns a new tint with the yellow foreground color applied
func (t *Tint) Yellow() *Tint {
	return t.WithFGColor(FG_YELLOW)
}

// Blue returns a new tint with the blue foreground color applied
func (t *Tint) Blue() *Tint {
	return t.WithFGColor(FG_BLUE)
}

// Magenta returns a new tint with the magenta foreground color applied
func (t *Tint) Magenta() *Tint {
	return t.WithFGColor(FG_MAGENTA)
}

// Cyan returns a new tint with the cyan foreground color applied
func (t *Tint) Cyan() *Tint {
	return t.WithFGColor(FG_CYAN)
}

// White returns a new tint with the white foreground color applied
func (t *Tint) White() *Tint {
	return t.WithFGColor(FG_WHITE)
}

// BrightBlack returns a new tint with the bright black foreground color applied
func (t *Tint) BrightBlack() *Tint {
	return t.WithFGColor(FG_BRIGHT_BLACK)
}

// BrightRed returns a new tint with the bright red foreground color applied
func (t *Tint) BrightRed() *Tint {
	return t.WithFGColor(FG_BRIGHT_RED)
}

// BrightGreen returns a new tint with the bright green foreground color applied
func (t *Tint) BrightGreen() *Tint {
	return t.WithFGColor(FG_BRIGHT_GREEN)
}

// BrightYellow returns a new tint with the bright yellow foreground color applied
func (t *Tint) BrightYellow() *Tint {
	return t.WithFGColor(FG_BRIGHT_YELLOW)
}

// BrightBlue returns a new tint with the bright blue foreground color applied
func (t *Tint) BrightBlue() *Tint {
	return t.WithFGColor(FG_BRIGHT_BLUE)
}

// BrightMagenta returns a new tint with the bright magenta foreground color applied
func (t *Tint) BrightMagenta() *Tint {
	return t.WithFGColor(FG_BRIGHT_MAGENTA)
}

// BrightCyan returns a new tint with the bright cyan foreground color applied
func (t *Tint) BrightCyan() *Tint {
	return t.WithFGColor(FG_BRIGHT_CYAN)
}

// BrightWhite returns a new tint with the bright white foreground color applied
func (t *Tint) BrightWhite() *Tint {
	return t.WithFGColor(FG_BRIGHT_WHITE)
}

// OnNormal returns a new tint with the base background color applied
func (t *Tint) OnNormal() *Tint {
	return t.WithBGColor(_NORMAL_BG)
}

// OnBlack returns a new tint with the black background color applied
func (t *Tint) OnBlack() *Tint {
	return t.WithBGColor(BG_BLACK)
}

// OnRed returns a new tint with the red background color applied
func (t *Tint) OnRed() *Tint {
	return t.WithBGColor(BG_RED)
}

// OnGreen returns a new tint with the green background color applied
func (t *Tint) OnGreen() *Tint {
	return t.WithBGColor(BG_GREEN)
}

// OnYellow returns a new tint with the yellow background color applied
func (t *Tint) OnYellow() *Tint {
	return t.WithBGColor(BG_YELLOW)
}

// OnBlue returns a new tint with the blue background color applied
func (t *Tint) OnBlue() *Tint {
	return t.WithBGColor(BG_BLUE)
}

// OnMagenta returns a new tint with the magenta background color applied
func (t *Tint) OnMagenta() *Tint {
	return t.WithBGColor(BG_MAGENTA)
}

// OnCyan returns a new tint with the cyan background color applied
func (t *Tint) OnCyan() *Tint {
	return t.WithBGColor(BG_CYAN)
}

// OnWhite returns a new tint with the white background color applied
func (t *Tint) OnWhite() *Tint {
	return t.WithBGColor(BG_WHITE)
}

// OnBrightBlack returns a new tint with the bright black background color applied
func (t *Tint) OnBrightBlack() *Tint {
	return t.WithBGColor(BG_BRIGHT_BLACK)
}

// OnBrightRed returns a new tint with the bright red background color applied
func (t *Tint) OnBrightRed() *Tint {
	return t.WithBGColor(BG_BRIGHT_RED)
}

// OnBrightGreen returns a new tint with the bright green background color applied
func (t *Tint) OnBrightGreen() *Tint {
	return t.WithBGColor(BG_BRIGHT_GREEN)
}

// OnBrightYellow returns a new tint with the bright yellow background color applied
func (t *Tint) OnBrightYellow() *Tint {
	return t.WithBGColor(BG_BRIGHT_YELLOW)
}

// OnBrightBlue returns a new tint with the bright blue background color applied
func (t *Tint) OnBrightBlue() *Tint {
	return t.WithBGColor(BG_BRIGHT_BLUE)
}

// OnBrightMagenta returns a new tint with the bright magenta background color applied
func (t *Tint) OnBrightMagenta() *Tint {
	return t.WithBGColor(BG_BRIGHT_MAGENTA)
}

// OnBrightCyan returns a new tint with the bright cyan background color applied
func (t *Tint) OnBrightCyan() *Tint {
	return t.WithBGColor(BG_BRIGHT_CYAN)
}

// OnBrightWhite returns a new tint with the bright white background color applied
func (t *Tint) OnBrightWhite() *Tint {
	return t.WithBGColor(BG_BRIGHT_WHITE)
}
