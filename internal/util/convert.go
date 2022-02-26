package util

func FromBoolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
func FromIntToBool(i int) bool {
	return i >= 1
}
