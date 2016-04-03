package helpers

import (
	"github.com/Cobblestone-Bridge/TurtleOverlord/models"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Login function
func Login(c *mgo.Database, email string, password string) (user *models.User, err error) {
	err = c.C("users").Find(bson.M{"e": email}).One(&user)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		user = nil
	}
	return
}
