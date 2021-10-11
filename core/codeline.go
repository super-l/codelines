package core

/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

import (
	"bufio"
	"fmt"
	"github.com/super-l/codelines/utils"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetResultInfo(config Config) ResultInfo {
	var resultInfo ResultInfo

	// 总文件个数
	resultInfo.AllFileCount = getAllFileCount(config)

	// 统计有效文件个数
	allFiles, err := getVerifyFilesPathList(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	resultInfo.ValidFileCount = int64(len(allFiles))

	resultInfo.VaildLineCount, resultInfo.EmptyLineCount, resultInfo.SingleCommentsLineCount, resultInfo.MulitCommentsLineCount = getVerifyFileListInfo(allFiles)

	return resultInfo
}


// 总文件个数
func getAllFileCount(config Config) int64 {
	var allFileCount = int64(0)
	baseFolder := utils.ToLinuxPath(config.CodeFolderPath)
	filepath.Walk(baseFolder, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		allFileCount ++
		return nil
	})
	return allFileCount
}

// 递归获取某个目录下的所有文件
func getVerifyFilesPathList(config Config)(result []string, err error) {
	baseFolder := utils.ToLinuxPath(config.CodeFolderPath)
	filepath.Walk(baseFolder, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}
		if fi.IsDir() {
			// 忽略的目录
			if IsIgnoreDir(config, path){
				return filepath.SkipDir
			}
		}
		if !fi.IsDir() {
			// 忽略MAC隐藏文件
			if config.IgnoreMacStyleFile {
				if utils.IsMacStyleFile(fi.Name()){
					return nil
				}
			}
			// 单独忽略的文件
			if IsIgnoreFile(config,path) {
				return nil
			}

			// 忽略不需要的文件类型
			if IsIgnoreFileSuffix(config, fi.Name()){
				return nil
			}


			result = append(result, path)
		}
		return nil
	})
	return result, nil
}

func getVerifyFileListInfo(filelist []string) (verifyLineCount int64, emptyLineCount int64, singleCommentsLineCount int64, multiCommentsLineCount int64){
	for _, filePath := range filelist {
		currentVerifyLineCount, currentEmptyLineCount, currentSingleCommentsLineCount, currentMultiCommentsLineCount, err := getVerifyLinesCountForFile(filePath)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(fmt.Sprintf("文件路径:%s 有效行数:%d 空行数:%d 单行注释行数:%d,块注释行数:%d", filePath, currentVerifyLineCount, currentEmptyLineCount, currentSingleCommentsLineCount,currentMultiCommentsLineCount))
		verifyLineCount = verifyLineCount + currentVerifyLineCount
		emptyLineCount = emptyLineCount + currentEmptyLineCount
		singleCommentsLineCount = singleCommentsLineCount + currentSingleCommentsLineCount
		multiCommentsLineCount = multiCommentsLineCount + currentMultiCommentsLineCount
	}
	return
}

func getVerifyLinesCountForFile(filepath string) (int64,int64,int64,int64, error) {
	var verifyLineCount = int64(0)
	var emptyLineCount = int64(0)
	var singleCommentLineCount = int64(0)
	var multiCommentLineCount = int64(0)

	file, err := os.Open(filepath)
	if err != nil {
		return 0,0,0,0,err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	startComment := false
	for {
		byteStr, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !isPrefix {
			str := string(byteStr)

			if strings.HasPrefix(strings.Trim(str," "), "/*") && strings.HasSuffix(str, "*/") {
				singleCommentLineCount ++
				continue
			}

			if strings.HasSuffix(strings.Trim(str," "), "*/") {
				startComment = false
				multiCommentLineCount ++
				continue
			}

			// 块注释 /**/注释
			if strings.HasPrefix(str, "/*") {
				startComment = true
				multiCommentLineCount ++
				continue
			}

			if startComment {
				multiCommentLineCount ++
				continue
			}

			// 空字符不统计
			if 0 == len(byteStr) || str == "\r\n"{
				emptyLineCount ++
				continue
			}

			// 常规单行注释
			strtrim := strings.Trim(str," ")
			if len(strtrim) >= 2 && strtrim[0:2] == "//" {
				singleCommentLineCount ++
				continue
			}

			verifyLineCount ++
		}

	}
	return verifyLineCount, emptyLineCount, singleCommentLineCount,multiCommentLineCount, nil
}


// 判断该文件夹是否在被排除的范围之内
func IsIgnoreDir(config Config, dirpath string) bool {
	for _, dir := range config.IgnoreFolderPath {
		currentPath := config.CodeFolderPath + "/" + dir
		if currentPath == dirpath {
			return true
		}
	}
	return false
}

// 判断是否是要忽略的后缀
func IsIgnoreFileSuffix(config Config, filename string) bool {
	filesuffix := path.Ext(filename)
	for _, suffix := range config.ValidFileSuffix {
		if filesuffix == suffix {
			return false
		}
	}
	return true
}

// 判断是否是单独忽略的文件
func IsIgnoreFile(config Config, filePath string) bool {
	for _, ingoreFileName := range config.IgnoreFileName {
		currentPath := config.CodeFolderPath + "/" + ingoreFileName
		if currentPath == filePath {
			return true
		}
	}
	return false
}