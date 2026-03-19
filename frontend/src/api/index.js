import axios from 'axios'

// 创建 axios 实例，配置基础 URL
const request = axios.create({
  baseURL: 'http://localhost:8080/api', // 后端地址
  timeout: 5000 // 请求超时时间
})

// 封装所有 API 函数
export const api = {
  // 1. 重置系统
  resetSystem() {
    return request.post('/scheduler/reset')
  },

  // 2. 批量创建进程
  createProcesses(processes) {
    return request.post('/process/create', { processes })
  },

  // 3. 初始化调度器
  initScheduler(data) {
    return request.post('/scheduler/init', data)
  },

  // 4. 执行一步调度
  runStep() {
    return request.post('/scheduler/step')
  },

  // 5. 阻塞当前进程
  blockProcess() {
    return request.post('/process/block')
  },

  // 6. 唤醒指定进程
  wakeupProcess(name) {
    return request.post('/process/wakeup', { name })
  },

  // 7. 获取当前系统状态
  getStatus() {
    return request.get('/status')
  },

  // 8. 获取调度结果
  getResult() {
    return request.get('/result')
  }
}