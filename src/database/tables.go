package database

import "github.com/scylladb/gocqlx/v2/table"

var ticketMetadata = table.Metadata{
	Name:    "go.tickets",
	Columns: []string{"ticket_id", "price", "movie", "seat", "status", "user_id"},
	PartKey: []string{"ticket_id"},
	SortKey: []string{"ticket_id"},
}

func New() *table.Table {
	return table.New(ticketMetadata)
}
