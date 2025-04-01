<template>
  <div class="projects-container">
    <div class="header">
      <h2>项目管理</h2>
      <a-button type="primary">创建项目</a-button>
    </div>
    <a-table :data="projectList" :columns="columns" />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getProjectList } from '@/api/project'

const projectList = ref([])
const columns = [
  { title: '项目名称', dataIndex: 'name' },
  { title: '创建时间', dataIndex: 'createdAt' },
  { title: '状态', dataIndex: 'status' },
]

onMounted(async () => {
  const res = await getProjectList()
  projectList.value = res.data
})
</script>

<style lang="less" scoped>
.projects-container {
  padding: 20px;

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
}
</style>
