package dbDriverTemp

import _ "embed"

type SqliteTemplate struct{}

//go:embed 
var SqliteServiceTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return SqliteServiceTemplate
}
