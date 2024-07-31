package flags

import (
    "fmt"
    "strings"
)

type Database string

const (
    Postgresql Database = "postgresql"
    Sqlite Database = "sqlite"
)

var DatabaseTypes = []string{string(Postgresql), string(Sqlite)}

func (d Database) String() string {
    return string(d)
}

func (d *Database) Type() string {
    return "Database"
}

func (d *Database) Set(value string) error {
    for _, datbse := range DatabaseTypes {
        if datbse == value {
            *d = Database(value)
            return nil
        }
    }

    return fmt.Errorf("Databases available to use: %s", strings.Join(DatabaseTypes, ", "))
}
