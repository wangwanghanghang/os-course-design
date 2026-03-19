<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>🎮 调度执行控制</span>
        <span style="font-size: 14px; color: #666;">当前系统时间：{{ store.systemTime }}</span>
      </div>
    </template>

    <div class="control-group">
      <el-button type="primary" @click="handleStep" :disabled="isAutoRunning">单步执行</el-button>
      <el-button type="success" @click="handleAutoStart" :disabled="isAutoRunning">自动执行</el-button>
      <el-button type="warning" @click="handleAutoStop" :disabled="!isAutoRunning">暂停</el-button>
      <el-divider direction="vertical" />
      <el-button type="danger" @click="handleBlock">阻塞当前进程</el-button>
      <el-input v-model="wakeupName" placeholder="输入要唤醒的进程名" style="width: 150px;" />
      <el-button @click="handleWakeup">唤醒进程</el-button>
    </div>
  </el-card>
</template>

<script setup>
import { ref } from 'vue'
import { useProcessStore } from '../stores/processStore'
import { ElMessage } from 'element-plus'

const store = useProcessStore()
const isAutoRunning = ref(false)
const autoTimer = ref(null)
const wakeupName = ref('')

// 单步执行
const handleStep = async () => {
  try {
    const status = await store.runStep()
    if (status === 'finished') {
      ElMessage.success('所有进程执行完成！')
      await store.fetchResult()
    }
  } catch (e) {
    ElMessage.error('执行失败：' + e.message)
  }
}

// 自动执行
const handleAutoStart = () => {
  isAutoRunning.value = true
  autoTimer.value = setInterval(async () => {
    const status = await store.runStep()
    if (status === 'finished') {
      handleAutoStop()
      ElMessage.success('所有进程执行完成！')
      await store.fetchResult()
    }
  }, 800) // 每800ms执行一步，可调整
}

// 暂停
const handleAutoStop = () => {
  isAutoRunning.value = false
  if (autoTimer.value) {
    clearInterval(autoTimer.value)
    autoTimer.value = null
  }
}

// 阻塞
const handleBlock = async () => {
  try {
    await store.blockProcess()
    ElMessage.success('阻塞成功！')
  } catch (e) {
    ElMessage.error('阻塞失败：' + e.message)
  }
}

// 唤醒进程
const handleWakeup = async () => {
  if (!wakeupName.value.trim()) {
    ElMessage.warning('请输入要唤醒的进程名！（如 P1、P2，注意大小写）')
    return
  }
  try {
    // 调用store里封装好的方法
    await store.wakeupProcess(wakeupName.value.trim())
    ElMessage.success('唤醒成功！')
    wakeupName.value = '' // 成功后清空输入框
  } catch (e) {
    // 捕获后端返回的错误，弹出红色提示
    ElMessage.error(e.message)
  }
}
</script>

<style scoped>
.panel-card {
  margin-bottom: 20px;
}
.control-group {
  display: flex;
  align-items: center;
  gap: 10px;
  flex-wrap: wrap;
}
</style>