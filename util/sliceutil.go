package util

import "strconv"

func SliceInt2String1(s []int) string {
	if len(s) < 1 {
		return ""
	}

	ss := strconv.Itoa(s[0])
	for i := 1 ; i < len(s); i++ {
		ss += "," + strconv.Itoa(s[i])
	}
	return ss
}

func SliceInt2String2(s []int) string {
	if len(s) < 1 {
		return ""
	}
	b := make([]byte, 0, 256)
	b = append(b, strconv.Itoa(s[0])...)
	for i := 1; i < len(s); i++ {
		b = append(b, ',')
		b = append(b, strconv.Itoa(s[i])...)
	}
	return string(b)
}
