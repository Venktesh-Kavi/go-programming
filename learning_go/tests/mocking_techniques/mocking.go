package mocking_techniques

import (
	"database/sql"
	"fmt"
)

type sqlOpener func(string, string) (*sql.DB, error)

// Higher Order Functions
func OpenDB(user, password, addr, db string, open sqlOpener) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return open("mysql", conn)
}
