package n81i

// 步骤
// 读取语言包
// 解析项目的语言包文件
// 生成 EXCEL
type N81i struct {
	sourceDir string
}

type DirReader interface {
	DirRead()
}

type FileWriter interface {
	FileWrite()
}
