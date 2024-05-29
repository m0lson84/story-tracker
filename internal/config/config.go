package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
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
func NewConfig() Config {
	// Load the current environment
	root := getRoot()
	loadEnv(root)

	// Get the configuration values
	app := getAppConfig()
	db := getDatabaseConfig()

	return Config{
		App: app,
		DB:  db,
	}
}

func getAppConfig() App {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	return App{
		Env:  os.Getenv("APP_ENV"),
		Port: port,
	}
}

func getDatabaseConfig() DB {
	return DB{
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func getRoot() string {
	current, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goMod := filepath.Join(current, "go.mod")
		if _, err := os.Stat(goMod); err == nil {
			break
		}

		parent := filepath.Dir(current)
		if parent == current {
			panic("go.mod not found")
		}

		current = parent
	}

	return current
}

func loadEnv(root string) {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	if env == "testing" {
		godotenv.Load(root + "/test/.env.testing")
	}

	godotenv.Load(root + "/.env." + env)
}
