package imdb

import (
	"errors"
)

type value struct {
	data []byte
}

type database struct {
	entries map[string]value
}

func NewDB() *database {

	db := &database{
		entries: make(map[string]value),
	}

	return db
}

func (db *database) Get(key string) (value, error) {
	if val, ok := db.entries[key]; ok {
		return val, nil
	}

	return value{}, errors.New("Key not found in database")
}

func (db *database) Set(key string, val []byte) error {
	if key == "" {
		return errors.New("Key to be inserted should be of atleast length 1")
	}
	db.entries[key] = value{
		data: val,
	}

	return nil
}

func (db *database) Unset(key string) {
	db.entries[key] = value{}
}
