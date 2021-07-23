package function

import (
	"crypto/md5"
	"encoding/hex"
)

/*
* description: 计算md5
* author: shahao
* created on: 2021/3/9 19:20
* param param_1:
* param param_2:
* return return_1:
 */
func CreateMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
