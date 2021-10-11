package core
/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

import (
	"encoding/json"
	"os"
)

func LoadConfigData() (Config, error) {
	filePtr, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer filePtr.Close()
	var config Config
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
