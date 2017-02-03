package countries

import "strconv"

func PadLeftZero(n int, length int) string {
	s := strconv.Itoa(n)

	for len(s) < length {
		s = "0" + s
	}

	return s
}
