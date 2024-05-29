//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

type (
	Gen     mg.Namespace
	DB      mg.Namespace
	Docker  mg.Namespace
	Migrate mg.Namespace
)

// Application
var (
	project = "story-tracker"
	version = "0.1.0"
)

// Build
var (
	outputDir  = "bin"
	binaryName = project
	binaryPath = "./" + filepath.Join(outputDir, binaryName)
)

// Configuration
var (
	envFile = filepath.Join("config", "local.env")
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// Get project dependencies
func Get() error {
	fmt.Println("Downloading dependencies...")
	return sh.RunV("go", "mod", "download")
}

// Build the application
func Build() error {
	mg.Deps(Gen.Sqlc)
	fmt.Println("Building web application...")
	return sh.RunV("go", "build", "-v", "-o", binaryPath, "cmd/api/main.go")
}

// Run the application
func Run() error {
	fmt.Println("Running application...")
	return sh.RunV("go", "run", "cmd/api/main.go")
}

// Live reload
func Watch() error {
	fmt.Println("Running %s in watch mode...", binaryName)
	return sh.RunV("go", "tool", "air")
}

// Clean the binary
func Clean() {
	fmt.Println("Cleaning project...")
	os.RemoveAll(outputDir)
}

// Run test suite
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "tool", "gotestsum", "./...", "-v")
}

// Run test suite with code coverage
func Coverage() error {
	fmt.Println("Running tests with coverage...")
	return sh.RunV("go", "tool", "gotestsum", "./...", "-coverprofile=coverage.out")
}

// Run static analysis
func Lint() error {
	return sh.RunV("go", "vet", "./...")
}

// Run code formatting
func Fmt() error {
	return sh.RunV("go", "fmt", "./...")
}

// sqlc code generation
func (Gen) Sqlc() error {
	fmt.Println("Generating code from SQL...")
	return sh.RunV("sqlc", "generate")
}

// Build docker image
func (Docker) Build() error {
	fmt.Println("Building docker image...")
	return sh.RunV("docker", "build", "-t", project, "-f", "build/Dockerfile", ".")
}

// Run docker container
func (Docker) Run() error {
	fmt.Println("Running containerized application...")
	return sh.RunV("docker", "run", "-it", "-p", "8080:8080", "-t", project+":latest")
}

// Standup local database instance
func (DB) Up() error {
	fmt.Println("Starting database instance...")
	composeFile := filepath.Join("tools", "db", "compose.db.yml")
	return sh.RunV("docker", "compose", "-f", composeFile, "--env-file", envFile, "up", "-d")
}

// Shutdown local database instance
func (DB) Down() error {
	fmt.Println("Shutting down database...")
	composeFile := filepath.Join("tools", "db", "compose.db.yml")
	return sh.RunV("docker", "compose", "-f", composeFile, "down")
}

// Execute pending database migrations
func (Migrate) Up() error {
	fmt.Println("Running database migrations...")
	return sh.RunV("dbmate", "--env-file", envFile, "--migrations-dir", "db/schema", "--migrations-table", "migrations", "up")
}

// Revert the last database migration
func (Migrate) Down() error {
	fmt.Println("Reverting last database migration...")
	return sh.RunV("dbmate", "--env-file", envFile, "--migrations-dir", "db/schema", "--migrations-table", "migrations", "down")
}
