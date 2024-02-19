package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHOST     string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
}

func InitConfig() *Config {

	//read configuration from a config file
	configFilePath := "C:/Users/Administrator/Desktop/go-workspace/src/BOOK_STORE/config/config.yaml"
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error reading config file:", err)
	}

	//unmarshal the config
	var cnfg Config
	err = viper.Unmarshal(&cnfg)
	if err != nil {
		log.Fatal("error unmarshalling config:", err)
	}

	return &cnfg

}
