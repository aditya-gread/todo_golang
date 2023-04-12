package utils

import "strconv"

func GetParamInINT(ID string) int {
	int_id, _ := strconv.Atoi(ID)
	return int_id
}
