package common

import "time"

type BaseFields struct {
	UtcCreate   *time.Time `gorm:"utc_create;default:null" json:"utcCreate,omitempty"`
	UtcModified *time.Time `gorm:"utc_modified;default:null" json:"utcModified,omitempty"`
}
