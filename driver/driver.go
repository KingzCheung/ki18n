package driver

// 解析驱动接口

type Driver interface {
	// 解析数据
	Parse(col int) map[string]string
}
