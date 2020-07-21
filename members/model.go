package members

type Member struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
}
