//Author  :     knight
//Date    :     2020/07/12 00:05:56
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     Order的model

package model

import "time"

//Order 订单
type Order struct {
	OrderID     *int64      `gorm:"primary_key;column:order_id" json:"id" form:"id"`
	Location    *string     `gorm:"column:location" json:"location" form:"location"`
	Money       *float64    `gorm:"column:money" json:"money" form:"money"`
	Type        *string     `gorm:"column:type;size:30" json:"type" form:"type"`
	Title       *string     `gorm:"column:title" json:"title" form:"title"`
	Detail      *string     `gorm:"column:detail" json:"detail" form:"detail"`
	Phone       *string     `gorm:"column:phone;size:11" json:"phone" form:"phone"`
	PublishTime *time.Time `gorm:"column:publish_time" json:"publish_time" form:"publish_time"`
	StartTime   *time.Time `gorm:"column:start_time" json:"start_time" form:"start_time"`
	FinishTime  *time.Time `gorm:"column:finish_time" json:"finish_time" form:"finish_time"`
	images      []*Image
}

//TableName 设置表名
func (Order) TableName() string {
	return "orders"
}
