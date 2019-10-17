package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/xorm"
	Users "github.com/projects/vue-go-1st/api/models/users"
	"log"
	"time"
)

var signKey string

func init() {
	signKey = "MySecret" /*os.Getenv("SECRET")*/
}

func CreateToken(id int64) (tokenString string, err error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = id

	token.Claims = claims

	tokenString, err = token.SignedString([]byte(signKey))

	return
}

func ParseToken(val string) (id int64, err error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}
		id = int64(claims["id"].(float64))
		return

	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token has expired")
		default:
			return 0, errors.New("error to parse token or token does not exist")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}

func GetUserFromToken(db *xorm.Engine, tokenVal string) (user Users.User, err error) {
	if len(tokenVal) == 0 {
		err = errors.New("No token present")
		return
	}
	userId, err := ParseToken(tokenVal)
	if err != nil {
		log.Println("ParseToken error")
		log.Println(err)
		err = errors.New("Token is invalid")
		return
	}
	if userId < 1 {
		err = errors.New("Token is missing required data.")
		return
	}
	user.ID = userId
	users, err := Users.Index(db, &user)
	if err != nil || len(users) == 0 {
		log.Println(err)
		err = errors.New("Unable to get user from token")
		return
	}
	user = users[0]
	user.Password = ""

	return
}
