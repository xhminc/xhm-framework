package generator

import (
	"fmt"
	"sync"
)

type tag struct {
	lock  sync.Mutex
	name  string
	begin int64
	end   int64
}

func (t *tag) InitTag(tag string) error {

	if len(tag) == 0 {
		return fmt.Errorf("Tag name is empty\n")
	}

	t.name = tag
	if !t.isTagExists() {
		return fmt.Errorf("Tag not exists, tag name: %s\n", t.name)
	}

	err := t.getTagIdSegment()
	if err != nil {
		return err
	}

	return nil
}

func (t *tag) getTagIdSegment() error {

	if t.begin == t.end {

		tx := db.Begin()

		updateRet := tx.Table(tableName).Where("biz_tag = ?", t.name).
			Update("max_id = max_id + step").RowsAffected

		if updateRet == 0 {
			tx.Rollback()
			return fmt.Errorf("Update segment table fail, tag name: %s\n", t.name)
		}

		var s idSegment
		tx.Select("biz_tag, max_id, step").Table(tableName).Where("biz_tag = ?", t.name).Scan(&s)

		if len(*s.BizTag) == 0 {
			tx.Rollback()
			return fmt.Errorf("Select segment table fail, tag name: %s\n", t.name)
		}

		t.end = *s.MaxId
		t.begin = t.end - int64(*s.Step)

		tx.Commit()
	}

	return nil
}

func (t *tag) nextId() (int64, error) {

	t.lock.Lock()
	defer t.lock.Unlock()

	err := t.getTagIdSegment()
	if err != nil {
		return -1, err
	}

	var ret int64
	ret = t.begin
	t.begin += 1

	return ret, nil
}

func (t *tag) isTagExists() bool {
	var count int64
	db.Table(tableName).Where("biz_tag = ?", t.name).Count(&count)
	return count > 0
}
