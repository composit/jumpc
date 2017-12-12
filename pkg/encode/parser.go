package encode

import (
	"fmt"
	"regexp"
)

func GetPwd(input string) (string, error) {
	r := regexp.MustCompile("password=(.*)")
	matches := r.FindStringSubmatch(input)

	if len(matches) != 2 {
		return "", fmt.Errorf("improperly formatted input: %s", input)
	}

	return matches[1], nil
}
