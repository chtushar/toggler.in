package configs

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *App
	once sync.Once
)

func Get() *App {
	once.Do(func() {
		setDefaults()
		readConfig()

		cfg = &App{
			Port: viper.GetInt(keyPort),
			Enviroment: viper.GetString(keyEnv),
			Production: isProduction(viper.GetString(keyEnv)),
			JWTSecret: viper.GetString(keyJWTSecret),
			DB: &DB{
				Host: viper.GetString(keyDBHost),
				Port: viper.GetInt(keyDBPort),
				User: viper.GetString(keyDBUser),
				Password: viper.GetString(keyDBPass),
				Name: viper.GetString(keyDBName),
			},
			Logger: &Logger{},
		}
	})
	return cfg
}

// Set default values for configs
func setDefaults(){
	// Loop over all the  default keys in the configs
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

// Read configs from env file
func readConfig(){
	// Gets the home directory similar to _dirname in node.js
	homeDir, err := os.UserHomeDir()
	// Log error if homeDir is not found
	if err != nil {
		log.Fatalf("Failed to find home directory")
	}

	// Set the path to the config file
	viper.SetConfigFile(".env")
	viper.AddConfigPath(homeDir)

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}