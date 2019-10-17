package users

type Users []User

type User struct {
	ID       int64  `xorm:"id"`
	Username string `xorm:"username`
	Password string `xorm:"password"`
}

func (u *User) TableName() string {
	return "users"
}
