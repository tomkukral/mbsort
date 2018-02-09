package main

// find position of string in array
func getPosition(s *[]string, value string) int {
	for i, v := range *s {
		if v == value {
			return i
		}
	}

	return -1
}
