package members

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	DB *gorm.DB
}

func (database *Database) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	database.DB, err = gorm.Open(Dbdriver, DBURL)
	defer database.DB.Close()
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	database.DB.AutoMigrate(&Member{})

	memberOne := Member{Name: "Ariel", Email: "ariel@ariel", Surname: "Silva", ProfilePic: "meu documentos"}
	_, err = memberOne.SaveMember(database.DB)
	if err != nil {
		print(err)
	}

	memberTwo := Member{Name: "Priscila", Email: "priscila@priscil", Surname: "Silva", ProfilePic: "meu computador"}
	_, err = memberTwo.SaveMember(database.DB)
	if err != nil {
		print(err)
	}

	_, err = memberOne.FindAllMembers(database.DB)
	if err != nil {
		print(err)
	}
	mFound, err := memberOne.FindMemberByID(database.DB, 10)
	print(mFound)

}
