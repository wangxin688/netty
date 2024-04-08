package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

var Settings *Config

var ProjectPath string

type Environment string

const (
	Dev   Environment = "dev"
	Stage Environment = "stage"
	Prod  Environment = "prod"
)

type Config struct {
	AccessTokenExpireMinutes  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRE_MINUTES"`
	RefreshTokenExpireMinutes time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRE_MINUTES"`
	SecretKey                 string        `mapstructure:"SECRET_KEY"`
	ServerAddress             string        `mapstructure:"SERVER_ADDRESS"`
	ServerTimeZone            string        `mapstructure:"SERVER_TIMEZONE"`
	AllowedHosts              []string      `mapstructure:"ALLOWED_HOSTS"`
	DBHost                    string        `mapstructure:"DB_HOST"`
	DBPort                    int           `mapstructure:"DB_PORT"`
	DBUser                    string        `mapstructure:"DB_USER"`
	DBPassword                string        `mapstructure:"DB_PASSWORD"`
	DBName                    string        `mapstructure:"DB_NAME"`
	MaxLifetime               time.Duration `mapstructure:"MAX_LIFETIME"`
	MaxIdleTime               time.Duration `mapstructure:"MAX_IDLE_TIME"`
	MaxOpenConns              int           `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleConns              int           `mapstructure:"MAX_IDLE_CONNS"`
	DBDebugMode               bool          `mapstructure:"DB_DEBUG_MODE"`
	RedisHost                 string        `mapstructure:"REDIS_HOST"`
	RedisPort                 int           `mapstructure:"REDIS_PORT"`
	RedisPassword             string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB                   int           `mapstructure:"REDIS_DB"`

	Env Environment `mapstructure:"ENV"`
}

// SetupConfig sets up the Config struct from the .env file in the root of the project.
// If the .env file is not found, it will set defaults for the fields.
// The .env file must be in KEY=VALUE format. The fields of the Config struct must
// match the keys in the .env file.
//
// It returns an error if there is an issue reading the file or unmarshaling the
// config into the Config struct.
func SetupConfig() (err error) {

	var settings = Settings

	_, currentPath, _, _ := runtime.Caller(0)

	projectPath := filepath.Dir(filepath.Dir(currentPath))

	// Set defaults for the config file values
	viper.SetDefault("AccessTokenExpireMinutes", time.Duration(60*24)*time.Minute)
	viper.SetDefault("RefreshTokenExpireMinutes", time.Duration(60*24*8)*time.Minute)
	viper.SetDefault("ServerTimeZone", "Asia/Shanghai")
	viper.SetDefault("DBDebugMode", false)
	viper.SetDefault("RedisPort", 6379)
	viper.SetDefault("RedisDB", 0)

	// Set the config file name and type
	viper.SetConfigFile(projectPath + "/.env")
	viper.SetConfigType("env")

	// Attempt to read the config file
	if err = viper.ReadInConfig(); err != nil {
		// logger.Errorf("Error reading config file, %s", err)
		return err
	}

	// Unmarshal the config file values into the Config struct
	if err = viper.Unmarshal(&settings); err != nil {
		return fmt.Errorf("error unmarshaling viper config into `Settings`: %w", err)
	}

	// Update the ProjectPath var
	ProjectPath = projectPath

	Settings = settings

	return nil
}

// BuildPgDsn returns a connection string for connecting to the Postgres database.
// The connection string is in the format:
//
//	host=<dbHost> port=<dbPort> user=<dbUser> dbname=<dbName> password=<dbPassword> sslmode=disable
//
// The connection string is used to establish a connection to the Postgres database.
func BuildPgDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		Settings.DBHost, Settings.DBPort, Settings.DBUser, Settings.DBName, Settings.DBPassword)
}

// BuildRedisDsn builds a connection string for connecting to a Redis database.
// The connection string is in the format:
//
//	redis://:<redisPassword>@<redisHost>:<redisPort>/<redisDB>
//
// The connection string is used to establish a connection to the Redis database.
func BuildRedisDsn() string {
	return fmt.Sprintf("redis://:%s:@%s:%d/%d",
		Settings.RedisPassword, Settings.RedisHost, Settings.RedisPort, Settings.RedisDB)
}
