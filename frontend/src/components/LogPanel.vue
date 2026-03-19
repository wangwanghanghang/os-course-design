<template>
  <el-card class="panel-card">
    <template #header>
      <div class="card-header">
        <span>📜 执行日志</span>
      </div>
    </template>

    <div class="log-container" ref="logContainer">
      <div v-for="(log, index) in store.logs" :key="index" class="log-item">
        {{ log }}
      </div>
    </div>
  </el-card>
</template>

<script setup>
import { useProcessStore } from '../stores/processStore'
import { watch, ref, nextTick } from 'vue'

const store = useProcessStore()
const logContainer = ref(null)

// 监听 logs 变化，自动滚动到底部
watch(() => store.logs, async () => {
  await nextTick()
  if (logContainer.value) {
    logContainer.value.scrollTop = logContainer.value.scrollHeight
  }
}, { deep: true })
</script>

<style scoped>
.panel-card {
  margin-bottom: 20px;
}
.log-container {
  height: 200px;
  overflow-y: auto;
  border: 1px solid #ddd;
  padding: 10px;
  border-radius: 4px;
  background-color: #f5f7fa;
}
.log-item {
  font-size: 13px;
  color: #333;
  margin: 4px 0;
  font-family: 'Consolas', monospace;
}
</style>