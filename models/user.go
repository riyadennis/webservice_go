package models

type User struct {
	Id       int
	Name     string
	Password string
	Role     string
}

func (u User) Save() {

}
