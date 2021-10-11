package core
/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

type ResultInfo struct {
	AllFileCount int64              `json:"AllFileCount"`// 总文件数
	ValidFileCount int64            `json:"ValidFileCount"`// 有效文件数
	VaildLineCount int64            `json:"VaildLineCount"`// 有效文件行数
	EmptyLineCount int64            `json:"EmptyLineCount"`// 空格数
	SingleCommentsLineCount int64   `json:"SingleCommentsLineCount"`// 注释行数
	MulitCommentsLineCount  int64   `json:"MulitCommentsLineCount"`// 注释行数
}
