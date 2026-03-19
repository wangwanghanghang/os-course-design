package api

import (
	"fmt"
	"net/http"
	"os-course-design/entity"
	"os-course-design/manager"
	"os-course-design/model"
	"os-course-design/scheduler"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置所有 API 路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS 跨域中间件：允许前端访问
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 定义 API 路由
	api := r.Group("/api")
	{
		// 进程控制相关
		api.POST("/process/create", CreateProcessHandler)
		api.POST("/process/block", BlockProcessHandler)
		api.POST("/process/wakeup", WakeupProcessHandler)

		// 调度相关
		api.POST("/scheduler/init", InitSchedulerHandler)
		api.POST("/scheduler/step", RunStepHandler)
		api.POST("/scheduler/auto", RunAutoHandler) // 可选，先做 step
		api.POST("/scheduler/reset", ResetHandler)

		// 状态和结果
		api.GET("/status", GetStatusHandler)
		api.GET("/result", GetResultHandler)
	}

	return r
}

// CreateProcessHandler 处理创建进程请求
func CreateProcessHandler(c *gin.Context) {
	var req model.CreateProcessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Code: 400, Msg: "参数错误"})
		return
	}

	for _, p := range req.Processes {
		manager.CreateProcess(p.Name, p.InitPriority, p.ArrivalTime, p.TotalTime)
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "进程创建成功"})
}

// BlockProcessHandler 处理阻塞当前进程请求
func BlockProcessHandler(c *gin.Context) {
	manager.BlockCurrentProcess()
	scheduler.AddLog("[操作] 阻塞当前执行进程")
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "阻塞成功"})
}

// WakeupProcessHandler 处理唤醒进程请求
func WakeupProcessHandler(c *gin.Context) {
	var req model.WakeupProcessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Code: 400, Msg: "参数错误"})
		return
	}

	// 调用唤醒函数，接收是否成功的结果
	success := manager.WakeupProcess(req.Name)
	if !success {
		// 没找到进程，返回错误提示
		c.JSON(http.StatusOK, model.Response{Code: 400, Msg: "唤醒失败：等待队列中没有找到该进程，请检查进程名是否正确"})
		return
	}

	// 只有成功了，才加日志、返回成功
	scheduler.AddLog(fmt.Sprintf("[操作] 成功唤醒进程 %s", req.Name))
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "唤醒成功"})
}

// InitSchedulerHandler 处理初始化调度器请求
func InitSchedulerHandler(c *gin.Context) {
	var req model.InitSchedulerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Code: 400, Msg: "参数错误"})
		return
	}

	scheduler.InitScheduler(req.AlgorithmType, req.TimeSlice, req.PriorityStep)
	// 如果是多级反馈队列，初始化队列
	if req.AlgorithmType == "multilevel_feedback" {
		scheduler.InitMultilevelQueues()
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "调度器初始化成功"})
}

// RunStepHandler 执行一步调度
func RunStepHandler(c *gin.Context) {
	// 检查是否所有进程都完成了
	allFinished := true
	for _, p := range manager.AllProcesses {
		if p.State != entity.Finish {
			allFinished = false
			break
		}
	}
	if allFinished {
		c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "所有进程已执行完成", Data: "finished"})
		return
	}

	// 根据当前算法执行一步
	switch scheduler.CurrentAlgorithm {
	case "priority_rr":
		scheduler.RunPriorityRRStep()
	case "multilevel_feedback":
		scheduler.RunMultilevelFeedbackStep()
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "执行一步成功"})
}

// RunAutoHandler 自动运行调度器直到所有进程完成
func RunAutoHandler(c *gin.Context) {
	// 检查是否所有进程都完成了
	allFinished := true
	for _, p := range manager.AllProcesses {
		if p.State != entity.Finish {
			allFinished = false
			break
		}
	}
	if allFinished {
		c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "所有进程已执行完成", Data: "finished"})
		return
	}

	// 自动运行调度器直到所有进程完成
	for !allFinished {
		// 根据当前算法执行一步
		switch scheduler.CurrentAlgorithm {
		case "priority_rr":
			scheduler.RunPriorityRRStep()
		case "multilevel_feedback":
			scheduler.RunMultilevelFeedbackStep()
		}

		// 检查是否所有进程都完成了
		allFinished = true
		for _, p := range manager.AllProcesses {
			if p.State != entity.Finish {
				allFinished = false
				break
			}
		}
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "自动调度完成"})
}

// ResetHandler 重置系统
func ResetHandler(c *gin.Context) {
	manager.InitSystem()
	scheduler.ClearLogs()
	scheduler.InitMultilevelQueues()
	c.JSON(http.StatusOK, model.Response{Code: 200, Msg: "系统重置成功"})
}

// GetStatusHandler 获取当前系统状态
func GetStatusHandler(c *gin.Context) {
	data := model.StatusData{
		AllProcesses:   manager.AllProcesses,
		ReadyQueue:     manager.ReadyQueue,
		WaitQueue:      manager.WaitQueue,
		CurrentProcess: manager.CurrentProcess,
		SystemTime:     manager.SystemTime,
		Logs:           scheduler.Logs,
	}
	// 多级反馈队列特殊处理：把多级队列也放进去（可选，这里简化）
	c.JSON(http.StatusOK, model.Response{Code: 200, Data: data})
}

// GetResultHandler 获取调度结果
func GetResultHandler(c *gin.Context) {
	totalTurnaround := 0
	count := 0
	for _, p := range manager.AllProcesses {
		if p.State == entity.Finish {
			totalTurnaround += p.TurnaroundTime
			count++
		}
	}

	avgTurnaround := 0.0
	if count > 0 {
		avgTurnaround = float64(totalTurnaround) / float64(count)
	}

	data := model.ResultData{
		Processes:         manager.AllProcesses,
		AvgTurnaroundTime: avgTurnaround,
	}

	c.JSON(http.StatusOK, model.Response{Code: 200, Data: data})
}
