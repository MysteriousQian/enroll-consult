package models

type Year struct {
	Name string `json:"year" gorm:"type:varchar(10); not null; primary_key; comment:年份"`
}

// 获取表名
func (Year) TableName() string {
	return "ec_year"
}
