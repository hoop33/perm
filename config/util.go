package config

import (
	"path/filepath"

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
