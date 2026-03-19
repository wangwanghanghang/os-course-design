<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>🏆 调度结果统计</span>
      </div>
    </template>

    <div v-if="!store.result" class="empty-result">
      请先执行调度，所有进程完成后会显示结果
    </div>

    <div v-else>
      <el-alert
        :title="'平均周转时间：' + store.result.avgTurnaroundTime.toFixed(2)"
        type="success"
        show-icon
        style="margin-bottom: 20px;"
      />

      <el-table :data="store.result.processes" border stripe>
        <el-table-column prop="name" label="进程名" width="80" />
        <el-table-column prop="arrivalTime" label="到达时间" width="90" />
        <el-table-column prop="finishTime" label="结束时间" width="90" />
        <el-table-column prop="turnaroundTime" label="周转时间" width="90" />
      </el-table>
    </div>
  </el-card>
</template>

<script setup>
import { useProcessStore } from '../stores/processStore'

const store = useProcessStore()
</script>

<style scoped>
.panel-card {
  margin-bottom: 20px;
}
.empty-result {
  color: #999;
  text-align: center;
  padding: 40px;
}
</style>