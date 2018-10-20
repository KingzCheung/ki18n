package typer

// 解析驱动接口

type Typer interface {
	// 解析数据
	// 获取对应列表数数据，反正 MAP
	Parse(col int) map[string]string
}
