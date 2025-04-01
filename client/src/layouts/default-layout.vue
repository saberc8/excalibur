<template>
  <a-layout class="layout">
    <a-layout-sider collapsible v-model:collapsed="collapsed" :width="220">
      <div class="space-header">
        <template v-if="currentSpace">
          <a-dropdown trigger="click" position="br">
            <div class="space-selector">
              <span>{{ currentSpace.name }}</span>
              <icon-down />
            </div>
            <template #content>
              <div class="space-dropdown">
                <div class="current-space">
                  <div class="space-name">{{ currentSpace.name }}</div>
                  <div class="space-actions">
                    <a-button type="text" size="small" @click="handleSpaceSettings">
                      <template #icon><icon-settings /></template>
                      空间设置
                    </a-button>
                    <a-button type="text" size="small" @click="handleInviteMembers">
                      <template #icon><icon-user-add /></template>
                      邀请成员
                    </a-button>
                  </div>
                </div>
                <a-divider style="margin: 8px 0" />
                <div class="other-spaces">
                  <div class="section-title">其他空间</div>
                  <div class="space-list">
                    <a-doption
                      v-for="space in otherSpaces"
                      :key="space.id"
                      class="space-item"
                      @click="switchSpace(space)"
                    >
                      {{ space.name }}
                    </a-doption>
                  </div>
                </div>
                <a-divider style="margin: 8px 0" />
                <a-doption @click="handleCreateSpace">
                  <icon-plus /> 创建新空间
                </a-doption>
              </div>
            </template>
          </a-dropdown>
        </template>
        <div v-else class="empty-space">
          <div class="empty-text">还没有空间</div>
          <a-button type="primary" size="small" @click="handleCreateSpace">
            创建空间
          </a-button>
        </div>
      </div>
      
      <template v-if="currentSpace">
        <a-menu
          :selected-keys="selectedKeys"
          :style="{ width: '100%' }"
          @menu-item-click="onMenuItemClick"
        >
          <a-menu-item key="home">
            <template #icon><icon-home /></template>
            控制台
          </a-menu-item>
          <a-menu-item key="page1">
            <template #icon><icon-file /></template>
            页面1
          </a-menu-item>
          <a-menu-item key="page2">
            <template #icon><icon-file /></template>
            页面2
          </a-menu-item>
          <a-menu-item key="projects">
            <template #icon><icon-folder /></template>
            项目管理
          </a-menu-item>
          <a-menu-item key="tasks">
            <template #icon><icon-schedule /></template>
            任务管理
          </a-menu-item>

          <a-divider style="margin: 8px 0" />
          
          <div class="menu-section-title">我的项目</div>
          <a-menu-item 
            v-for="project in projectList" 
            :key="`project-${project.id}`"
          >
            <template #icon><icon-file /></template>
            {{ project.name }}
          </a-menu-item>
        </a-menu>
      </template>
    </a-layout-sider>
    <a-layout>
      <a-layout-header>
        <a-space :size="24" class="header-wrapper">
          <a-button type="text" @click="collapsed = !collapsed">
            <template #icon>
              <icon-menu-fold v-if="collapsed" />
              <icon-menu-unfold v-else />
            </template>
          </a-button>
          <a-space class="header-right">
            <a-dropdown>
              <a-avatar :size="32">
                <icon-user />
              </a-avatar>
              <template #content>
                <a-doption class="user-option">{{ username }}</a-doption>
                <a-doption class="user-option" @click="handleLogout">退出登录</a-doption>
              </template>
            </a-dropdown>
          </a-space>
        </a-space>
      </a-layout-header>
      <a-layout-content>
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
  <create-space-modal
    v-model:visible="createSpaceVisible"
    @success="submitCreateSpace"
  />
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  IconHome,
  IconFolder,
  IconSchedule,
  IconMenuFold,
  IconMenuUnfold,
  IconUser,
  IconDown,
  IconPlus,
  IconSettings,
  IconUserAdd,
  IconFile,
} from '@arco-design/web-vue/es/icon'
import { getSpaceList, createSpace } from '@/api/space'
import { getProjectList } from '@/api/project'
import { Message } from '@arco-design/web-vue'
import CreateSpaceModal from '@/components/space/create-space-modal.vue'
import { useSpaceStore } from '@/stores/space'

const router = useRouter()
const route = useRoute()
const collapsed = ref(false)
const username = ref('管理员')

const selectedKeys = computed(() => [route.name])

const onMenuItemClick = (key) => {
  if (key.startsWith('project-')) {
    const projectId = key.replace('project-', '')
    router.push(`/${currentSpace.value.path}/project/${projectId}`)
  } else {
    router.push(`/${currentSpace.value.path}/${key}`)
  }
}

