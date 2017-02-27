package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(str string) string {
	var returnMD5String string

	hash := md5.New()
	hash.Write([]byte(str))

	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String
}
