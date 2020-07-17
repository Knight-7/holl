//Author  :     knight
//Date    :     2020/07/12 21:14:54
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     图片

package model

//Image 图片
type Image struct {
	OrderID   *int64  `gorm:"primary_key;column:order_id" json:"order_id"`
	ImageName *string `gorm:"primary_key;column:image_name" json:"image_name"`
}

//TableName 设置表名
func TableName() string {
	return "images"
}
