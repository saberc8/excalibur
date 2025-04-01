<template>
  <div class="home-page">
    <a-card title="项目概览">
      <template #extra>
        <a-button type="primary" @click="openCreateProject">
          创建调查项目
        </a-button>
      </template>
      <p>欢迎使用项目管理系统</p>
    </a-card>

    <a-modal
      v-model:visible="visible"
      title="创建调查项目"
      @ok="handleCreateProject"
      @cancel="handleCancel"
    >
      <a-form :model="form" :style="{ width: '100%' }">
        <a-form-item field="name" label="项目名称">
          <a-input v-model="form.name" placeholder="请输入项目名称" />
        </a-form-item>
        <a-form-item field="description" label="项目描述">
          <a-textarea
            v-model="form.description"
            placeholder="请输入项目描述"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Message } from '@arco-design/web-vue'
import { useSpaceStore } from '@/stores/space'
import { createProject } from '@/api/project'

const visible = ref(false)
const form = ref({
  name: '',
  description: '',
  status: 1
})

const openCreateProject = () => {
  const currentSpace = useSpaceStore.getCurrentSpace()
  if (!currentSpace) {
    Message.error('请先选择空间')
    return
  }
  form.value.spaceId = currentSpace.id
  visible.value = true
}

const handleCancel = () => {
  visible.value = false
  form.value = {
    name: '',
    description: '',
    status: 1
  }
}

const handleCreateProject = async () => {
  try {
    await createProject(form.value)
    Message.success('创建成功')
    handleCancel()
  } catch (error) {
    Message.error('创建失败')
  }
}
</script>

<style lang="less" scoped>
.home-page {
  padding: 16px;
  background-color: var(--color-fill-2);
  min-height: 100%;
}
</style>
