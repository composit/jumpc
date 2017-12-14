package server

import (
	"fmt"
	"strconv"
)

// validate port performs simple port validation
// it simply checks if the port is numerical
func validatePort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return fmt.Errorf("bad port specified: %s", port)
	}

	return nil
}
