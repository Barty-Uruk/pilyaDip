package models

import (
	"time"

	"github.com/go-pg/pg"
)

var (
	db *pg.DB
)

type ormStruct struct {
	Deleted   bool      `json:"-" sql:"default:false"`
	ID        int       `json:"id" sql:",pk"`
	CreatedAt time.Time `sql:"default:now()" json:"created_at" description:"Дата создания"`
	UpdatedAt time.Time `json:"updated_at" `
}

// ConnectDB initialize connection to package var
func ConnectDB(conn *pg.DB) {
	db = conn
}
