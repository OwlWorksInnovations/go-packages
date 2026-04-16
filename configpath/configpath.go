package configpath

import (
	"os"
	"path/filepath"
)

func CreateConfigPath(configFolder string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	configPath := filepath.Join(home, configFolder)
	os.MkdirAll(configPath, 0755)
}

func GetConfigPath(configFolder string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, configFolder)
}
