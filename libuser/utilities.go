// Contains various utility functions
package libuser

import (
	"fmt"
	"net/http"
	"regexp"
)

func VerifyEmail(email string) (bool, error) {
	r := regexp.MustCompile(`(?i)^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$`)
	matched := r.MatchString(email)
	if !matched {
		return false, nil
	}

	return true, nil
}

func HandleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "An error occured: %v", err)
}
