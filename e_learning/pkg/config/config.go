package config

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

func GetConfig() Config {
	var myConfig Config

	myConfig.Db.Dsn = "postgres://postgres:123456@localhost:5432/e-learning?sslmode=disable"
	myConfig.Port = 8000
	myConfig.Env = "dev"
	myConfig.Db.MaxOpenConns = 15
	myConfig.Db.MaxIdleConns = 15
	myConfig.Db.MaxIdleTime = "15m"
	return myConfig
}
