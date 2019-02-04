package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	var err error
	config = viper.New()

	// Viper will automatically override config options from the ENV using this prefix.
	// ie. GORT_PORT=8080 GORT_HOST=localhost ./gort
	config.SetEnvPrefix("GORT")
	config.AutomaticEnv()

	// Set defaults.
	config.SetDefault("host", "")
	config.SetDefault("port", "8080")
	config.SetDefault("env", "development")

	// Read our config file.
	config.SetConfigType("yaml")
	config.SetConfigName(config.GetString("env"))
	config.AddConfigPath("config/")

	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
