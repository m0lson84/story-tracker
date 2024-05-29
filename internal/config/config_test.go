package config

import (
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
	env := NewConfig("testing")

	// ASSERT
	assert.Equal(t, app, env.App, "parse app env variables")
	assert.Equal(t, db, env.DB, "parse db env variables")
}
