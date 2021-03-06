package config

import (
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

// Dir returns the configuration directory for the application
func Dir() string {
	return filepath.Join(xdg.ConfigHome, AppName)
}

// HistoryFile returns the full path of the history file
func HistoryFile() string {
	return filepath.Join(Dir(), ".history")
}

// AddPrefixIfAbsent adds the specified prefix to the specified string if not present
func AddPrefixIfAbsent(str, prefix string) string {
	if !strings.HasPrefix(str, prefix) {
		return prefix + str
	}
	return str
}

// FirstNonBlank returns the first non-blank string
func FirstNonBlank(strs ...string) string {
	for _, s := range strs {
		if s != "" {
			return s
		}
	}
	return ""
}
