package members

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID         primitive.ObjectID `json:"_id, omitempty" bson:"_id, omitempty"`
	Name       string             `json:"name" bson:"name"`
	Surname    string             `json:"surname" bson:"surnname"`
	Email      string             `json:"email" bson:"email"`
	ProfilePic string             `json:"profile_pic" bson:"profile_pic"`
}
