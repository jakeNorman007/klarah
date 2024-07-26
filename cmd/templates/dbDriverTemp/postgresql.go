package dbDriverTemp

import (
    _ "embed"
)

type PostgresqlTemplate struct{}

//go:embed files/db/postgresql.tmpl
var postgresqlServiceTemplate []byte

func (m PostgresqlTemplate) Service() []byte {
	return postgresqlServiceTemplate
}
