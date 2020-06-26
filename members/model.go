package members

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50)`
	Surname    string `gorm:"type:varchar(100)`
	Email      string `gorm:"type:varchar(100);unique_index"`
	ProfilePic string `gorm:"type:varchar(200)`
}
