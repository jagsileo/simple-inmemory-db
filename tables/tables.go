package tables

import (
	"github.com/google/uuid"
	"time"
)

type Table struct {
	ID                uuid.UUID
	Name              string
	Entries           []map[string][]interface{}
	MetaData          *tableMetaData
	LastTransactionID string
}

type tableMetaData struct {
	createdDate  string
	currentSize  int64
	lastModified string
	numEntries   int64
	size         int64
}

func NewTable() *Table {
	currentTime := time.Now()
	table := &Table{
		ID:      uuid.New(),
		Entries: []map[string][]interface{}{},
		MetaData: &tableMetaData{
			createdDate:  currentTime.Format("2019/10/30 15:01:04"),
			currentSize:  0,
			lastModified: currentTime.Format("2019/10/30 15:01:04"),
			numEntries:   0,
			size:         0,
		},
		LastTransactionID: "",
	}

	return table
}
