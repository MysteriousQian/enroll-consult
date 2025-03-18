package models

type AcceptDetail struct {
	Id           int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;comment:详情ID"`
	Province     string `json:"province" gorm:"type:varchar(100);not null;comment:省份"`
	Year         string `json:"year" gorm:"type:varchar(100);not null;comment:年份"`
	Subject      string `json:"subject" gorm:"type:varchar(100);not null;comment:科类"`
	Batch        string `json:"batch" gorm:"type:varchar(100);not null;comment:专业组/批次"`
	Major        string `json:"major" gorm:"type:varchar(100);not null;comment:专业"`
	AcceptCount  int    `json:"accept_count" gorm:"type:int unsigned;default:0;comment:录取人数"`
	LowestScore  int    `json:"lowest_score" gorm:"type:int;comment:最低分数"`
	LowestRank   int    `json:"lowest_rank" gorm:"type:int unsigned;comment:最低名次"`
	HighestScore int    `json:"highest_score" gorm:"type:int;comment:最高分数"`
	Average      int    `json:"average" gorm:"type:int;comment:平均分"`
	ControlLine  string `json:"control_line" gorm:"type:varchar(100);comment:控制线"`
	Description  string `json:"description" gorm:"type:varchar(200);comment:备注"`
	CreateTime   int64  `json:"create_time" gorm:"type:bigint unsigned;default:0;not null;comment:创建时间"`
	UpdateTime   int64  `json:"update_time" gorm:"type:bigint unsigned;default:0;not null;comment:更新时间"`
}

// 获取表名
func (AcceptDetail) TableName() string {
	return "ec_accept_detail"
}

func (model AcceptDetail) SelectAcceptDetailsList(major, province, year string) (details []AcceptDetail, err error) {
	tx := DB.Model(&model)
	if major != "" {
		tx.Where("major LIKE ?", "%"+major+"%")
	}
	if province != "" {
		tx.Where("province = ?", province)
	}
	if year != "" {
		tx.Where("year = ?", year)
	}
	err = tx.Order("year DESC").Find(&details).Error
	return
}

func (model AcceptDetail) SelectByMajorProvince() (acceptDetail []AcceptDetail, err error) {
	err = DB.Model(&model).
		Where("major = ? AND province = ?", model.Major, model.Province).
		Order("year DESC").
		Find(&acceptDetail).Error
	return acceptDetail, err
}
