package main

import (
	"fmt"
	"os/user"
	"strings"
)

// find position of string in array
func getPosition(s []string, value string) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}

	return 0
}

// remove element from slice
func removeItem(s []string, value string) []string {
	ret := []string{}
	for _, i := range s {
		if i != value {
			ret = append(ret, i)
		}
	}

	return ret
}

// replace $HOME with homedir
func expandPath(path string) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Reading user information: %s", err)
	}
	return strings.Replace(path, "$HOME", user.HomeDir, -1), nil
}
