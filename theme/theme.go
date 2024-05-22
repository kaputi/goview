package theme

type Theme struct {
	UnicodeChars Chars
	AsciiChars   Chars
	classes      map[string]string
	useAscii     bool
}

func New() *Theme {
	return &Theme{
		UnicodeChars: Chars{
			DoubleCorners:                  []string{"╔", "╗", "╚", "╝"},
			SingleCorners:                  []string{"┌", "┐", "└", "┘"},
			DoubleTee:                      []string{"╠", "╣", "╦", "╩"},
			SingleTee:                      []string{"├", "┤", "┬", "┴"},
			DoubleCross:                    "╬",
			SingleCross:                    "┼",
			DoubleHorizonatlSingleVertical: "╪",
			DoubleVerticalSingleHorizontal: "╫",
			DoubleHorizontal:               "═",
			SingleHorizontal:               "─",
			DoubleVertical:                 "║",
			SingleVertical:                 "│",
			SingleSideToDoubleTee:          []string{"╞", "╡", "╥", "╨"},
			DoubleSideToSingleTee:          []string{"╟", "╢", "╤", "╧"},
		},
		AsciiChars: Chars{
			DoubleCorners:                  []string{"++", "++", "++", "++"},
			SingleCorners:                  []string{"+", "+", "+", "+"},
			DoubleTee:                      []string{"++", "++", "++", "++"},
			SingleTee:                      []string{"+", "+", "+", "+"},
			DoubleCross:                    "++",
			SingleCross:                    "+",
			DoubleHorizonatlSingleVertical: "+",
			DoubleVerticalSingleHorizontal: "++",
			DoubleHorizontal:               "=",
			SingleHorizontal:               "-",
			DoubleVertical:                 "||",
			SingleVertical:                 "|",
			SingleSideToDoubleTee:          []string{"|", "|", "++", "++"},
			DoubleSideToSingleTee:          []string{"++", "++", "|", "|"},
		},
		classes: map[string]string{
			"grid":           "blue",
			"note1":          "green",
			"note2":          "red",
			"fixed":          "yellow",
			"selectedFixed":  "black,yellow",
			"placed":         "white",
			"selectedPlaced": "black,white",
			"error":          "red",
			"selectedError":  "black,red",
			"note":           "cyan",
			"selectedNote":   "black,cyan",
		},
	}
}

func (t *Theme) UseAscii(option bool) {
	t.useAscii = option
}

func (t *Theme) GetChars() Chars {
	if t.useAscii {
		return t.AsciiChars
	} else {
		return t.UnicodeChars
	}
}

func (t *Theme) GetClassColor(class string) string {
	color, ok := t.classes[class]
	if !ok {
		return t.classes["grid"]
	}
	return color
}
