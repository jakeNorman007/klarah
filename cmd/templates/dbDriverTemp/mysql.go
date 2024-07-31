package dbDriverTemp

import _ "embed"

type MysqlTemplate struct{}

//go:embed 
var mysqlServiceTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}
