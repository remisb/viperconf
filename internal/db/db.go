package db

type DbConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	DisableTLS bool
}
