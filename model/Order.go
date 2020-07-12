//Author  :     knight
//Date    :     2020/07/12 00:05:56
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     Order的model

package model

import "time"

//Order 订单
type Order struct {
	ID     int64     `gorm:"primary_key;column:id" json:"id"`
	Location    string    `gorm:"column:location" json:"location"`
	Money       float32   `gorm:"column:money" json:"money"`
	Type        string    `gorm:"column:type;size:30" json:"type"`
	Title       string    `gorm:"column:title" json:"title"`
	Detail      string    `gorm:"column:detal" json:"detail"`
	Phone       string    `gorm:"column:phone;size:11" json:"phone"`
	PublishTime time.Time `gorm:"column:publish_time" json:"publish_time"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	FinishTime  time.Time `gorm:"column:finish_time" json:"finish_time"`
	Image       Image     `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
	Deal        Deal      `gorm:"foreignkey:OrderID;association_foreignkey:ID"`
}

//TableName 设置表名
func (Order) TableName() string {
	return "orders"
}
