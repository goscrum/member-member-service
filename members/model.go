package members

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50) json:"name"`
	Surname    string `gorm:"type:varchar(100) json:"surnam"`
	Email      string `gorm:"type:varchar(100);unique_index" json:"email"`
	ProfilePic string `gorm:"type:varchar(200) json:"profile_pic"`
}