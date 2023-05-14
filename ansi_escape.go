package gotint

type TintStyle string
type TintFGColor string
type TintBGColor string

const (
	_RESET = "\033[0m"
)

const (
	// Styles
	_NORMAL_ST TintStyle = ""
	BOLD       TintStyle = "\033[1m"
	DIM        TintStyle = "\033[2m"
	ITALIC     TintStyle = "\033[3m"
	UNDERLINE  TintStyle = "\033[4m"
	BLINK      TintStyle = "\033[5m"
	REVERSE    TintStyle = "\033[7m"
	HIDDEN     TintStyle = "\033[8m"
)

const (
	_NORMAL_FG TintFGColor = ""

	// Foreground Colors
	FG_BLACK   TintFGColor = "\033[30m"
	FG_RED     TintFGColor = "\033[31m"
	FG_GREEN   TintFGColor = "\033[32m"
	FG_YELLOW  TintFGColor = "\033[33m"
	FG_BLUE    TintFGColor = "\033[34m"
	FG_MAGENTA TintFGColor = "\033[35m"
	FG_CYAN    TintFGColor = "\033[36m"
	FG_WHITE   TintFGColor = "\033[37m"

	// Bright Foreground Colors
	FG_BRIGHT_BLACK   TintFGColor = "\033[90m"
	FG_BRIGHT_RED     TintFGColor = "\033[91m"
	FG_BRIGHT_GREEN   TintFGColor = "\033[92m"
	FG_BRIGHT_YELLOW  TintFGColor = "\033[93m"
	FG_BRIGHT_BLUE    TintFGColor = "\033[94m"
	FG_BRIGHT_MAGENTA TintFGColor = "\033[95m"
	FG_BRIGHT_CYAN    TintFGColor = "\033[96m"
	FG_BRIGHT_WHITE   TintFGColor = "\033[97m"
)

const (
	_NORMAL_BG TintBGColor = ""

	// Background Colors
	BG_BLACK   TintBGColor = "\033[40m"
	BG_RED     TintBGColor = "\033[41m"
	BG_GREEN   TintBGColor = "\033[42m"
	BG_YELLOW  TintBGColor = "\033[43m"
	BG_BLUE    TintBGColor = "\033[44m"
	BG_MAGENTA TintBGColor = "\033[45m"
	BG_CYAN    TintBGColor = "\033[46m"
	BG_WHITE   TintBGColor = "\033[47m"

	// Bright Background Colors
	BG_BRIGHT_BLACK   TintBGColor = "\033[100m"
	BG_BRIGHT_RED     TintBGColor = "\033[101m"
	BG_BRIGHT_GREEN   TintBGColor = "\033[102m"
	BG_BRIGHT_YELLOW  TintBGColor = "\033[103m"
	BG_BRIGHT_BLUE    TintBGColor = "\033[104m"
	BG_BRIGHT_MAGENTA TintBGColor = "\033[105m"
	BG_BRIGHT_CYAN    TintBGColor = "\033[106m"
	BG_BRIGHT_WHITE   TintBGColor = "\033[107m"
)
