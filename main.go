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
	"github.com/gookit/color"
	"github.com/super-l/codelines/core"
)

func showBanner() {
	banner := `
codelines v1.0
Created by superl[忘忧草安全]             QQ: 86717375

               _      _ _ _                 
              | |    ( ) (_)                
  ___ ___   __| | ___|/| |_ _ __   ___  ___ 
 / __/ _ \ / _' |/ _ \ | | | '_ \ / _ \/ __|
| (_| (_) | (_| |  __/ | | | | | |  __/\__ \
\___\___/ \__,_|\___| |_|_|_| |_|\___||___/

`
	color.Green.Println(banner)
}

func main(){
	showBanner()
	config, err := core.LoadConfigData()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result := core.GetResultInfo(config)
	resultJson, _ := json.Marshal(result)
	fmt.Println(string(resultJson))
}


