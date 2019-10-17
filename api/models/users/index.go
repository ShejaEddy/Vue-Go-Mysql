package users

import (
	"github.com/go-xorm/xorm"
)

func Index(db *xorm.Engine, findBy *User) (users Users, err error) {

	if findBy == nil {
		findBy = &User{}
	}

	sess := db.Where("1 = 1")

	if findBy.ID > 0 {
		sess = sess.Where("id = ?", findBy.ID)
	}

	if len(findBy.Username) > 0 {
		sess = sess.Where("username = ?", findBy.Username)
	}

	if err = sess.Find(&users); err != nil {
		return
	}
	return
}
