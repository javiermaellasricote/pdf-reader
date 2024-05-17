package file

import "strconv"

func checkByteIsEOL(value byte) bool {
	return value == '\r' || value == '\n'
}

func checkByteIsNumber(value byte) bool {
	_, err := strconv.Atoi(string(value))
	if err != nil {
		return false
	}
	return true
}
