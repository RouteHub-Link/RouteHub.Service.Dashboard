package configuration

import (
	"encoding/json"
	"os"
	"sync"
)

const (
	configFile = "config.json"
)

var (
	onceStatic   sync.Once
	staticConfig *StaticConfig
)

type StaticConfig struct {
	S3 S3ClientConfig `json:"s3"`
}

func createConfigFile() error {
	// Create a new file
	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the empty configuration to the schema
	config := StaticConfig{
		S3: S3ClientConfig{},
	}

	configJson, _ := json.Marshal(config)

	_, err = file.Write(configJson)

	return err
}

func initializeStaticConfig() (*StaticConfig, error) {
	// Check if the config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Create the config file if it does not exist
		if err := createConfigFile(); err != nil {
			return nil, err
		}
	}

	// Open the config file
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the config file
	config := StaticConfig{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	return &config, err
}

func updateStaticConfig(config *StaticConfig) error {
	// Open the config file
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the updated configuration to the file
	configJson, _ := json.Marshal(config)
	_, err = file.WriteAt(configJson, 0)
	staticConfig = config
	return err
}

func getStaticConfig() (*StaticConfig, error) {
	onceErr := new(error)
	onceStatic.Do(func() {
		config, err := initializeStaticConfig()
		if err != nil {
			onceErr = &err
			return
		}
		staticConfig = config
	})

	return staticConfig, *onceErr
}
