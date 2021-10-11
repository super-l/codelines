package utils
/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

import "strings"

func ToLinuxPath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

// 判断是否是MAC隐藏样式文件
func IsMacStyleFile(filename string) bool {
	if filename[0:2] == "._" {
		return true
	}
	return false
}

//获取得文件的扩展名，最后一个.后面的内容
func GetExt(f string) (ext string) {
	index := strings.LastIndex(f, ".")
	data := []byte(f)
	for i := index + 1; i < len(data); i++ {
		ext = ext + string([]byte{data[i]})
	}
	return
}
