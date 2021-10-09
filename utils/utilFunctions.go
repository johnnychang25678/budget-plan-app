package utils

func IsInSliceInt(sl []int, target int) bool {
	for _, v := range sl {
		if target == v {
			return true
		}
	}
	return false
}
