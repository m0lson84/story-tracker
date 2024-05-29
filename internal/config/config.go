package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

// Config the configuration for the story tracker application
type Config struct {
	// Database related configuration
	DB DB
	// Application related configuration
	App
}

// App the configuration for the application
type App struct {
	// The current execution environment
	Env string
	// The port the server will listen on
	Port int
}

// DB the configuration for the database connection
type DB struct {
	// The name of the database
	Name string
	// The host for connecting to the database
	Host string
	// The port for connecting to the database
	Port string
	// The username to use when connecting to the database
	User string
	// The password credentials for connecting to the database
	Password string
}

// DataSource returns the data source URL for the database connection
func (d DB) DataSource() string {
	return "postgres://" + d.User + ":" + d.Password + "@" + d.Host + ":" + d.Port + "/" + d.Name + "?sslmode=disable"
}

// NewConfig creates a new configuration for the application
func NewConfig(name string) Config {
	env, err := loadEnv(name)
	if err != nil {
		panic(err)
	}

	return Config{
		App: env.app(),
		DB:  env.db(),
	}
}

type env struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	AppPort    string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBUser     string `mapstructure:"DB_USER"`
}

func getRoot() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	current := cwd
	for {
		goMod := filepath.Join(current, "go.mod")
		if _, err := os.Stat(goMod); err == nil {
			break
		}

		parent := filepath.Dir(current)
		if parent == current {
			return cwd
		}

		current = parent
	}

	return current
}

// loadEnv loads the configuration from the environment
func loadEnv(name string) (env, error) {
	env := env{}

	// Assume local if no name is provided
	if name == "" {
		name = "local"
	}

	// Get the root of the project
	root := getRoot()

	// Define where to look for the configuration files
	viper.AddConfigPath(filepath.Join(root, "config"))
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	// Set defaults for the configuration
	setDefaults()

	// Load existing environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	viper.ReadInConfig()

	// Unmarshal the configuration
	if err := viper.Unmarshal(&env); err != nil {
		return env, err
	}

	return env, nil
}

// setDefaults sets the default values for the configuration
func setDefaults() {
	// Application defaults
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("PORT", "8080")

	// Database defaults
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "root")
}

// app get the application configuration
func (e env) app() App {
	port, err := strconv.Atoi(e.AppPort)
	if err != nil {
		panic(err)
	}

	return App{
		Env:  e.AppEnv,
		Port: port,
	}
}

func (e env) db() DB {
	return DB{
		Name:     e.DBName,
		Host:     e.DBHost,
		Port:     e.DBPort,
		User:     e.DBUser,
		Password: e.DBPassword,
	}
}
