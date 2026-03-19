package scheduler

import (
	"fmt"
)

// 全局调度配置
var (
	CurrentAlgorithm string   // 当前使用的算法
	TimeSlice        int      // 时间片大小
	PriorityStep     int      // 优先级下降步长
	Logs             []string // 执行日志
)

// InitScheduler 初始化调度器配置
func InitScheduler(algorithmType string, timeSlice, priorityStep int) {
	CurrentAlgorithm = algorithmType
	TimeSlice = timeSlice
	PriorityStep = priorityStep
	Logs = []string{}
	// 修复日志乱码，用 fmt.Sprintf 格式化数字
	AddLog(fmt.Sprintf("调度器初始化完成：算法=%s，基础时间片=%d，优先级下降步长=%d", algorithmType, timeSlice, priorityStep))
}

// AddLog 添加一条执行日志
func AddLog(log string) {
	Logs = append(Logs, log)
}

// ClearLogs 清空日志
func ClearLogs() {
	Logs = []string{}
}
