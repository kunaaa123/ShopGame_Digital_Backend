package configs

type serverConfig struct {
	RESTPort      string
	MySQLHost     string
	MySQLUsername string
	MySQLPassword string
	MySQLDatabase string
}

type Config struct {
	Server *serverConfig
}

func New() *Config {
	c := &Config{}
	c.Server = &serverConfig{
		RESTPort:      ":8080",
		MySQLHost:     "localhost:3306",
		MySQLUsername: "root",
		MySQLPassword: "",
		MySQLDatabase: "gameshop",
	}
	return c
}
