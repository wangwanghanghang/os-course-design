package scheduler

import (
	"fmt"
	"os-course-design/entity"
	"os-course-design/manager"
)

// 多级反馈队列的全局变量
var (
	MultilevelQueues [3][]*entity.PCB // 3 个队列，编号 0/1/2（0 是第 1 级，优先级最高）
)

// InitMultilevelQueues 初始化多级反馈队列
func InitMultilevelQueues() {
	MultilevelQueues = [3][]*entity.PCB{}
}

// 辅助函数：根据队列级别（0/1/2）返回对应队列的时间片大小
func getQueueTimeSlice(level int) int {
	switch level {
	case 0:
		return TimeSlice // 第1级：基础时间片
	case 1:
		return TimeSlice * 2 // 第2级：基础×2
	case 2:
		return TimeSlice * 4 // 第3级：基础×4
	default:
		return TimeSlice
	}
}

// RunMultilevelFeedbackStep 执行一步多级反馈队列调度（单时间单位推进）
func RunMultilevelFeedbackStep() {
	// 1. 检查新到达的进程，加入第 1 级队列（Queue 0）
	for _, pcb := range manager.AllProcesses {
		if pcb.ArrivalTime <= manager.SystemTime &&
			pcb.State == entity.Ready &&
			pcb.StartTime == -1 &&
			!isInMultilevelQueue(pcb) {
			MultilevelQueues[0] = append(MultilevelQueues[0], pcb)
			pcb.State = entity.Ready
			AddLog(fmt.Sprintf("[时间 %d] 进程 %s 到达，加入第 1 级队列", manager.SystemTime, pcb.Name))
		}
	}

	// 2. 如果当前没有执行进程，从高到低找非空队列，选进程执行
	if manager.CurrentProcess == nil {
		found := false
		var currentQueueLevel int
		for i := 0; i < 3; i++ {
			if len(MultilevelQueues[i]) > 0 {
				// 找到非空队列，取第一个进程
				manager.CurrentProcess = MultilevelQueues[i][0]
				MultilevelQueues[i] = MultilevelQueues[i][1:]
				currentQueueLevel = i
				// 修改状态为执行态
				manager.CurrentProcess.State = entity.Execute
				// 重置本次时间片已用时长
				manager.CurrentProcess.TimeSliceUsed = 0
				// 记录首次开始时间
				if manager.CurrentProcess.StartTime == -1 {
					manager.CurrentProcess.StartTime = manager.SystemTime
				}
				AddLog(fmt.Sprintf("[时间 %d] 进程 %s 从第 %d 级队列开始执行（队列时间片：%d）",
					manager.SystemTime, manager.CurrentProcess.Name, currentQueueLevel+1, getQueueTimeSlice(currentQueueLevel)))
				found = true
				break
			}
		}
		if !found {
			AddLog(fmt.Sprintf("[时间 %d] 所有队列为空，系统时间+1", manager.SystemTime))
			manager.SystemTime++
			return
		}
		return // 选完进程就返回，下一次step再执行
	}

	// 3. 有正在执行的进程，先确定它所属的队列级别和对应时间片
	currentProc := manager.CurrentProcess
	// 计算当前进程应该用的时间片（根据它的已用时间，判断它属于哪个队列级别）
	queueLevel := 0
	if currentProc.UsedTime >= TimeSlice {
		queueLevel = 1
	}
	if currentProc.UsedTime >= TimeSlice+TimeSlice*2 {
		queueLevel = 2
	}
	queueTimeSlice := getQueueTimeSlice(queueLevel)

	// 4. 执行1个时间单位
	currentProc.UsedTime += 1
	currentProc.RemainTime -= 1
	currentProc.TimeSliceUsed += 1
	manager.SystemTime += 1

	// 5. 打印执行日志
	AddLog(fmt.Sprintf("[时间 %d] 进程 %s 执行1个时间单位，剩余运行时间：%d，本次时间片已用：%d/%d",
		manager.SystemTime, currentProc.Name, currentProc.RemainTime, currentProc.TimeSliceUsed, queueTimeSlice))

	// 6. 判断进程是否执行完成
	if currentProc.RemainTime <= 0 {
		AddLog(fmt.Sprintf("[时间 %d] 进程 %s 执行完成", manager.SystemTime, currentProc.Name))
		manager.FinishCurrentProcess()
		return
	}

	// 7. 判断当前队列的时间片是否用完
	if currentProc.TimeSliceUsed >= queueTimeSlice {
		// 时间片用完，降级到下一级队列
		nextLevel := queueLevel + 1
		if nextLevel > 2 {
			nextLevel = 2 // 最高级是第3级（索引2）
		}
		MultilevelQueues[nextLevel] = append(MultilevelQueues[nextLevel], currentProc)
		currentProc.State = entity.Ready
		AddLog(fmt.Sprintf("[时间 %d] 进程 %s 时间片用完，降级到第 %d 级队列", manager.SystemTime, currentProc.Name, nextLevel+1))
		// 清空当前执行进程，下一次调度选新的
		manager.CurrentProcess = nil
		return
	}

	// 8. 时间片没用完、进程也没完成，继续保持执行态，下一次step继续执行
}

// 辅助函数：判断进程是否在多级队列中
func isInMultilevelQueue(pcb *entity.PCB) bool {
	for i := 0; i < 3; i++ {
		for _, p := range MultilevelQueues[i] {
			if p.Name == pcb.Name {
				return true
			}
		}
	}
	return false
}
