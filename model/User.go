//Author  :     knight
//Date    :     2020/07/11 23:46:06
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     用户的model

package model

//User 用户
type User struct {
	OpenID       string `gorm:"primary_key;column:open_id;size:30" json:"openid"`
	AvatarURL    string `gorm:"column:avatar_url" json:"avatar_url"`
	City         string `gorm:"column:city;size:30" json:"city"`
	Province     string `gorm:"column:province;size:30" json:"province"`
	Country      string `gorm:"column:country;size:30" json:"country"`
	Credit       int64  `gorm:"column:credit;default:100" json:"credit"`
	Gender       int    `gorm:"column:gender" json:"gender"`
	Language     string `gorm:"column:language;size:30" json:"language"`
	PublishDeals []Deal `gorm:"foreignkey:PublishID;association_foreignkey:OpenID"`
	ReceiveDeals []Deal `gorm:"foreignkey:ReceiveID;association_foreignkey:OpenID"`
}

//TableName 设置表名
func (User) TableName() string {
	return "users"
}
