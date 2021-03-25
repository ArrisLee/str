package str

import "strconv"

func ToInt(in string) (int, error) {
	return strconv.Atoi(in)
}

func ToInt32(in string) (int32, error) {
	out, err := strconv.ParseInt(in, 10, 32)
	return int32(out), err
}

func ToInt64(in string) (int64, error) {
	return strconv.ParseInt(in, 10, 64)
}
