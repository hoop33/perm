package config

import "github.com/logrusorgru/aurora"

var au aurora.Aurora

// EnableColor turns color output on or off
func EnableColor(enable bool) {
	au = aurora.NewAurora(enable)
}

// Error returns a string colorized as error
func Error(s string) string {
	return au.Red(s).String()
}

// Warning returns a string colorized as warning
func Warning(s string) string {
	return au.Yellow(s).String()
}

// Info returns a string colorized as info
func Info(s string) string {
	return au.Green(s).String()
}

// Header returns a string colorized as header
func Header(s string) string {
	return au.Blue(s).String()
}

// Default returns a string colorized as default
func Default(s string) string {
	return s
}

// Text returns a string colorized as text
func Text(s string) string {
	return au.Magenta(s).String()
}

func init() {
	au = aurora.NewAurora(true)
}
