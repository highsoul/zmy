package model

import (
	"github.com/jinzhu/gorm"
)

/***********************
    Gorm 数据库实体
************************/
type Point struct {
	gorm.Model
	Lng float64
	Lat float64
}

/*Point Insert to Sqlite*/
func (p Point) Insert(db gorm.DB) {
	if db.NewRecord(p) {
		db.Create(&p)
	}
}

/* 获取最后一个point记录 */
func (p *Point) GetLast(db gorm.DB) {
	db.Last(p)
}
