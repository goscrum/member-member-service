package members

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50)" json:"name"`
	Surname    string `gorm:"type:varchar(100)" json:"surname"`
	Email      string `gorm:"type:varchar(100); unique_index" json:"email"`
	ProfilePic string `gorm:"type:varchar(200)" json:"profile_pic"`
}

func (u *Member) SaveMember(db *gorm.DB) (*Member, error) {

	err := db.Create(&u).Error
	if err != nil {
		return &Member{}, err
	}
	return u, nil
}

func (u *Member) FindAllMembers(db *gorm.DB) (*[]Member, error) {
	var err error
	members := []Member{}
	err = db.Model(&Member{}).Limit(100).Find(&members).Error
	if err != nil {
		return &[]Member{}, err
	}
	return &members, err
}

func (u *Member) FindMemberByID(db *gorm.DB, uid uint32) (*Member, error) {
	err := db.Model(Member{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Member{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Member{}, errors.New("Member Not Found")
	}
	return u, err
}
