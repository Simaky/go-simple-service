package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const defaultConfigPath = "config.json"

// Config contains all config values
var Config Configuration

// Configuration contains all config values
type Configuration struct {
	LogLevel    string   `json:"log_level"`
	LogFilePath string   `json:"log_file_path"`
	Port        string   `json:"port"`
	Database    Database `json:"database"`
}

// Database contains DB variables
type Database struct {
	Username string `json:"username"`
	DBName   string `json:"db_name"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

// LoadConfig loads configuration from file config.json
func LoadConfig() error {
	cfgFile, err := os.Open(defaultConfigPath)
	if err != nil {
		return err
	}
	defer cfgFile.Close()

	cfgByte, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(cfgByte, &Config)
}
