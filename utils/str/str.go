package str

import "errors"

func CharAt(str string, pos int) (string, error) {
	if pos >= len(str) {
		return "", errors.New("invalid index")
	}

	return string(str[pos]), nil
}

func ReplaceCharAt(str string, replacement string, index int) string {
	ret := ""

	for i := 0; i < len(str); i++ {
		if i != index {
			ret += string(str[i])
		} else {
			ret += replacement
		}
	}

	return ret
}

func Reverse(str string) string {
	ret := ""

	for i := len(str) - 1; i >= 0; i-- {
		ret += string(str[i])
	}

	return ret
}
