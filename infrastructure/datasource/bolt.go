package datasource

import (
	"go.etcd.io/bbolt"
	"time"
)

// BoltDB must defer db.Close()
func BoltDB(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func BoltDBReadOnly(path string) (*bbolt.DB, error) {
	db, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: 3 * time.Second, ReadOnly: true})
	if err != nil {
		return nil, err
	}
	return db, nil
}
