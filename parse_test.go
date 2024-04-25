package env

import (
	"os"
	"testing"
)

// func TestParse(t *testing.T) {
// 	type Config struct {
// 		Port     int    `env:"PORT"`
// 		Password string `env:"PASSWORD"`
// 		IsOnline bool   `env:"IS_ONLINE"`
// 	}

// 	var config Config
// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if config.Port != 0 {
// 		t.Errorf("Expected 0, got %d", config.Port)
// 	}

// 	if config.Password != "" {
// 		t.Errorf("Expected empty string, got %s", config.Password)
// 	}

// 	if config.IsOnline != false {
// 		t.Errorf("Expected false, got %t", config.IsOnline)
// 	}

// 	// Set environment variables
// 	os.Setenv("PORT", "8080")
// 	os.Setenv("PASSWORD", "password")
// 	os.Setenv("IS_ONLINE", "true")

// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if config.Port != 8080 {
// 		t.Errorf("Expected 8080, got %d", config.Port)
// 	}

// 	if config.Password != "password" {
// 		t.Errorf("Expected password, got %s", config.Password)
// 	}

// 	if config.IsOnline != true {
// 		t.Errorf("Expected true, got %t", config.IsOnline)
// 	}
// }

// func TestParseWithDefaultValue(t *testing.T) {
// 	type Config struct {
// 		Port int `env:"PORT_WITH_DEFAULT_CONFIG,default=9999"`
// 	}

// 	var config Config
// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if config.Port != 9999 {
// 		t.Errorf("Expected 9999, got %d", config.Port)
// 	}

// 	// Set environment variables
// 	os.Setenv("PORT_WITH_DEFAULT_CONFIG", "8080")
// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if config.Port != 8080 {
// 		t.Errorf("Expected 8080, got %d", config.Port)
// 	}
// }

// func TestParseWithRequired(t *testing.T) {
// 	type Config struct {
// 		Port int `env:"PORT_WITH_REQUIRED,required"`
// 	}

// 	var config Config
// 	if err := Parse(&config); err == nil {
// 		t.Fatalf("Expected Parse() to return an error, but go nil")
// 	} else if err.Error() != "PORT_WITH_REQUIRED is required" {
// 		t.Fatalf("Expected PORT_WITH_REQUIRED is required, got %s", err.Error())
// 	}
// }

// func TestParseWithSlice(t *testing.T) {
// 	type Config struct {
// 		Hosts []string `env:"HOSTS"`
// 	}

// 	var config Config
// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if len(config.Hosts) != 0 {
// 		t.Fatalf("Expected 0, got %d", len(config.Hosts))
// 	}

// 	// Set environment variables
// 	os.Setenv("HOSTS", "localhost,live.local")
// 	if err := Parse(&config); err != nil {
// 		t.Fatalf("Expected Parse() to not return an error: %s", err.Error())
// 	}

// 	if len(config.Hosts) != 2 {
// 		t.Fatalf("Expected 2, got %d", len(config.Hosts))
// 	}

// 	if config.Hosts[0] != "localhost" {
// 		t.Fatalf("Expected localhost, got %s", config.Hosts[0])
// 	}

// 	if config.Hosts[1] != "live.local" {
// 		t.Fatalf("Expected live.local, got %s", config.Hosts[1])
// 	}
// }

func TestParseWithNestConfig(t *testing.T) {
	type Config struct {
		Port int `env:"PORT"`

		Database struct {
			Host string `env:"DB_HOST"`
			Port int    `env:"DB_PORT"`
			User string `env:"DB_USER"`
			Pass string `env:"DB_PASS"`
			Name string `env:"DB_NAME"`
		}
	}

	var config Config

	// Set environment variables
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "root")
	os.Setenv("DB_NAME", "test")

	if err := Parse(&config); err != nil {
		t.Fatalf("Expected Parse() to not return an error: %v", err)
	}

	if config.Port != 8080 {
		t.Errorf("Expected 8080, got %d", config.Port)
	}

	if config.Database.Host != "localhost" {
		t.Errorf("Expected localhost, got %s", config.Database.Host)
	}

	if config.Database.Port != 3306 {
		t.Errorf("Expected 3306, got %d", config.Database.Port)
	}

	if config.Database.User != "root" {
		t.Errorf("Expected root, got %s", config.Database.User)
	}

	if config.Database.Pass != "root" {
		t.Errorf("Expected root, got %s", config.Database.Pass)
	}

	if config.Database.Name != "test" {
		t.Errorf("Expected test, got %s", config.Database.Name)
	}
}
