package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/m0lson84/story-tracker/db"
	"github.com/m0lson84/story-tracker/internal/config"
	"github.com/m0lson84/story-tracker/internal/logging"
	"go.uber.org/zap"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Close closes the database connection.
	// It logs a message indicating the disconnection from the specific database.
	// If the connection is successfully closed, it returns nil.
	// If an error occurs while closing the connection, it returns the error.
	Close() error
	// Connect establishes a connection to the database.
	// It returns an error if the connection cannot be established.
	Connect() error
	// Execute returns the database query interface.
	// It is used to execute queries on the database.
	Execute() *db.Queries
	// Health checks the health of the database connection by pinging the database.
	// It returns a map with keys indicating various health statistics.
	Health() map[string]string
}

// DB represents a database service.
type DB struct {
	// The connection to the database.
	conn *pgx.Conn
	// The configuration for the database connection.
	config *pgx.ConnConfig
	// Application logging instance.
	logger *zap.SugaredLogger
	// The database query interface.
	query *db.Queries
}

// NewDatabase creates a new database service.
func NewDatabase(config config.DB) Service {
	logger := logging.NewNamed("Services.DB")
	connString := config.DataSource()

	cfg, err := pgx.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	cfg.Tracer = newTracer(logger)

	return &DB{
		config: cfg,
		logger: logger,
	}
}

func (s *DB) Close() error {
	s.logger.Infow("Disconnected from database...", "database", s.config.Database)
	ctx := context.Background()
	return s.conn.Close(ctx)
}

func (s *DB) Connect() error {
	s.logger.Infow("Connecting to database...", "database", s.config.Database)
	ctx := context.Background()

	conn, err := pgx.ConnectConfig(ctx, s.config)
	if err != nil {
		s.logger.Panicw("Failed to connect to database", "error", err)
	}

	s.conn = conn
	s.query = db.New(s.conn)

	return nil
}

func (s *DB) Execute() *db.Queries {
	return s.query
}

func (s *DB) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.conn.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		s.logger.Errorw("Database is down", "error", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	return stats
}