const handleLogout = () => {
  // 实现登出逻辑
  router.push('/login')
}

const currentSpace = ref(null)
const spaceList = ref([])

onMounted(async () => {
  const res = await getSpaceList()
  spaceList.value = res.data
  if (spaceList.value.length > 0) {
    // 使用 store 保存当前空间
    const savedSpace = useSpaceStore.getCurrentSpace()
    currentSpace.value = savedSpace || spaceList.value[0]
    useSpaceStore.setCurrentSpace(currentSpace.value)
    localStorage.setItem('lastSpacePath', currentSpace.value.path)
    // 如果当前在根路径，重定向到空间首页
    if (route.path === '/') {
      router.push(`/${currentSpace.value.path}/home`)
    }
  }
})

const createSpaceVisible = ref(false)
const spaceFormRef = ref(null)
const spaceForm = ref({
  name: '',
  description: '',
  status: 1
})

const handleCreateSpace = () => {
  createSpaceVisible.value = true
  spaceForm.value = {
    name: '',
    description: '',
    status: 1
  }
}

const submitCreateSpace = async (formData) => {
  try {
    const res = await createSpace(formData)
    if (res.data) {
      Message.success('创建成功')
      createSpaceVisible.value = false
      // 刷新空间列表
      const listRes = await getSpaceList()
      spaceList.value = listRes.data
      if (!currentSpace.value && spaceList.value.length > 0) {
        currentSpace.value = spaceList.value[0]
      }
    }
  } catch (err) {
    console.error(err)
  }
}

const otherSpaces = computed(() => {
  return spaceList.value.filter(space => space.id !== currentSpace.value?.id)
})

const switchSpace = (space) => {
  currentSpace.value = space
  useSpaceStore.setCurrentSpace(space)
  localStorage.setItem('lastSpacePath', space.path)
  // 切换空间时，跳转到对应空间的首页
  router.push(`/${space.path}/home`)
}

const handleSpaceSettings = () => {
  // 实现空间设置逻辑
}

const handleInviteMembers = () => {
  // 实现邀请成员逻辑
}

const projectList = ref([])

// 添加获取项目列表的方法
const fetchProjectList = async () => {
  if (currentSpace.value) {
    try {
      const res = await getProjectList({ spaceId: currentSpace.value.id })
      projectList.value = res.data
    } catch (err) {
      console.error('获取项目列表失败:', err)
    }
  }
}

// 监听当前空间变化
watch(() => currentSpace.value, () => {
  fetchProjectList()
}, { immediate: true })
</script>

<style lang="less" scoped>
.layout {
  height: 100vh;
}

.space-header {
  height: 64px;
  padding: 0 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.space-selector {
  display: flex;
  align-items: center;
  justify-content: space-between;  /* 两端对齐 */
  gap: 8px;
  cursor: pointer;
  color: #000000;
  padding: 8px 12px;  /* 增大内边距 */
  border-radius: 4px;
  width: 100%;  /* 占满容器宽度 */
  font-weight: 600;  /* 字体加粗 */
  font-size: 16px;  /* 增大字体大小 */

  &:hover {
    background: rgba(255, 255, 255, 0.1);
  }
}

.space-dropdown {
  padding: 12px 0;
  min-width: 280px;
  overflow-y: auto;
}

.current-space {
  padding: 12px 16px;
  
  .space-name {
    font-weight: bold;
    margin-bottom: 8px;
    font-size: 16px;
  }
  
  .space-actions {
    display: flex;
    flex-direction: row;  /* 改为水平方向 */
    justify-content: space-between; /* 两端对齐 */
    gap: 12px; /* 按钮之间的间距 */

    :deep(.arco-btn) {
      flex: 1; /* 按钮等宽 */
    }
  }
}

.other-spaces {
  padding: 0 16px;
  
  .section-title {
    color: #86909c;
    font-size: 12px;
    padding: 4px 0;
  }
}

.space-list {
  max-height: 300px;
  overflow-y: auto;
  
  .space-item {
    color: #333;
    padding: 8px 12px;
    margin: 2px 0;
    border-radius: 4px;
    
    &:hover {
      background-color: var(--color-fill-2);
    }
  }
}

.empty-space {
  text-align: center;
  color: #fff;
  
  .empty-text {
    margin-bottom: 8px;
    font-size: 14px;
  }
}

.logo {
  height: 64px;
  line-height: 64px;
  text-align: center;
  color: #fff;
  font-size: 16px;
  font-weight: bold;
}

.header-wrapper {
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.header-right {
  margin-right: 16px;
}

.user-option {
  color: #333;
}

.menu-section-title {
  padding: 8px 16px;
  color: #86909c;
  font-size: 12px;
}
</style>
