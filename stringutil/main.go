package stringutil

func PadRight(str, pad string, length int) string {
	length--
	for {
		if len(str) > length {
			return str
		}
		str += pad
	}
}

func PadLeft(str, pad string, length int) string {
	length--
	for {
		if len(str) > length {
			return str
		}
		str = pad + str
	}
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}
