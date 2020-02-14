package config

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// Dir returns the configuration directory for the application
func Dir() string {
	return filepath.Join(xdg.ConfigHome, AppName)
}
