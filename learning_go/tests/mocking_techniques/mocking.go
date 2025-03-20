package mocking_techniques

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
)

type sqlOpener func(string, string) (*sql.DB, error)

// OpenDB Higher Order Functions
func OpenDB(user, password, addr, db string, open sqlOpener) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return open("mysql", conn)
}

var MkySqlOpener = sql.Open

// OpenDBMonkeyPatch OpenDB my monkey patching.
func OpenDBMonkeyPatch(user, password, addr, db string) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return MkySqlOpener("mysql", conn)
}

// Mocking using interfaces, typical file reading use case.

func FileNotReadToCapError(msg string) error {
	return errors.New(msg)
}

func FileIOError(msg string) error {
	return errors.New(msg)
}

func ReadFile(fn string, numOfBytes int) error {
	f, err := os.Open(fn)
	if err != nil {
		return FileIOError("error opening file")
	}
	fi, err := f.Stat()
	data, err := ReadContents(f, numOfBytes)
	if len(data) != int(fi.Size()) {
		return FileNotReadToCapError("file not fully read")
	}
	if err != nil {
		return FileIOError("error reading contents from file")
	}
	fmt.Printf("%s\n", string(data))
	return nil
}

func ReadContents(r io.ReadCloser, numOfByte int) (bytes []byte, error error) {
	defer r.Close()
	data := make([]byte, numOfByte)
	_, err := r.Read(data)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return nil, err
	}
	return data, nil
}
