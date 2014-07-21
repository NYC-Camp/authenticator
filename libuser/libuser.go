// Libuser contains functionality for user management
package libuser

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// Define user type
type User struct {
	Username  string
	Email     string
	UUID      uuid.UUID
	Verified  bool
	Created   time.Time
	Updated   time.Time
	LastLogin time.Time
}
