package stringutil

import ()

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
