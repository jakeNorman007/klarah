package dbDriverTemp

import _ "embed"

type SqliteTemplate struct{}

//go:embed files/db/sqlite.tmpl
var SqliteServiceTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return SqliteServiceTemplate
}
