package gotint

type TintString struct {
	// The string to be tinted
	content string

	// The tint to be applied
	tint *Tint
}

// NewTintString creates a new tint string
func NewTintString(str string) *TintString {
	return &TintString{
		content: str,
		tint:    NewTint(),
	}
}

// String returns the tinted string
func (ts *TintString) String() string {
	return ts.tint.String() + ts.content + _RESET
}

// Tint returns the tint to be applied
func (ts *TintString) Tint() *Tint {
	return ts.tint
}

// Content returns the string to be tinted
func (ts *TintString) Content() string {
	return ts.content
}

// WithTint sets the tint to be applied and returns the TintString
func (ts *TintString) WithTint(tint *Tint) *TintString {
	ts.tint = tint
	return ts
}

// WithContent sets the string to be tinted and returns the TintString
func (ts *TintString) WithContent(content string) *TintString {
	ts.content = content
	return ts
}
