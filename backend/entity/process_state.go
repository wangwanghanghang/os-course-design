package entity

// ProcessState 定义进程状态类型
type ProcessState int

// 使用 iota 定义 4 种进程状态常量
const (
	Execute ProcessState = iota // 执行态 E
	Ready                       // 就绪态 R
	Wait                        // 等待态 W
	Finish                      // 完成态 F
)

// String 方法：让状态能打印成可读的字符串（方便调试和日志）
func (s ProcessState) String() string {
	switch s {
	case Execute:
		return "E"
	case Ready:
		return "R"
	case Wait:
		return "W"
	case Finish:
		return "F"
	default:
		return "Unknown"
	}
}
