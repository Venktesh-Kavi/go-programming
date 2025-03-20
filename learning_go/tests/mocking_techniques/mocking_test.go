package mocking_techniques

import (
	"database/sql"
	"errors"
	"io"
	"reflect"
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

func TestReadContents(t *testing.T) {
	subtests := []struct {
		name         string
		rc           io.ReadCloser
		expectedData []byte
		expectedErr  error
		numOfBytes   int
	}{
		{
			name: "happy path read",
			rc: mockReadCloser{
				expectedData: []byte("hello world"),
				expectedErr:  nil,
			},
			expectedData: []byte("hello world"),
			expectedErr:  nil,
			numOfBytes:   50,
		},
		{
			name: "read with constrained byte capacity",
			rc: mockReadCloser{
				expectedData: []byte("hello world"),
				expectedErr:  nil,
			},
			expectedData: []byte("hello world"),
			expectedErr:  nil,
			numOfBytes:   2,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			actualData, actualErr := ReadContents(subtest.rc, subtest.numOfBytes)
			if reflect.DeepEqual(string(actualData), subtest.expectedData) {
				t.Errorf("ReadContents() actualData = %v, expectedData = %v", actualData, subtest.expectedData)
			}
			if !errors.Is(actualErr, subtest.expectedErr) {
				t.Errorf("ReadContents() actualErr = %v, expectedErr = %v", actualErr, subtest.expectedErr)
			}
		})
	}
}

// mockReadCloser mock type which implements io.ReadCloser interface. It has state so that each mock can provided an expected data & err.
type mockReadCloser struct {
	expectedData []byte
	expectedErr  error
}

func (mr mockReadCloser) Read(p []byte) (n int, err error) {
	// copy the expected content from mr.expectedData to p
	copy(mr.expectedData, p)
	return 0, mr.expectedErr
}

func (mr mockReadCloser) Close() error {
	// NoOp
	return nil
}
