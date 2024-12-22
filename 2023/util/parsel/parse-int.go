package parsel

import "strconv"

func ParseInt(value string) (int, error) {
	return strconv.Atoi(value)
}
