package circletest

func StrInSlice(a string, s []string) bool {

	isInSlice := false
	for _, v := range s {
		if v == a {
			isInSlice = true
			break
		}
	}
	return isInSlice

}
