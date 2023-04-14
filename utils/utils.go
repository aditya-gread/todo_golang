package utils

import "strconv"

// converts string into int
func GetParamInINT(ID string) int {
	int_id, _ := strconv.Atoi(ID)
	return int_id
}
