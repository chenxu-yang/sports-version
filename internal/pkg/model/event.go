package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Event struct {
	ID          int32
	UserOpenID  int32
	CourtID     int32
	StartTime   time.Time
	EndTime     time.Time
	CreatedTime time.Time
	UpdatedTime time.Time
}

// TableName get sql table name.获取数据库名字
func (obj *Event) TableName() string {
	return "t_event"
}

// Create 创建记录
func (obj *Event) Create(event *Event) (*Event, error) {
	err := db.Get().Create(event).Error
	return event, err
}

// Get 获取
func (obj *Event) Get(event *Event) (*Event, error) {
	result := new(Event)
	err := db.Get().Table(obj.TableName()).Where(event).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *Event) Gets(event *Event) ([]Event, error) {
	results := make([]Event, 0)
	err := db.Get().Table(obj.TableName()).Where(event).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Event) Update(event *Event) (*Event, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", event.ID).Updates(event).Error
	return event, err
}

// Delete 删除
func (obj *Event) Delete(event *Event) error {
	return db.Get().Delete(event, "id = ?", event.ID).Error
}
