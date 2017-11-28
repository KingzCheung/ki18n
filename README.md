# ki18n XLSX -> JSON 生成语言包工具

### 安装

下载源码丢到`gopath`下,运行以下命令:

```shell
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

Excel 格式:

| key   | 简体中文  | 繁体中文  | 英文          | ...  |
| ----- | ----- | ----- | ----------- | ---- |
| hello | 你好,世界 | 你好,世界 | hello wolrd | ...  |
| ...   | ...   | ...   | ...         | ...  |

默认三种语言包输出 `zh-CN`,`zh-HK`,`en-US`

可以通过 i18n.ini 配置文件扩展:

```ini
[default]
language = zh-CN,zh-HK,en-US
```

Excel 文件取命令行工作目录下的`language.xlsx`  可以通过 `-f`指定:

```shell
ki18n -f yourxlsxname.xlsx
```

默认会生成一个`lang`目录,里面对应生成`i18n.ini`配置的JSON

合并语言包为一个JSON使用 `-m`选项:

```shell
ki18n -m -f yourxlsxname.xlsx
```

默认会生成一个叫`locales.json`的文件

