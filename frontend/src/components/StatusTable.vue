<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>📊 进程实时状态</span>
      </div>
    </template>

    <el-table :data="store.allProcesses" border stripe>
      <el-table-column prop="name" label="进程名" width="80" />
      <el-table-column prop="initPriority" label="初始优先级" width="100" />
      <el-table-column prop="currPriority" label="当前优先级" width="100" />
      <el-table-column prop="arrivalTime" label="到达时间" width="90" />
      <el-table-column prop="totalTime" label="总运行时间" width="90" />
      <el-table-column prop="usedTime" label="已用时间" width="90" />
      <el-table-column prop="remainTime" label="剩余时间" width="90" />
      <el-table-column prop="state" label="状态" width="100">
        <template #default="{ row }">
          <el-tag :type="getStateType(row.state)">
            {{ getStateText(row.state) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="startTime" label="开始时间" width="90" />
      <el-table-column prop="finishTime" label="结束时间" width="90" />
      <el-table-column prop="turnaroundTime" label="周转时间" width="90" />
    </el-table>
  </el-card>
</template>

<script setup>
import { useProcessStore } from '../stores/processStore'

const store = useProcessStore()

// 状态映射：数字 -> 文字
const getStateText = (state) => {
  const map = { 0: 'E执行', 1: 'R就绪', 2: 'W等待', 3: 'F完成' }
  return map[state] || '未知'
}

// 状态映射：数字 -> Element Plus Tag 颜色
const getStateType = (state) => {
  const map = { 0: 'danger', 1: 'success', 2: 'warning', 3: 'info' }
  return map[state] || ''
}
</script>

<style scoped>
.panel-card {
  margin-bottom: 20px;
}
</style>