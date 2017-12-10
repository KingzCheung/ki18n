# ki18n XLSX => (JSON,PHP,iOS) 生成语言包工具

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
[DEFAULT]
language = zh-CN,zh-HK,en-US
```

#### JSON

Excel 文件取命令行工作目录下的`language.xlsx`  可以通过 `-f`指定:

```shell
ki18n -f yourxlsxname.xlsx
```

默认会生成一个`lang`目录,里面对应生成`i18n.ini`配置的JSON



有一些前端项目使用的语言包可能需要合并为一个JSON 包，此时可以向`i18n.ini` 添加以下选项

```ini
merge = true
```

默认会生成一个叫`locales.json`的文件

> 当此选项开启时， -t 选项会失效，并强行输出JSON包

#### PHP

生成 php格式的语言包:

```shell
ki18n -t=php -f=yourxlsxname.xlsx
```

默认生成php格式的数组,支持 `thinkphp`,`laravel`

####  iOS

生成 iOS格式语言包:

```shell
ki18n -t=strings -f=yourxlsxname.xlsx
```



### 添加CSV 格式

`csv`格式文件默认分割符为 `;` 如果发现有冲突可以在`i18n.ini` 配置

```ini
splitter = ,
```

