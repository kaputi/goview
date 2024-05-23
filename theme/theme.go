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
			"Grid":           "blue",
			"Note1Grid":      "green",
			"Note2Grid":      "red",
			"Fixed":          "yellow",
			"SelectedFixed":  "black,yellow",
			"Placed":         "white",
			"SelectedPlaced": "black,white",
			"Error":          "red",
			"SelectedError":  "black,red",
			"Note":           "cyan",
			"SelectedNote":   "black,cyan",
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
