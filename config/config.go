package config

// Configurations for future use if needed
type Config struct {
    Port string
}

func GetConfig() *Config {
    return &Config{
        Port: "8080",
    }
}