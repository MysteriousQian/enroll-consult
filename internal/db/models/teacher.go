package models

// 用户管理
type Teacher struct {
	Id            int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT;comment:账号ID"`
	Name          string `json:"name" gorm:"type:varchar(150);default '';comment:账号名称;"`
	Avatar        string `json:"avatar" gorm:"type:varchar(255);default '';comment:头像"`
	Subject       string `json:"subject" gorm:"type:varchar(100);not null;comment:领域"`
	Qualification string `json:"qualification" gorm:"type:varchar(100);not null;comment:学历"`
	Description   string `json:"description" gorm:"type:varchar(200);comment:简介"`
	CreateTime    int64  `json:"create_time" gorm:"type:bigint unsigned;default:0;not null;comment:创建时间"`
	UpdateTime    int64  `json:"update_time" gorm:"type:bigint unsigned;default:0;not null;comment:更新时间"`
}

var TeacherField = []string{
	"name",
	"avatar",
	"subject",
	"qualification",
	"description",
	"update_time",
}

// 获取表名
func (Teacher) TableName() string {
	return "ec_teacher"
}

func (model Teacher) SelectTeachers(page, size int, name string) (teacherList []Teacher, err error) {
	tx := DB.Model(&model)
	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}
	err = tx.Order("create_time DESC").
		Offset((page - 1) * size).
		Limit(size).
		Scan(&teacherList).Error
	return
}

func (model Teacher) Create() (err error) {
	return DB.Model(&model).Create(&model).Error
}

func (model Teacher) Update() (err error) {
	return DB.Model(&model).Select(TeacherField).Updates(model).Error
}

func (model Teacher) Delete() (err error) {
	return DB.Model(&model).Where("id = ?", model.Id).Delete(&model).Error
}
