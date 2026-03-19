<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>📦 队列可视化</span>
      </div>
    </template>

    <div class="queue-container">
      <!-- 就绪队列 -->
      <div class="queue-box">
        <h4>✅ 就绪队列 (Ready)</h4>
        <div class="queue-cards">
          <!-- 修复：加 ?. 可选链，并且判断 !store.readyQueue -->
          <div v-if="!store.readyQueue || store.readyQueue.length === 0" class="empty-text">队列为空</div>
          <div v-else v-for="pcb in store.readyQueue" :key="pcb.name" class="process-card ready">
            {{ pcb.name }}<br/>
            <small>优先级: {{ pcb.currPriority }}</small>
          </div>
        </div>
      </div>

      <!-- 当前执行 -->
      <div class="queue-box">
        <h4>⚡ 当前执行 (Execute)</h4>
        <div class="queue-cards">
          <!-- 修复：先判断 !store.currentProcess -->
          <div v-if="!store.currentProcess" class="empty-text">CPU空闲</div>
          <div v-else class="process-card execute">
            {{ store.currentProcess.name }}<br/>
            <small>剩余: {{ store.currentProcess.remainTime }}</small>
          </div>
        </div>
      </div>

      <!-- 等待队列 -->
      <div class="queue-box">
        <h4>⏸️ 等待队列 (Wait)</h4>
        <div class="queue-cards">
          <!-- 修复：加 ?. 可选链，并且判断 !store.waitQueue -->
          <div v-if="!store.waitQueue || store.waitQueue.length === 0" class="empty-text">队列为空</div>
          <div v-else v-for="pcb in store.waitQueue" :key="pcb.name" class="process-card wait">
            {{ pcb.name }}
          </div>
        </div>
      </div>
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
.queue-container {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}
.queue-box {
  flex: 1;
  min-width: 200px;
}
.queue-cards {
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
  min-height: 80px;
  border: 1px dashed #ddd;
  padding: 10px;
  border-radius: 4px;
}
.process-card {
  width: 80px;
  height: 60px;
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: white;
  font-weight: bold;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
.ready { background-color: #67C23A; }
.execute { background-color: #F56C6C; }
.wait { background-color: #E6A23C; }
.empty-text {
  color: #999;
  line-height: 60px;
}
</style>