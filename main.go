// ki18n 一个语言包转换工具
// 用例

// 命令行设计
// 主要实现两个功能：
// 1,实现excel或者csv到对应目录文本工作
// 2,对应用的文本到excel 或者csv的转换工作

//命令设计
// gi18n build filename.xlsx [--lang=zh-cn,zh-hk,en-us] [--format=json] # build file to json/php
// gi18n build filename.xlsx [--lang=zh-cn] [--cell] # build file to json/php
// gi18n reverse filedir 反向通过项目的语言包生成运营人员可编辑的excel文件

package main

import (
	"github.com/kingzcheung/ki18n/cmd"
)

func main() {
	cmd.Execute()
}
