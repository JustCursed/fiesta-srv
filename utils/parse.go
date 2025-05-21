package utils

import "strings"

func IfLoc(loc string) bool {
	return strings.HasPrefix(loc, "!")
}
