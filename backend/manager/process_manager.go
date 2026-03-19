package manager

import (
	"fmt"
	"os-course-design/entity"
	"sort"
)

// 全局变量：维护系统中的所有进程和队列
var (
	AllProcesses   []*entity.PCB // 所有进程列表
	ReadyQueue     []*entity.PCB // 就绪队列
	WaitQueue      []*entity.PCB // 等待队列
	CurrentProcess *entity.PCB   // 当前正在执行的进程
	SystemTime     int           // 当前系统时间（模拟时间片）
)

// InitSystem 初始化系统：清空所有队列和进程，重置系统时间
func InitSystem() {
	AllProcesses = []*entity.PCB{}
	ReadyQueue = []*entity.PCB{}
	WaitQueue = []*entity.PCB{}
	CurrentProcess = nil
	SystemTime = 0
	fmt.Println("系统已重置，所有数据已清空")
}

// CreateProcess 创建一个新进程并加入系统
func CreateProcess(name string, initPriority, arrivalTime, totalTime int) *entity.PCB {
	pcb := entity.NewPCB(name, initPriority, arrivalTime, totalTime)
	// 新增：打印创建的进程信息
	fmt.Printf("✅ 成功创建进程：%s，到达时间：%d，总运行时间：%d\n", pcb.Name, pcb.ArrivalTime, pcb.TotalTime)

	AllProcesses = append(AllProcesses, pcb)
	// 如果进程到达时间 <= 当前系统时间，直接加入就绪队列
	if arrivalTime <= SystemTime {
		AddToReadyQueue(pcb)
		fmt.Printf("📥 进程 %s 已加入就绪队列，就绪队列长度：%d\n", name, len(ReadyQueue))
	}
	return pcb
}

// AddToReadyQueue 将进程加入就绪队列，并按优先级排序（优先级高的在前）
func AddToReadyQueue(pcb *entity.PCB) {
	pcb.State = entity.Ready
	pcb.TimeSliceUsed = 0 // 重置时间片计数
	ReadyQueue = append(ReadyQueue, pcb)
	// 排序：按 CurrPriority 从小到大排（数值越小优先级越高）
	sort.Slice(ReadyQueue, func(i, j int) bool {
		return ReadyQueue[i].CurrPriority < ReadyQueue[j].CurrPriority
	})
}

// BlockCurrentProcess 阻塞当前正在执行的进程
func BlockCurrentProcess() {
	if CurrentProcess == nil {
		fmt.Println("❌ 没有正在执行的进程，无法阻塞")
		return
	}
	// 状态改为等待，加入等待队列
	CurrentProcess.State = entity.Wait
	WaitQueue = append(WaitQueue, CurrentProcess)
	fmt.Printf("⏸️ 进程 %s 已被阻塞，加入等待队列\n", CurrentProcess.Name)
	// 清空当前执行进程，触发调度
	CurrentProcess = nil
}

// WakeupProcess 唤醒指定名称的等待进程，返回是否唤醒成功
func WakeupProcess(name string) bool {
	// 从等待队列中找到该进程
	var targetPCB *entity.PCB
	targetIndex := -1
	for i, pcb := range WaitQueue {
		if pcb.Name == name {
			targetPCB = pcb
			targetIndex = i
			break
		}
	}
	if targetPCB == nil {
		fmt.Printf("❌ 等待队列中没有找到进程 %s，唤醒失败\n", name)
		return false // 没找到进程，返回失败
	}

	// 从等待队列中移除
	WaitQueue = append(WaitQueue[:targetIndex], WaitQueue[targetIndex+1:]...)
	// 加入就绪队列
	AddToReadyQueue(targetPCB)
	fmt.Printf("✅ 进程 %s 已被唤醒，加入就绪队列\n", name)

	// 任务书要求：如果被唤醒进程优先级 > 当前执行进程，立即抢占 CPU
	if CurrentProcess != nil && targetPCB.CurrPriority < CurrentProcess.CurrPriority {
		// 抢占：把当前进程放回就绪队列，清空当前进程，触发调度
		fmt.Printf("⚡ 进程 %s 优先级更高，抢占当前执行的进程 %s\n", targetPCB.Name, CurrentProcess.Name)
		AddToReadyQueue(CurrentProcess)
		CurrentProcess = nil
	}
	return true // 唤醒成功
}

// FinishCurrentProcess 结束当前正在执行的进程
func FinishCurrentProcess() {
	if CurrentProcess == nil {
		return
	}
	// 状态改为完成，记录结束时间和周转时间
	CurrentProcess.State = entity.Finish
	CurrentProcess.FinishTime = SystemTime
	CurrentProcess.TurnaroundTime = CurrentProcess.FinishTime - CurrentProcess.ArrivalTime
	fmt.Printf("🏁 进程 %s 执行完成，周转时间：%d\n", CurrentProcess.Name, CurrentProcess.TurnaroundTime)
	// 清空当前执行进程，触发调度
	CurrentProcess = nil
}
