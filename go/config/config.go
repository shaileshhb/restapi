package config

import "github.com/spf13/viper"

// Config stores all configuration of the application.
type Config struct {
	DBRootPassword  string `mapstructure:"DB_ROOT_PASSWORD"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBDatabase      string `mapstructure:"DB_DATABASE"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBHost          string `mapstructure:"DB_HOST"`
	AccessSecretKey string `mapstructure:"ACCESS_SECRET"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
