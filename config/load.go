package config

import (
	"os"

	"encoding/json"

	"github.com/go-playground/validator/v10"

	"github.com/dwisiswant0/chatgptui/common"
)

func Load(path string) (common.Config, error) {
	var cfg common.Config

	file, err := os.Open(path)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
