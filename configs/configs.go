package configs

// Configs
var ServiceName *string
var Port *int
var EnableQueue *bool
var EnableJob *bool
var DBConfig *DBConfig
var Redis *RedisConfig
var MongoDB *MongoDbConfig

func BootConfig() {
	env := NewEnvironment()

	ServiceName = env.GetString("SERVICE_NAME")
	Port = env.GetInt("PORT", 3000)
	EnableQueue = env.GetBool("ENABLE_QUEUE", false)
	EnableJob = env.GetBool("ENABLE_JOB", false)
	DBConfig = &DBConfig{
		Host:     *env.GetString("POSTGRES_HOST"),
		Port:     *env.GetInt("POSTGRES_PORT"),
		Username: *env.GetString("POSTGRES_USER"),
		Password: *env.GetString("POSTGRES_PASSWORD"),
		DBName:   *env.GetString("POSTGRES_NAME"),
		TimeZone: "Asia/Bangkok",
	}
	Redis = &RedisConfig{
		Host: *env.GetString("REDIS_HOST", "172.16.23.27"),
		Port: *env.GetInt("REDIS_PORT", 6379),
	}
	MongoDB = &MongoDbConfig{
		User:     *env.GetString("MONGODB_USER"),
		Password: *env.GetString("MONGODB_PASSWORD"),
		Host:     *env.GetString("MONGODB_HOST"),
		Port:     *env.GetString("MONGODB_PORT", "27017"),
		Database: *env.GetString("MONGODB_DATABASE"),
	}
}
