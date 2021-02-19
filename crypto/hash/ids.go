package hash

import "github.com/speps/go-hashids"

// 加密
func EncodeID(salt string, minLength int, params []int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.Encode(params)
		if err == nil {
			return e
		}
	}
	return ""
}

// 解密
func DecodeID(salt string, minLength int, hash string) []int {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err == nil {
		e, err := h.DecodeWithError(hash)
		if err == nil {
			return e
		}
	}
	return []int{}
}
