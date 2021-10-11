package core
/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

type Config struct {
	CodeFolderPath string      `json:"CodeFolderPath"`    // 代码文件夹根路径
	ValidFileSuffix []string   `json:"ValidFileSuffix"`   // 有效文件后缀
	IgnoreFolderPath []string  `json:"IgnoreFolderPath"`  // 忽略的文件夹目录
	IgnoreFileName   []string  `json:"IgnoreFileName"`    // 忽略的文件完整路径
	IgnoreBlankLines bool      `json:"IgnoreBlankLines"`  // 是否忽略空行
	IgnoreComments   bool      `json:"IgnoreComments"`    // 是否忽略注释
	IgnoreMacStyleFile bool    `json:"IgnoreMacStyleFile"`// 是否忽略MAC系统自动生成的隐藏文件
}



