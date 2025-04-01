<template>
  <div class="tasks-container">
    <div class="header">
      <h2>任务管理</h2>
      <a-button type="primary">创建任务</a-button>
    </div>
    <a-table :data="taskList" :columns="columns" />
  </div>
</template>

<script setup>
import { ref } from 'vue'

const taskList = ref([])
const columns = [
  { title: '任务名称', dataIndex: 'name' },
  { title: '所属项目', dataIndex: 'projectName' },
  { title: '负责人', dataIndex: 'assignee' },
  { title: '截止日期', dataIndex: 'dueDate' },
  { title: '状态', dataIndex: 'status' },
]
</script>

<style lang="less" scoped>
.tasks-container {
  padding: 20px;

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
}
</style>
