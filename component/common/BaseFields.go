package common

import "time"

type BaseFields struct {
	Id          *int64     `gorm:"id;primary;not null" json:"id,omitempty"`
	UtcCreate   *time.Time `gorm:"utc_create;null" json:"utcCreate,omitempty"`
	UtcModified *time.Time `gorm:"utc_modified;null" json:"utcModified,omitempty"`
}
