package config

import "github.com/logrusorgru/aurora"

// Error returns a string colorized as error
func Error(s string) string {
	return aurora.Red(s).String()
}

// Warning returns a string colorized as warning
func Warning(s string) string {
	return aurora.Yellow(s).String()
}

// Info returns a string colorized as info
func Info(s string) string {
	return aurora.Green(s).String()
}

// Default returns a string colorized as default
func Default(s string) string {
	return s
}

// Text returns a string colorized as text
func Text(s string) string {
	return aurora.Magenta(s).String()
}
