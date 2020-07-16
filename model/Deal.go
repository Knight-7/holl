//Author  :     knight
//Date    :     2020/07/12 21:14:46
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     订单

package model

//Deal 订单
type Deal struct {
	OrderID   int64  `gorm:"column:order_id" json:"order_id"`
	PublishID string `gorm:"column:publish_id" json:"publish_id"`
	ReceiveID string `gorm:"column:receive_id" json:"receive_id"`
}

//TableName 自定义表名
func (Deal) TableName() string {
	return "deals"
}
