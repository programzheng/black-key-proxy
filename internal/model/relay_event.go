package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type RelayEvent struct {
	gorm.Model
	RelayID    uint
	Relay      Relay
	Identifier sql.NullString `gorm:"index:relay_event_identifier_key_key,unique"`
	Key        string         `gorm:"index:relay_event_identifier_key_key,unique"`
}
