package view

import "strings"

var ColorFg = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}
var BgColor = map[string]string{
	"black":   "\033[40m",
	"red":     "\033[41m",
	"green":   "\033[42m",
	"yellow":  "\033[43m",
	"blue":    "\033[44m",
	"magenta": "\033[45m",
	"cyan":    "\033[46m",
	"white":   "\033[47m",
}

var ResetColor = "\033[0m"

// takes a string of fg,bg colors separated by comma and returns the color code
func GetColorCode(colorStr string) string {
	str := ""

	colors := strings.Split(colorStr, ",")
	for i, color := range colors {
		colorMap := ColorFg
		if i == 1 {
			colorMap = BgColor
		}
		code, ok := colorMap[color]
		if ok {
			str += code
		}
	}

	if len(str) == 0 {
		str = ResetColor
	}

	return str
}
