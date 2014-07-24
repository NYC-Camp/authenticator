// An implementation of user storage in MySQL
package main

import (
	"errors"

	"github.com/coopernurse/gorp"
	"github.com/nyc-camp/authenticator/libuser"
)

const USER_NOT_UPDATED = "userstoragemysql: User Not Updated."
const USER_NOT_DELETED = "userstoragemysql: User not deleted or does not exist."

type UserStorageMySQL struct {
	Dbmap *gorp.DbMap
}

func (usm UserStorageMySQL) RetrieveUser(username string) (libuser.User, error) {
	return libuser.User{}, nil
}

func (usm UserStorageMySQL) CreateUser(user libuser.User) (bool, error) {
	err := usm.Dbmap.Insert(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (usm UserStorageMySQL) UpdateUser(user libuser.User) (bool, error) {
	count, err := usm.Dbmap.Update(user)
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
