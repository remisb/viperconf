package main

import (
	"github.com/remisb/viperconf/internal/db"
	"github.com/remisb/viperconf/internal/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

const (
	configFileName = "config"
	envPrefix      = "vc"
)

func main() {
	// list all command to be server by admin app

	readConfigFile()

}

func readConfigFile() {

	// setup cli flags
	pflag.CommandLine.IntP("port", "p", 8080, "api service port")
	pflag.String("db-name", "postgres", "Database name")
	pflag.String("db-host", "localhost", "Database host")
	pflag.String("db-port", "5432", "Database port")
	pflag.String("db-user", "postgres", "Database user")
	pflag.String("db-password", "postgres", "Database password")
	pflag.Bool("db-disable-tls", true, "Database disable TLS")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		log.Sugar.Errorf("bind CLI flags error: %v", err)
	}

	// setup config file variables
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// setup environment variables
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	bindEnv("port")
	bindEnv("db-host")
	bindEnv("db-port")
	bindEnv("db-name")
	bindEnv("db-user")
	bindEnv("db-password")
	bindEnv("db-disable-tls")

	xPort := viper.GetInt("port")

	if err := viper.ReadInConfig(); err != nil {
		log.Sugar.Errorf("read config error: %v", err)
	}

	//var dbConfig db.DbConfig
	//viper.UnmarshalKey("Db", &dbConfig)

	dbConfig := db.DbConfig{
		Name: viper.GetString("db-name"),
		Host: viper.GetString("db-host"),
		Port: viper.GetString("db-port"),
		User: viper.GetString("db-user"),
		Password: viper.GetString("db-password"),
		DisableTLS: viper.GetBool("db-disable-tls"),
	}

	log.Sugar.Infof("Port: %d)", xPort)
	log.Sugar.Infof("dbConfig: %+v", dbConfig)
}

func bindEnv(name string) {
	if err := viper.BindEnv("name"); err != nil {
		log.Sugar.Errorf("bind env var %s error: %v", name, err)
	}
}
