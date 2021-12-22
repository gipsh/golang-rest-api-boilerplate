package config

import "github.com/spf13/viper"

type AppConfig struct {
	Stage       string `mapstructure:"STAGE"`
	APIBindPort string `mapstructure:"API_BIND_PORT"`
	DBFile      string `mapstructure:"APP_DB_FILE"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
