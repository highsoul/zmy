package model

import (
	"github.com/jinzhu/gorm"
)

/***********************
    Gorm 数据库实体
************************/
type Point struct {
	ID       uint `gorm:"primary_key"`
	CreateAt string
	Lng      float64
	Lat      float64
}

/*Point Insert to Sqlite*/
func (p Point) Insert(db gorm.DB) {
	if db.NewRecord(p) {
		db.Debug().Create(&p)
	}
}

/* 获取指定ID的记录 */
func (p *Point) Get(db gorm.DB, id int) {
	db.First(p, id)
}

/* 获取最后一个point记录 */
func (p *Point) GetLast(db gorm.DB) {
	db.Last(p)
}

/*获取所有记录*/
func (p *Point) GetAll(db gorm.DB) []Point {
	points := []Point{}
	db.Order("create_at desc").Find(&points)
	return points
}

func (p *Point) Delete(db gorm.DB) {
	db.Unscoped().Delete(p)
}
