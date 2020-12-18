package db

import (
	"github.com/jackc/pgx/v4"
)

var (
	Pool *pgx.Conn
)
