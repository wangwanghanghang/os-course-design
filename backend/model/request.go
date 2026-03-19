package model

// CreateProcessRequest 前端请求：创建进程
type CreateProcessRequest struct {
	Processes []ProcessInfo `json:"processes"` // 批量创建多个进程
}

// ProcessInfo 单个进程的信息
type ProcessInfo struct {
	Name         string `json:"name"`
	InitPriority int    `json:"initPriority"`
	ArrivalTime  int    `json:"arrivalTime"`
	TotalTime    int    `json:"totalTime"`
}

// InitSchedulerRequest 前端请求：初始化调度算法
type InitSchedulerRequest struct {
	AlgorithmType string `json:"algorithmType"` // "priority_rr" 或 "multilevel_feedback"
	TimeSlice     int    `json:"timeSlice"`     // 时间片大小
	PriorityStep  int    `json:"priorityStep"`  // 优先级下降步长（仅优先级算法用）
}

// WakeupProcessRequest 前端请求：唤醒进程
type WakeupProcessRequest struct {
	Name string `json:"name"` // 要唤醒的进程名
}
