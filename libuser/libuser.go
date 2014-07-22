// Libuser contains functionality for user management
package libuser

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// Define user type
type User struct {
	Username string
	Email    string
	UUID     uuid.UUID
	// @TODO: Ensure this is never sent in a response.
	password  []byte
	Verified  bool
	Enabled   bool
	Created   time.Time
	Updated   time.Time
	LastLogin time.Time
}

type UserStorage interface {
	// Retrieve a user.
	// @TODO: Potentially use a paratially filled in user object to retrieve the
	// full user object. This would allow retrieval via usernam, email, or UUID
	RetrieveUser(username string) (User, error)
	CreateUser(user User) (bool, error)
	UpdateUser(user User) (bool, error)
	DisableUser(user User) (bool, error)
	DeleteUser(user User) (bool, error)
}
