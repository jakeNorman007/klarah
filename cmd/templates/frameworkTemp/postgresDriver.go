package frameworkTemp

import _"embed"

type PostgresqlTemplate struct{}

var postgresqlServiceTemplate []byte

var postgresqlEnvTemplate []byte

func (m PostgresqlTemplate) Service() []byte {
	return postgresqlServiceTemplate
}

func (m PostgresqlTemplate) Env() []byte {
	return postgresqlEnvTemplate
}
