package encode

import (
	"fmt"
	"regexp"
)

func GetPwd(input []byte) ([]byte, error) {
	r := regexp.MustCompile("password=(.*)")
	matches := r.FindSubmatch(input)

	if len(matches) != 2 {
		return []byte{}, fmt.Errorf("improperly formatted input: %s", input)
	}

	return matches[1], nil
}
