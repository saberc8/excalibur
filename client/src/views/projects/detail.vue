<template>
  <div class="project-detail">
    <div class="header">
      <h2>{{ project.name }}</h2>
    </div>
    <a-tabs>
      <a-tab-pane key="overview" title="概览">
        <p>{{ project.description }}</p>
      </a-tab-pane>
      <a-tab-pane key="tasks" title="任务列表">
        <a-table :data="tasks" :columns="taskColumns" />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const project = ref({})
const tasks = ref([])

const taskColumns = [
  { title: '任务名称', dataIndex: 'name' },
  { title: '负责人', dataIndex: 'assignee' },
  { title: '状态', dataIndex: 'status' },
]
</script>

<style lang="less" scoped>
.project-detail {
  padding: 20px;
  
  .header {
    margin-bottom: 24px;
  }
}
</style>
