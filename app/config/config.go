package config

import "github.com/spf13/viper"

type AppConfig struct {
	Stage       string `mapstructure:"STAGE"`
	APIBindPort string `mapstructure:"API_BIND_PORT"`
	DBName      string `mapstructure:"APP_DB_NAME"`
	DBHost      string `mapstructure:"APP_DB_HOST"`
	DBPort      string `mapstructure:"APP_DB_PORT"`
	DBUsername  string `mapstructure:"APP_DB_USERNAME"`
	DBPassword  string `mapstructure:"APP_DB_PASSWORD"`
}

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	//	viper.SetDefault("FORCE_CT", "true")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
