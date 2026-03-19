package scheduler

import (
	"fmt"
	"os-course-design/entity"
	"os-course-design/manager"
)

// RunPriorityRRStep 执行一步调度（单时间单位推进）
func RunPriorityRRStep() {
	// 1. 先检查是否有新到达的进程，加入就绪队列
	for _, pcb := range manager.AllProcesses {
		// 条件：到达时间<=当前系统时间、就绪态、还没开始执行、不在就绪队列里
		if pcb.ArrivalTime <= manager.SystemTime &&
			pcb.State == entity.Ready &&
			pcb.StartTime == -1 &&
			!isInQueue(pcb, manager.ReadyQueue) {
			manager.AddToReadyQueue(pcb)
			AddLog(fmt.Sprintf("[时间 %d] 进程 %s 到达，加入就绪队列", manager.SystemTime, pcb.Name))
		}
	}

	// 2. 如果当前没有正在执行的进程，从就绪队列选优先级最高的
	if manager.CurrentProcess == nil {
		// 就绪队列为空，系统时间+1，跳过本次调度
		if len(manager.ReadyQueue) == 0 {
			AddLog(fmt.Sprintf("[时间 %d] 就绪队列为空，系统时间+1", manager.SystemTime))
			manager.SystemTime++
			return
		}
		// 取就绪队列第一个（已经按优先级排好序了）
		manager.CurrentProcess = manager.ReadyQueue[0]
		// 从就绪队列移除
		manager.ReadyQueue = manager.ReadyQueue[1:]
		// 修改状态为执行态
		manager.CurrentProcess.State = entity.Execute
		// 重置本次时间片的已用时长
		manager.CurrentProcess.TimeSliceUsed = 0
		// 如果是第一次执行，记录开始时间
		if manager.CurrentProcess.StartTime == -1 {
			manager.CurrentProcess.StartTime = manager.SystemTime
		}
		AddLog(fmt.Sprintf("[时间 %d] 进程 %s 开始执行（当前优先级：%d）", manager.SystemTime, manager.CurrentProcess.Name, manager.CurrentProcess.CurrPriority))
		return // 选完进程就返回，本次只更新状态，下一次step再执行
	}

	// 3. 有正在执行的进程，执行1个时间单位
	currentProc := manager.CurrentProcess
	// 3.1 更新进程执行时间
	currentProc.UsedTime += 1
	currentProc.RemainTime -= 1
	currentProc.TimeSliceUsed += 1
	manager.SystemTime += 1

	// 3.2 打印本次执行日志
	AddLog(fmt.Sprintf("[时间 %d] 进程 %s 执行1个时间单位，剩余运行时间：%d，本次时间片已用：%d/%d",
		manager.SystemTime, currentProc.Name, currentProc.RemainTime, currentProc.TimeSliceUsed, TimeSlice))

	// 4. 判断进程是否执行完成
	if currentProc.RemainTime <= 0 {
		AddLog(fmt.Sprintf("[时间 %d] 进程 %s 执行完成", manager.SystemTime, currentProc.Name))
		manager.FinishCurrentProcess()
		return
	}

	// 5. 判断时间片是否用完
	if currentProc.TimeSliceUsed >= TimeSlice {
		// 时间片用完：优先级下降，放回就绪队列
		currentProc.CurrPriority += PriorityStep
		AddLog(fmt.Sprintf("[时间 %d] 进程 %s 时间片用完，优先级下降为 %d，放回就绪队列", manager.SystemTime, currentProc.Name, currentProc.CurrPriority))
		manager.AddToReadyQueue(currentProc)
		// 清空当前执行进程，下一次调度选新的进程
		manager.CurrentProcess = nil
		return
	}

	// 6. 时间片没用完、进程也没完成，继续保持执行态，下一次step继续执行
}

// 辅助函数：判断进程是否已经在就绪队列里
func isInQueue(pcb *entity.PCB, queue []*entity.PCB) bool {
	for _, p := range queue {
		if p.Name == pcb.Name {
			return true
		}
	}
	return false
}
