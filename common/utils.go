package common

import (
	"log"
	"os"

	"path/filepath"
)

func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, ".chatgptui.json")
}
