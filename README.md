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