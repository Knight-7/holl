//Author  :     knight
//Date    :     2020/07/12 21:14:46
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     订单

package model

//Deal 订单
type Deal struct {
	OrderID     *int64  `gorm:"primarykey;column:order_id" json:"order_id"`
	Order       *Order
	PUblishUser *User  `grom:"foreignkey:PublishID" json:"publish_user"`
	PublishID   *string `gorm:"column:publish_id" json:"publish_id"`
	ReceiveUser *User  `gorm:"foreignkey:ReceiveID" json:"receive_user"`
	ReceiveID   *string `gorm:"column:receive_id" json:"receive_id"`
}

//TableName 自定义表名
func (Deal) TableName() string {
	return "deals"
}
