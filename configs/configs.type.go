package configs

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	TimeZone string
}

type RedisConfig struct {
	Host string
	Port int
}

type MongoDbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}
