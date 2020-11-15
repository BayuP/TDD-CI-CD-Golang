package model

import (
	"time"
)

type Base struct {
	CreatedBy    string    `json:"created_by"`
	CreatedTime  time.Time `json:"creted_time"`
	ModifiedBy   string    `json:"modified_by"`
	ModifiedTime time.Time `json:"modified_time"`
	DeletedBy    string    `json:"deleted_by"`
	DeletedTime  time.Time `json:"deleted_time"`
}
