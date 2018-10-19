package main

import "github.com/KingzCheung/ki18n/cmd"

// ki18n 一个语言包转换工具
// 用例
// ki18n json -f=language.xlsx -m
// ki18n php  -f=language.xlsx -m
// ki18n strings  -f=language.xlsx -m

func main() {
	cmd.Execute()
}
