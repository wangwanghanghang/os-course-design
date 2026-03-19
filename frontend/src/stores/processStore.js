import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '../api/index'

// 定义 Store
export const useProcessStore = defineStore('process', () => {
  // 1. 定义响应式状态（对应后端返回的数据）
  const allProcesses = ref([])       // 所有进程列表
  const readyQueue = ref([])         // 就绪队列
  const waitQueue = ref([])          // 等待队列
  const currentProcess = ref(null)   // 当前执行进程
  const systemTime = ref(0)          // 系统时间
  const logs = ref([])                // 执行日志
  const result = ref(null)            // 调度结果

  // 2. 定义 Actions（调用 API 并更新状态的方法）
  
  // 重置系统
  const resetSystem = async () => {
    await api.resetSystem()
    await fetchStatus() // 重置后立即获取最新状态
  }

  // 创建进程
  const createProcesses = async (processes) => {
    await api.createProcesses(processes)
    await fetchStatus()
  }

  // 初始化调度器
  const initScheduler = async (data) => {
    await api.initScheduler(data)
    await fetchStatus()
  }

  // 执行一步调度
  const runStep = async () => {
    const res = await api.runStep()
    await fetchStatus()
    return res.data.data // 返回是否 finished
  }

  // 阻塞进程
  const blockProcess = async () => {
    await api.blockProcess()
    await fetchStatus()
  }

// 唤醒进程
const wakeupProcess = async (name) => {
  const res = await api.wakeupProcess(name)
  // 关键：判断后端返回的业务状态码，不是200就抛出错误
  if (res.data.code !== 200) {
    throw new Error(res.data.msg || '唤醒失败')
  }
  // 只有成功了，才刷新状态
  await fetchStatus()
}

// 核心方法：从后端获取最新状态并更新
const fetchStatus = async () => {
  const res = await api.getStatus()
  const data = res.data.data
  // 修复：加空值合并运算符 ??，如果后端返回 null，就赋值为空数组 []
  allProcesses.value = data.allProcesses ?? []
  readyQueue.value = data.readyQueue ?? []
  waitQueue.value = data.waitQueue ?? []
  currentProcess.value = data.currentProcess ?? null
  systemTime.value = data.systemTime ?? 0
  logs.value = data.logs ?? []
}

  // 获取调度结果
  const fetchResult = async () => {
    const res = await api.getResult()
    result.value = res.data.data
  }

  // 3. 返回状态和方法，供组件使用
  return {
    allProcesses,
    readyQueue,
    waitQueue,
    currentProcess,
    systemTime,
    logs,
    result,
    resetSystem,
    createProcesses,
    initScheduler,
    runStep,
    blockProcess,
    wakeupProcess,
    fetchStatus,
    fetchResult
  }
})