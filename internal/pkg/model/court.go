package model

import "wxcloudrun-golang/internal/pkg/db"

type Court struct {
	ID       int32
	Name     string
	Location string
	Info     string
}

// TableName get sql table name.获取数据库名字
func (obj *Court) TableName() string {
	return "t_court"
}

// Create 创建记录
func (obj *Court) Create(count *Court) (*Court, error) {
	err := db.Get().Create(count).Error
	return count, err
}

// Get 获取
func (obj *Court) Get(count *Court) (*Court, error) {
	result := new(Court)
	err := db.Get().Table(obj.TableName()).Where(count).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *Court) Gets(count *Court) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Court) Update(count *Court) (*Court, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", count.ID).Updates(count).Error
	return count, err
}

// Delete 删除
func (obj *Court) Delete(count *Court) error {
	return db.Get().Delete(count, "id = ?", count.ID).Error
}
