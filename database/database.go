package database

import (
	"github.com/simple-inmemory-db/tables"
	"time"
)

type database struct {
	name     string
	tables   []tables.Table
	metaData *dbMeta
}

type dbMeta struct {
	createdDate      string
	lastModifiedDate string
	numTables        int64
	size             int64
}

func NewDB() *database {
	currentTime := time.Now()
	db := &database{
		tables: []tables.Table{},
		metaData: &dbMeta{
			createdDate:      currentTime.Format("2006/01/20 15:04:05"),
			lastModifiedDate: currentTime.Format("2006/01/20 15:04:05"),
			numTables:        0,
			size:             0,
		},
	}

	return db
}
