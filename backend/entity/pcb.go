package entity

// PCB 进程控制块结构体，包含任务书要求的所有字段
type PCB struct {
	Name           string       `json:"name"`           // 进程名，如 P1、P2
	InitPriority   int          `json:"initPriority"`   // 初始优先级（1-100，越小越高）
	CurrPriority   int          `json:"currPriority"`   // 当前优先级（动态变化）
	ArrivalTime    int          `json:"arrivalTime"`    // 到达时间
	TotalTime      int          `json:"totalTime"`      // 总需要运行时间
	UsedTime       int          `json:"usedTime"`       // 已用 CPU 时间
	RemainTime     int          `json:"remainTime"`     // 剩余运行时间
	State          ProcessState `json:"state"`          // 当前进程状态
	StartTime      int          `json:"startTime"`      // 首次开始执行时间（初始为 -1，表示未开始）
	FinishTime     int          `json:"finishTime"`     // 结束执行时间
	TurnaroundTime int          `json:"turnaroundTime"` // 周转时间（FinishTime - ArrivalTime）
	TimeSliceUsed  int          `json:"timeSliceUsed"`  // 【新增】本次时间片已使用的时间
}

// NewPCB 创建一个新的 PCB 对象，初始化默认值
func NewPCB(name string, initPriority, arrivalTime, totalTime int) *PCB {
	return &PCB{
		Name:          name,
		InitPriority:  initPriority,
		CurrPriority:  initPriority, // 初始当前优先级 = 初始优先级
		ArrivalTime:   arrivalTime,
		TotalTime:     totalTime,
		UsedTime:      0,
		RemainTime:    totalTime, // 初始剩余时间 = 总时间
		State:         Ready,     // 任务书要求：初始状态均为 R（就绪）
		StartTime:     -1,        // -1 表示还没开始执行
		TimeSliceUsed: 0,         // 新增字段初始化
	}
}
