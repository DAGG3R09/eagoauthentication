package dao

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Alias    string        `json:"alias" bson:"alias"`
	Username string        `json:"username" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

var userCollection = getUserCollection(session)

// GetAllUsers returns all users from users collection
func GetAllUsers() []user {
	var users []user

	err := userCollection.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}

	return users
}

// GetUserByEmail return the user if exists
func GetUserByEmail(email string) user {
	var user user

	err := userCollection.Find(bson.M{"email": email}).One(&user)

	if err == mgo.ErrNotFound {
		fmt.Println("Not Found.")
		return user
	} else if err != nil {
		panic(err)
	}

	return user
}

// CheckPassword compares the entered password with pass of user
func AuthenticateUser(email string, password string) bool {
	user := GetUserByEmail(email)

	if password == user.Password {
		return true
	}
	return false
}
