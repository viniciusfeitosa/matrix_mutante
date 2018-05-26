package db

import (
	redigo "github.com/gomodule/redigo/redis"
)

// MockDB is a mock to DB
type MockDB struct {
	Value string
	Err   []error
}

// NewDBPool the mock interface signature
func (db *MockDB) NewDBPool() *redigo.Pool {
	return nil
}
func (db *MockDB) popErr(errors []error) error {
	var err error
	if len(errors) > 0 {
		if len(errors) > 1 {
			err = errors[0]
			db.Err = errors[1:]
		} else {
			err = errors[0]
		}
	}
	return err
}

// GetValue is a method resposible to get a value from the MockDB using a key
func (db *MockDB) GetValue(key interface{}) (string, error) {
	return db.Value, db.popErr(db.Err)
}

// SetValue is a method resposible to set a value on the MockDB appling a key as reference
func (db *MockDB) SetValue(key interface{}, value interface{}) error {
	return db.popErr(db.Err)
}

// DelValue is a method resposible to delete a value from the MockDB by a key
func (db *MockDB) DelValue(key interface{}) error {
	return db.popErr(db.Err)
}

// EnqueueValue is the responsible to add a value on the queue
func (db *MockDB) EnqueueValue(queue string, uuid uint32) error {
	return db.popErr(db.Err)
}

// PopQueue is the responsible to get a value from the queue
func (db *MockDB) PopQueue(queue string, id int) (uint32, string, error) {
	return 0, db.Value, db.popErr(db.Err)
}
