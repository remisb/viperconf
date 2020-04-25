package main

import (
	"github.com/remisb/viperconf/internal/db"
	"github.com/remisb/viperconf/internal/log"
	"github.com/spf13/viper"
)

const (
	configFileName = "config"
	envPrefix = "VC"
)

func main() {
	// list all command to be server by admin app

	readConfigFile()
}

func readConfigFile() {
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// setup environment variables
	viper.SetEnvPrefix(envPrefix)
	viper.BindEnv("PORT")

	envPort := viper.GetString("port")

	err := viper.ReadInConfig()
	if err != nil {
		log.Sugar.Errorf("read config error: %s", err)
	}

	port := viper.GetInt("Port")
	dbConfig := db.DbConfig{}
	dbConfig.Name = viper.GetString("Db.Name")
	dbConfig.Host = viper.GetString("Db.Host")
	dbConfig.Port = viper.GetString("Db.Port")
	dbConfig.User = viper.GetString("Db.User")
	dbConfig.Password = viper.GetString("Db.Password")
	dbConfig.DisableTLS = viper.GetBool("Db.DisableTLS")

	log.Sugar.Infof("port: %d)", port)
	log.Sugar.Infof("envPort: %s)", envPort)

	log.Sugar.Infof("dbConfig: %v", dbConfig)
}
