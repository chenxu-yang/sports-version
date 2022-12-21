package model

import "wxcloudrun-golang/internal/pkg/db"

type Collect struct {
	ID        int32 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OpenId    int32 `json:"user_id" gorm:"column:user_id;type:int(11);not null;default:0;comment:'用户id'"`
	VideoId   int32 `json:"video_id" gorm:"column:video_id;type:int(11);not null;default:0;comment:'视频id'"`
	CreatedAt int32 `json:"created_at" gorm:"column:created_at;type:int(11);not null;default:0;comment:'创建时间'"`
}

func (obj *Collect) TableName() string {
	return "t_collect"
}

func (obj *Collect) Create(collect *Collect) (*Collect, error) {
	err := db.Get().Create(collect).Error
	return collect, err
}

func (obj *Collect) Get(collect *Collect) (*Collect, error) {
	result := new(Collect)
	err := db.Get().Table(obj.TableName()).Where(collect).First(result).Error
	return result, err
}

func (obj *Collect) Gets(collect *Collect) ([]Collect, error) {
	results := make([]Collect, 0)
	err := db.Get().Table(obj.TableName()).Where(collect).Find(&results).Error
	return results, err
}

func (obj *Collect) Update(collect *Collect) (*Collect, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", collect.ID).Updates(collect).Error
	return collect, err
}

func (obj *Collect) Delete(collect *Collect) error {
	return db.Get().Delete(collect, "id = ?", collect.ID).Error
}
