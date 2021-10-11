package main
/*
 * codelines -  A useful code line count software
 *
 * @github   https://github.com/super-l
 * @author   superl <86717375@qq.com>
 */

import (
	"encoding/json"
	"fmt"
	"github.com/super-l/codelines/core"
)

func main(){
	config, err := core.LoadConfigData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := core.GetResultInfo(config)
	resultJson, _ := json.Marshal(result)
	fmt.Println(string(resultJson))
}


