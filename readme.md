### Codelines简介

![codelines](https://github.com/super-l/codelines/blob/main/run.png)

一款跨平台，可以统计项目代码行数的软件(命令行软件，无界面)，支持多种自定义过滤。
主要用于代码安全审计服务相关的费用评估。

### 功能特色

1. 自定义项目文件夹路径(支持绝对路径和相对路径);
2. 自定义有效文件后缀(比如只统计php文件)；
3. 自定义忽略的文件夹路径(比如要忽略js,css等静态文件夹下的所有文件，或者cache缓存目录文件)；
4. 自定义忽略特定的文件(比如phpinfo.php,test.php等文件)；
5. 自定义是否忽略空行；
6. 自定义是否忽略单行注释(包括'//'单行注释 以及 '/*单行代码*/' 注释)；
7. 自定义是否忽略多行注释(包括'/* 之间的多行代码 */')；


### 运行结果
```
superldeMacBook-Pro:codelines superl$ ./codelines
文件路径:example/Email.php 有效行数:1947 空行数:391 单行注释行数:0,块注释行数:152
文件路径:example/Ftp.php 有效行数:492 空行数:103 单行注释行数:0,块注释行数:72
文件路径:example/Medoo.php 有效行数:892 空行数:203 单行注释行数:12,块注释行数:8
文件路径:example/Zip.php 有效行数:397 空行数:62 单行注释行数:0,块注释行数:73
文件路径:example/index.php 有效行数:5 空行数:3 单行注释行数:2,块注释行数:11

{"AllFileCount":6,"ValidFileCount":5,"VaildLineCount":3733,"EmptyLineCount":762,"SingleCommentsLineCount":14,"MulitCommentsLineCount":316}
```

最后返回的JSON数据，是总统计数据！说明如下：
```
	AllFileCount              // 总文件数
	ValidFileCount            // 有效文件数
	VaildLineCount            // 有效文件行数
	EmptyLineCount            // 空格数
	SingleCommentsLineCount   // 注释行数
	MulitCommentsLineCount    // 注释行数
```

### 使用说明
```
使用前，修改config.json，对参数进行配置，比如项目文件夹路径。
配置完成后，直接运行程序即可。

linux下运行：./codelines
windows下运行: codelines
```

### 编译说明
linux系统；
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/codelines
```

windows系统：
```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/windows_amd64/codelines.exe
```

### 其他信息

如果好用，点赞走起。后续版本，考虑加入导出html报表功能。