# ki18n EXCEL 语言包生成语言包工具

### 安装
 - [二进制下载](https://github.com/KingzCheung/ki18n/releases)
 - 源码编译
 ```shell
    下载并安装依赖
    go get -u github.com/KingzCheung/ki18n.git
    cd cd $GOPATH/src/github.com/KingzCheung/ki18n
    go build
 ```

交叉编译 `Windows`,`Linux`,`Mac`:

`Mac`下编译 Linux 和 Windows 64位可执行程序:

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

`Linux`下编译 Mac 和 Windows 64位可执行程序:

```shell
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

`Windows` 下编译 Mac 和 Linux 64位可执行程序:

```shell
#Mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build

#Linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build
```

### 用法

> 0.03 以下参考： [点击这里](https://github.com/KingzCheung/ki18n/tree/0.03)

通过以下命令查看有哪些命令

```shell
ki18n 是一个支持EXCEL，CSV 格式的语言文件快速转成JSON,PHP,IOS等平台的语言包工具

Usage:
  ki18n [flags]
  ki18n [command]

Available Commands:
  help        Help about any command
  json        生成 json 格式的语言包
  php         生成 php 格式的语言包
  strings     生成 strings 格式的语言包

Flags:
  -f, --file string        需要转换的源文件 (default "language.xlsx")
  -h, --help               help for ki18n
  -l, --language strings   转换目标的语言的列表

Use "ki18n [command] --help" for more information about a command.

```

比如生成 json格式的语言包:

```shell
ki18n json // 默认会找当前目录下的 language.xlsx 文件
ki18n json --file=your/dir/file.xlsx // 也可以通过 --file 来指定文件

// 通过--language 来指定需要输出的语言包
ki18n json -f=your/dir/file.xlsx --language=zh-CN,zh-HK,en-US
```
输出PHP格式
```shell
ki18n php
```

Excel 格式:

| key   | 简体中文  | 繁体中文  | 英文          | ...  |
| ----- | ----- | ----- | ----------- | ---- |
| hello | 你好,世界 | 你好,世界 | hello wolrd | ...  |
| ...   | ...   | ...   | ...         | ...  |

默认三种语言包输出 `zh-CN`,`zh-HK`,`en-US`

可以通过 i18n.yaml 配置文件扩展:

```yaml
language:
  - zh-CN
  - zh-HK
  - en-US
```
