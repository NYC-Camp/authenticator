// An implementation of user storage in MySQL
package main

import (
	"errors"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/coopernurse/gorp"
	"github.com/nyc-camp/authenticator/libuser"
)

const USER_NOT_UPDATED = "userstoragemysql: User Not Updated."
const USER_NOT_DELETED = "userstoragemysql: User not deleted or does not exist."

type UserStorageMySQL struct {
	Dbmap *gorp.DbMap
}

func (usm UserStorageMySQL) RetrieveUser(username string) (libuser.User, error) {
	var user libuser.User
	err := usm.Dbmap.SelectOne(&user, "SELECT * FROM user WHERE username=?", username)
	if err != nil {
		return libuser.User{}, err
	}

	return user, nil
}

func (usm UserStorageMySQL) CreateUser(user libuser.User) (bool, error) {
	err := usm.Dbmap.Insert(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (usm UserStorageMySQL) UpdateUser(user libuser.User) (bool, error) {
	count, err := usm.Dbmap.Update(&user)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New(USER_NOT_UPDATED)
	}
	return true, nil
}

func (usm UserStorageMySQL) DeleteUser(user libuser.User) (bool, error) {
	count, err := usm.Dbmap.Delete(user)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, errors.New(USER_NOT_DELETED)
	}

	return true, nil
}

func (usm UserStorageMySQL) CheckUsername(username string) (bool, error) {
	count, err := usm.Dbmap.SelectInt("select count(*) from user where username=?", username)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (usm UserStorageMySQL) CheckEmail(email string) (bool, error) {
	count, err := usm.Dbmap.SelectInt("select count(*) from user where email=?", email)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (usm UserStorageMySQL) CheckPassword(username, password string) (bool, error) {
	user, err := usm.RetrieveUser(username)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
