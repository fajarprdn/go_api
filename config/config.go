package config

import "github.com/spf13/viper"

type Config struct {
	path string
	name string
}

func (c Config) readConfigFile() *viper.Viper {
	viper.SetConfigName(c.name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		} else {
			panic("Config File Error")
		}
	}
	return viper.GetViper()
}

func (c Config) Get(key string) string {
	return c.readConfigFile().GetString(key)
}

func New(path string, configFileName string) Config {
	return Config{
		path: path,
		name: configFileName,
	}
}
