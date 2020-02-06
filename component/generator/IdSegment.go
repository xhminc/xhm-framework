package generator

type idSegment struct {
	BizTag *string `gorm:"biz_tag"`
	MaxId  *int64  `gorm:"max_id"`
	Step   *int    `gorm:"step"`
}
