package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type tracer struct {
	log *zap.SugaredLogger
}

func newTracer(logger *zap.SugaredLogger) *tracer {
	return &tracer{
		log: logger,
	}
}

func (t *tracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	t.log.Debugw("Database query started.", "sql", data.SQL, "args", data.Args)
	return ctx
}

func (t *tracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		t.log.Errorw("Error executing query.", "err", data.Err)
		return
	}
	t.log.Debugw("Database query completed.", "commandTag", data.CommandTag)
}
