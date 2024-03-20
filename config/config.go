package config

import (
	"encoding/json"
	"os"
)

// CosConfig stores the COS credentials and settings.
type CosConfig struct {
	SecretId  string `json:"SecretId"`
	SecretKey string `json:"SecretKey"`
	Region    string `json:"Region"`
	Bucket    string `json:"Bucket"`
}

// LoadConfig reads the configuration file and unmarshals it into CosConfig struct.
func LoadConfig(path string) (*CosConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &CosConfig{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
