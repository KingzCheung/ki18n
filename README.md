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

由于更换了excel的解析库，新的EXCEL必须要指定工作表，因此需要把工作表默认设置为 `Sheet1`。

表格内容如下，第一列为key值，第二表开始为对应的语言翻译。

| key      | zh-cn              | zh-hk              | en-us                         |
| -------- | ------------------ | ------------------ | ----------------------------- |
| hello    | 你好，世界         | 你好，世界         | Hello World                   |
| wellcome | 欢迎来到计算机世界 | 歡迎來到計算機世界 | Welcome to the computer world |

命令示例:

```bash
$ ./ki18n --help
目标是解析一定格式的语言文件，来生成对应项目中可直接使用的语言包文件

Usage:
  ki18n [flags]
  ki18n [command]

Available Commands:
  Print       
  build       构建项目可使用的i18n文件
  help        Help about any command

Flags:
  -h, --help      help for ki18n
  -v, --version   Print version information and quit

Use "ki18n [command] --help" for more information about a command.

```

生成语言包：

```bash
ki18n build language.csv #默认会生成json格式的语言包文件
ki18n build language.xlsx --format=php #指定生成php项目可用的语言包
ki18n build --lang=zh-cn,en-us,zh-hk --format=php language.csv #指定生成语言包名称，与表格的列对应，如果不指定 --lang 参数，生成的文件名会取表头的值
```