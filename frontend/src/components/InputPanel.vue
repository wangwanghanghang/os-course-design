<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>📝 进程信息与参数设置</span>
      </div>
    </template>

    <!-- 1. 进程数量输入 -->
    <div class="input-item">
      <span class="label">进程数量：</span>
      <el-input-number v-model="processCount" :min="1" :max="10" @change="initProcessList" />
    </div>

    <!-- 2. 进程信息表格（核心修复：加宽列宽、自适应输入框） -->
    <div class="table-wrapper">
      <el-table :data="processList" border style="width: 100%; min-width: 680px;" size="default">
        <el-table-column prop="name" label="进程名" width="90" align="center" />
        <el-table-column prop="initPriority" label="初始优先级 (1-100)" width="180" align="center">
          <template #default="{ row }">
            <el-input-number 
              v-model="row.initPriority" 
              :min="1" 
              :max="100" 
              style="width: 100%"
            />
          </template>
        </el-table-column>
        <el-table-column prop="arrivalTime" label="到达时间" width="140" align="center">
          <template #default="{ row }">
            <el-input-number 
              v-model="row.arrivalTime" 
              :min="0" 
              style="width: 100%"
            />
          </template>
        </el-table-column>
        <el-table-column prop="totalTime" label="总运行时间" width="140" align="center">
          <template #default="{ row }">
            <el-input-number 
              v-model="row.totalTime" 
              :min="1" 
              style="width: 100%"
            />
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 3. 调度算法与参数选择 -->
    <div class="input-item">
      <span class="label">调度算法：</span>
      <el-select v-model="schedulerParams.algorithmType" placeholder="请选择" style="width: 280px;">
        <el-option label="基于优先级的时间片轮转" value="priority_rr" />
        <el-option label="多级反馈队列轮转" value="multilevel_feedback" />
      </el-select>
    </div>
    <div class="input-item">
      <span class="label">时间片大小：</span>
      <el-input-number v-model="schedulerParams.timeSlice" :min="1" />
    </div>
    <div class="input-item" v-if="schedulerParams.algorithmType === 'priority_rr'">
      <span class="label">优先级下降步长：</span>
      <el-input-number v-model="schedulerParams.priorityStep" :min="1" />
    </div>

    <!-- 4. 操作按钮 -->
    <div class="button-group">
      <el-button type="primary" @click="handleCreateProcesses">创建进程</el-button>
      <el-button type="success" @click="handleInitScheduler">初始化调度器</el-button>
      <el-button @click="handleReset">重置系统</el-button>
    </div>
  </el-card>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useProcessStore } from '../stores/processStore'
import { ElMessage } from 'element-plus'

const store = useProcessStore()

// 1. 定义响应式数据
const processCount = ref(3) // 默认3个进程
const processList = ref([]) // 进程信息列表
const schedulerParams = reactive({
  algorithmType: 'priority_rr',
  timeSlice: 2,
  priorityStep: 2
})

// 2. 初始化进程列表
const initProcessList = () => {
  processList.value = []
  for (let i = 0; i < processCount.value; i++) {
    processList.value.push({
      name: `P${i + 1}`,
      initPriority: i + 1,
      arrivalTime: i,
      totalTime: 3
    })
  }
}

// 3. 操作方法
const handleCreateProcesses = async () => {
  try {
    // 给每个进程加上 name
    const processes = processList.value.map((p, i) => ({
      ...p,
      name: `P${i + 1}`
    }))
    await store.createProcesses(processes)
    ElMessage.success('进程创建成功！')
  } catch (e) {
    ElMessage.error('创建失败：' + e.message)
  }
}

const handleInitScheduler = async () => {
  try {
    await store.initScheduler(schedulerParams)
    ElMessage.success('调度器初始化成功！')
  } catch (e) {
    ElMessage.error('初始化失败：' + e.message)
  }
}

const handleReset = async () => {
  try {
    await store.resetSystem()
    ElMessage.success('系统重置成功！')
  } catch (e) {
    ElMessage.error('重置失败：' + e.message)
  }
}

// 4. 页面加载时初始化
initProcessList()
</script>

<style scoped>
.panel-card {
  margin-bottom: 20px;
}
.input-item {
  margin: 14px 0;
  display: flex;
  align-items: center;
  gap: 10px;
  flex-wrap: wrap;
}
.label {
  font-weight: 500;
  min-width: 110px;
}
/* 表格滚动容器优化 */
.table-wrapper {
  width: 100%;
  overflow-x: auto;
  margin: 12px 0;
  border-radius: 4px;
}
.button-group {
  margin-top: 22px;
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}
</style>