package hash

import "golang.org/x/crypto/bcrypt"

// 加密密码
func EncodePassword(pwd []byte) (hashedPwd []byte, err error) {
	hashedPwd, err = bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return
}

// 验证密码 hashedPwd是混淆后的密码
func ComparePassword(hashedPwd, plainPwd []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		return false
	}
	return true
}
