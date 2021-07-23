package function

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-06-01 18:42:41
 */

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

/*
* description: 获取绝对路径（不兼容于go run main.go运行模式）
* author: shahao
* created on: 2021/3/5 11:35
* param param_1:
* param param_2:
* return return_1:
 */
func GetAbsPath(relativePath string) string {
	execPath, _ := os.Executable()
	path, _ := filepath.Split(execPath)
	if relativePath == "" {
		return ""
	}
	//兼容go run main.go模式，请在开发模式下使用，生产环境打包请注释掉
	if gin.Mode() == gin.DebugMode {
		path, _ = os.Getwd()
	}
	return filepath.Join(path, relativePath)
}

/*
* description: 获取当前执行程序绝对路径（不兼容于go run main.go运行模式）
* author: shahao
* created on: 2021/3/5 11:36
* param param_1:
* param param_2:
* return return_1:
 */
func GetCurrentAbsPath() string {
	execPath, _ := os.Executable()
	path, _ := filepath.Split(execPath)
	//兼容go run main.go模式，请在开发模式下使用，生产环境打包请注释掉
	if gin.Mode() == gin.DebugMode {
		path, _ = os.Getwd()
	}
	return path
}

/*
* description: 检测路径不存在则创建
* author: shahao
* created on: 20-3-23 下午5:22
* param param_1:
* param param_2:
* return return_1:
 */
func FileNotExistsAndCreate(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}
