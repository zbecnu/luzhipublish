package model

type LuzhiPublish struct {
	Cover     string `gorm:"column:Cover"`
	Url       string `gorm:"column:Url"`
	CreatedAt int    `gorm:"column:CreatedAt"`
	Title     string `gorm:"column:Title"`
}

func (LuzhiPublish) TableName() string {
	return "LuzhiPublishes"
}
