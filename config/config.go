package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisAddr string
	RedisPass string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "ghost8797",
		DBName:     "ecommerce",
		RedisAddr:  "localhost:6379",
		RedisPass:  "ghost8797",
	}
}
