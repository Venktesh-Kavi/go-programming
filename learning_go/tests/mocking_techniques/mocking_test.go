package mocking_techniques

import (
	"database/sql"
	"errors"
	"testing"
)

func TestOpenDB(t *testing.T) {
	var mockErr error
	mockErr = errors.New("mock err")
	tests := []struct {
		name        string
		user        string
		db          string
		password    string
		addr        string
		open        sqlOpener
		expectedErr error
	}{
		{
			name:     "happy path open sql",
			user:     "user",
			db:       "db",
			password: "password",
			addr:     "localhost:3306",
			open: func(t string, conn string) (*sql.DB, error) {
				if conn != "user:password@localhost:3306/db" {
					return nil, errors.New("invalid connection")
				}
				return nil, nil
			},
		},
		{
			name: "connection failure",
			user: "user",
			db:   "db",
			addr: "localhost:3306",
			open: func(t string, conn string) (*sql.DB, error) {
				return nil, mockErr
			},
			expectedErr: mockErr,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenDB(tt.user, tt.password, tt.addr, tt.db, tt.open)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("OpenDB() error = %v, wantErr %v\n", err, tt.expectedErr)
			}
		})
	}

	// Monkey Patch
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MkySqlOpener = tt.open
			_, err := OpenDB(tt.user, tt.password, tt.addr, tt.db, MkySqlOpener)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("OpenDB() error = %v, wantErr %v\n", err, tt.expectedErr)
			}
		})
	}
}
