package models

type Major struct {
	Id          int64   `json:"id" gorm:"primary_key;AUTO_INCREMENT;comment:专业ID"`
	Name        string  `json:"name" gorm:"type:varchar(100);comment:名称"`
	Popularity  float64 `json:"popularity" gorm:"type:decimal(10,2);comment:热度"`
	EmployeDest string  `json:"employe_dest" gorm:"type:varchar(255);comment:就业方向"`
	WorkRate    float64 `json:"work_rate" gorm:"type:decimal(10,2);comment:就业率"`
	StudyRate   float64 `json:"study_rate" gorm:"type:decimal(10,2);comment:考研率"`
	Description string  `json:"description" gorm:"type:varchar(255);comment:描述"`
	CreateTime  int64   `json:"create_time" gorm:"type:bigint;comment:创建时间"`
	UpdateTime  int64   `json:"update_time" gorm:"type:bigint;comment:更新时间"`
}

var MajorField = []string{
	"name",
	"popularity",
	"employe_dest",
	"work_rate",
	"study_rate",
	"description",
	"update_time",
}

// 获取表名
func (Major) TableName() string {
	return "ec_major"
}

func (model Major) SelectMajors(page, size int, name string) (majorList []Major, err error) {
	tx := DB.Model(&model)
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	err = tx.Order("create_time DESC").
		Offset((page - 1) * size).
		Limit(size).
		Scan(&majorList).Error
	return
}

func (model Major) Create() (err error) {
	return DB.Model(&model).Create(&model).Error
}

func (model Major) Update() (err error) {
	return DB.Model(&model).Select(MajorField).Updates(model).Error
}

func (model Major) Delete() (err error) {
	return DB.Model(&model).Where("id = ?", model.Id).Delete(&model).Error
}
