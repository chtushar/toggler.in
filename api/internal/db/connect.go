package db

import (
	"sync"
)

// Creates a new connection to the database.
func NewConnection()  {

}

var (
	once   sync.Once
)

// Creates a new connection to the database if not present or returns the
// existing connection.
func GetConnection()  {
	once.Do(func() {
		NewConnection()
	})
}

// Connecting to the database.
func connect()  {

}