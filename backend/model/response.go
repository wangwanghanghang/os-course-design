package model

import "os-course-design/entity"

// Response 统一响应格式：所有 API 都返回这个格式，方便前端处理
type Response struct {
	Code int         `json:"code"` // 200 表示成功
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 实际数据
}

// StatusData 前端请求 /api/status 时返回的数据
type StatusData struct {
	AllProcesses   []*entity.PCB `json:"allProcesses"`
	ReadyQueue     []*entity.PCB `json:"readyQueue"`
	WaitQueue      []*entity.PCB `json:"waitQueue"`
	CurrentProcess *entity.PCB   `json:"currentProcess"`
	SystemTime     int           `json:"systemTime"`
	Logs           []string      `json:"logs"` // 执行日志
}

// ResultData 前端请求 /api/result 时返回的数据
type ResultData struct {
	Processes         []*entity.PCB `json:"processes"`
	AvgTurnaroundTime float64       `json:"avgTurnaroundTime"` // 平均周转时间
}
