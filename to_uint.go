package str

import "strconv"

func ToUint32(in string) (uint32, error) {
	out, err := strconv.ParseUint(in, 10, 32)
	return uint32(out), err
}

func ToUint64(in string) (uint64, error) {
	return strconv.ParseUint(in, 10, 64)
}
