package function

import "os"

/*
* description: 判断文件是否存在
* author: shahao
* created on: 2021/3/5 16:50
* param param_1:
* param param_2:
* return return_1:
 */
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
