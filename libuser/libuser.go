// Libuser contains functionality for user management
package libuser

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
	"code.google.com/p/go.crypto/bcrypt"
)

// Define user type
type User struct {
	UUID     string `db:"uid"`
	Username string
	Email    string
	// @TODO: Ensure this is never sent in a response.
	Password  []byte `json:"-"`
	Verified  bool
	Enabled   bool
	Created   time.Time
	Updated   time.Time
	LastLogin time.Time `db:"last_login"`
}

type UserStorage interface {
	// Retrieve a user.
	// @TODO: Potentially use a paratially filled in user object to retrieve the
	// full user object. This would allow retrieval via usernam, email, or UUID
	RetrieveUser(username string) (User, error)
	CreateUser(user User) (bool, error)
	UpdateUser(user User) (bool, error)
	DeleteUser(user User) (bool, error)
}

/// CreateUser creates an empty user and returns it.
func CreateUser() User {
	id := uuid.New()
	return User{
		UUID:    id,
		Created: time.Now(),
		Updated: time.Now(),
	}
}

// Clears out the memory that held the plaintext password
// http://stackoverflow.com/a/19828153
func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

// SetPassword takes a plaintext passwords and encrypts it.
// Call this function, do not directly set the password on the user struct.
func (u *User) SetPassword(password []byte) (bool, error) {
	defer clear(password)
	passwordHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	u.Password = passwordHash
	return true, nil
}
