package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	// ARRANGE
	app := App{
		Env:  "testing",
		Port: 8080,
	}
	db := DB{
		Name:     "test",
		Host:     "local",
		Port:     "42",
		User:     "foo",
		Password: "12345",
	}

	// ACT
	env := NewConfig()

	// ASSERT
	assert.Equal(t, app, env.App, "Correctly parsed app env variables")
	assert.Equal(t, db, env.DB, "Correctly parsed DB env variables")
}

func TestDBDataSource(t *testing.T) {
	// ARRANGE
	expected := "postgres://foo:12345@local:42/test?sslmode=disable"

	// ACT
	env := NewConfig()

	// ASSERT
	assert.Equal(t, expected, env.DB.DataSource(), "Correctly formatted DB data source")
}

func TestMain(m *testing.M) {
	// Setup
	setup()
	defer teardown()

	// Run tests
	code := m.Run()

	// Return exit code
	os.Exit(code)
}

func setup() {
	os.Setenv("APP_ENV", "testing")
}

func teardown() {
	os.Clearenv()
}
