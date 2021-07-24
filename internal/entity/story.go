package entity

import "database/sql"

// Story is ...
type Story struct {
	ID        int64
	Author    sql.NullString
	Title     sql.NullString
	Story     sql.NullString
	Status    sql.NullString
	Email     sql.NullString
	CreatedOn sql.NullTime
}
