package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tag is tag model of blog
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// BeforeCreate generates created_on field
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// BeforeUpdate updates modified_on field
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

// GetTags gets tags
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	// db in same models package
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetTagTotal gets total tags
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName checks tag
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// AddTag adds tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// ExistTagByID checks tag by id
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// DeleteTag deletes tag by id
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

// EditTag updates tag by id
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}
