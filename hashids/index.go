package hashids

import (
	hashids "github.com/speps/go-hashids/v2"
)

const (
	SALT = "this is my salt"
)

// encode int number to string
func Encode(rawInt int, salt string, minLength int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	if len(salt) == 0 {
		hd.Salt = SALT
	}
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{rawInt})
	return e
}

func Decode(salt string, minLength int) []int {
	hd := hashids.NewData()
	hd.Salt = salt
	if len(salt) == 0 {
		hd.Salt = SALT
	}
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(e)

	return d
}
