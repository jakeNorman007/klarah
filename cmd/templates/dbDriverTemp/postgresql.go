package dbDriverTemp

import _"embed"

type PostgresqlTemplate struct{}

//embed here
var postgresqlServiceTemplate []byte

var postgresqlEnvTemplate []byte

func (m PostgresqlTemplate) Service() []byte {
    return postgresqlServiceTemplate
}

func (m PostgresqlTemplate) Env() []byte {
    return postgresqlEnvTemplate
}
