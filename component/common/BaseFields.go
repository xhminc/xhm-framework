package common

import "time"

type BaseFields struct {
	UtcCreate   *time.Time `gorm:"utc_create;null" json:"utcCreate,omitempty"`
	UtcModified *time.Time `gorm:"utc_modified;null" json:"utcModified,omitempty"`
}
