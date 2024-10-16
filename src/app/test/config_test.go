// src/app/test/config_test.go
package test

import (
	"os"
	"testing"

	"event-router/src/app/internal/config" // Adjust the import path as necessary
)

func TestMain(m *testing.M) {
	// Setup: Create a temporary config file for testing
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Ensure the CONFIG_PATH environment variable is set to a known path during tests

	// Create a temporary config file for testing
	configFile, err := os.Create("./config.yaml")
	if err != nil {
		panic(err)
	}
	os.Setenv("CONFIG_PATH", configFile.Name())

	// Write some sample configuration to the temporary config file
	configContent := []byte(`service:
  host: localhost
  port: 8080`)
	if _, err := configFile.Write(configContent); err != nil {
		panic(err)
	}
	if err := configFile.Close(); err != nil {
		panic(err)
	}
}

func teardown() {
	os.Remove("./config.yaml") // Clean up the temp file after tests are done
}

func TestLoadConfig(t *testing.T) {
	config, err := config.LoadConfig()
	if err != nil {
		t.Errorf("Failed to load config: %v", err)
	}
	if config.Service.Host != "localhost" {
		t.Error("Expected host to be localhost")
	}
	// Add more assertions as necessary to validate your configuration settings
}
