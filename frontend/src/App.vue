<template>
  <div id="app">
    <div class="header">
      <h1>🎯 操作系统课程设计 - 进程调度模拟器</h1>
      <p>Vue3 + Go | 基于优先级的时间片轮转 & 多级反馈队列</p>
    </div>

    <div class="content">
      <!-- 左侧：输入和控制 -->
      <div class="left-col">
        <InputPanel />
        <ControlPanel />
        <QueueVisual />
      </div>

      <!-- 右侧：状态、日志、结果 -->
      <div class="right-col">
        <StatusTable />
        <LogPanel />
        <ResultPanel />
      </div>
    </div>
  </div>
</template>

<script setup>
// 引入所有组件
import InputPanel from './components/InputPanel.vue'
import ControlPanel from './components/ControlPanel.vue'
import StatusTable from './components/StatusTable.vue'
import QueueVisual from './components/QueueVisual.vue'
import LogPanel from './components/LogPanel.vue'
import ResultPanel from './components/ResultPanel.vue'
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
body {
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', Arial, sans-serif;
  background-color: #f0f2f5;
}
#app {
  padding: 20px;
}
.header {
  text-align: center;
  margin-bottom: 30px;
  background-color: #409EFF;
  color: white;
  padding: 20px;
  border-radius: 8px;
}
.header h1 {
  margin-bottom: 10px;
}
.content {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}
.left-col {
  flex: 1.3; /* 从1.2再加大一点，给左侧更多空间 */
  min-width: 700px; /* 最小宽度再加大，保证表格能完整显示 */
}
.right-col {
  flex: 1.7;
  min-width: 600px;
}
</style>
