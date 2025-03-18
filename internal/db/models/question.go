package models

type Question struct {
	Id         int64  `json:"id" gorm:"type:int;primary_key;AUTO_INCREMENT;comment:问题Id"`
	Title      string `json:"title" gorm:"type:varchar(255);comment:问题标题"`
	Content    string `json:"content" gorm:"type:text;comment:问题内容"`
	CreateTime int64  `json:"create_time" gorm:"type:bigint;comment:创建时间"`
	UpdateTime int64  `json:"update_time" gorm:"type:bigint;comment:更新时间"`
}

func (model Question) TableName() string {
	return "ec_question"
}

// 查询问题列表
func (model Question) SelectAllByPage(page, size int, content string) (questions []Question, total int64, err error) {
	tx := DB.Model(&model)
	if content != "" {
		tx.Where("content LIKE ?", "%"+content+"%")
	}
	err = tx.Count(&total).
		Offset((page - 1) * size).
		Limit(size).
		Scan(&questions).Error
	return
}

// 新增问题
func (model Question) Create() (err error) {
	return DB.Create(&model).Error
}

// 删除问题
func (model Question) Delete() (err error) {
	return DB.Model(&model).Where("id = ?", model.Id).Delete(&model).Error
}
