package models

type Province struct {
	Name string `json:"province"  gorm:"type:varchar(10); not null; primary_key; comment:省份"`
}

// 获取表名
func (Province) TableName() string {
	return "ec_province"
}
