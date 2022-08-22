package utils

// OutRange 判断index是否出界
func OutRange(s [][]byte, idx int) bool {
	return idx < len(s)
}
