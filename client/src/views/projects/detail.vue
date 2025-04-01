<template>
  <div class="project-detail">
    <div class="header">
      <h2>{{ project && project.name ? project.name : '项目详情' }}</h2>
      <a-button type="primary" @click="showCreateTaskModal">创建任务</a-button>
    </div>

    <a-card title="任务列表" :bordered="false">
      <a-tabs default-active-key="1">
        <a-tab-pane key="1" title="全部">
          <a-table :data="tasks" :loading="loading" row-key="id">
            <template #columns>
              <a-table-column title="任务标题" data-index="title">
                <template #cell="{ record }">
                  <a @click="viewTaskDetail(record)">{{ record.title }}</a>
                </template>
              </a-table-column>
              <a-table-column title="类型" data-index="type">
                <template #cell="{ record }">
                  <a-tag :color="getTaskTypeColor(record.type)">
                    {{ getTaskTypeName(record.type) }}
                  </a-tag>
                </template>
              </a-table-column>
              <a-table-column title="优先级" data-index="priority">
                <template #cell="{ record }">
                  <span style="display: none;">{{ JSON.stringify(record) }}</span>
                  <a-tag :color="getPriorityColor(Number(record.priority))">
                    {{ getPriorityName(Number(record.priority)) }}
                  </a-tag>
                </template>
              </a-table-column>
              <a-table-column title="状态" data-index="status">
                <template #cell="{ record }">
                  <a-tag :color="getStatusColor(record.status)">
                    {{ getStatusName(record.status) }}
                  </a-tag>
                </template>
              </a-table-column>
              <a-table-column title="负责人" data-index="assigneeId">
                <template #cell="{ record }">
                  {{ record.assigneeId ? getMemberName(record.assigneeId) : '未分配' }}
                </template>
              </a-table-column>
              <a-table-column title="开始日期" data-index="startDate">
                <template #cell="{ record }">
                  {{ formatDate(record.startDate) }}
                </template>
              </a-table-column>
              <a-table-column title="结束日期" data-index="endDate">
                <template #cell="{ record }">
                  {{ formatDate(record.endDate) }}
                </template>
              </a-table-column>
              <a-table-column title="操作">
                <template #cell="{ record }">
                  <a-space>
                    <a-button type="text" size="small" @click="editTask(record)">编辑</a-button>
                    <a-popconfirm
                      content="确定要删除此任务吗？"
                      @ok="confirmDelete(record.id)"
                    >
                      <a-button type="text" status="danger" size="small">删除</a-button>
                    </a-popconfirm>
                  </a-space>
                </template>
              </a-table-column>
            </template>
          </a-table>
        </a-tab-pane>
        <a-tab-pane key="2" title="Bug"></a-tab-pane>
        <a-tab-pane key="3" title="新需求"></a-tab-pane>
      </a-tabs>
    </a-card>

    <a-modal
      v-model:visible="taskModalVisible"
      :title="editingTask.id ? '编辑任务' : '创建任务'"
      @ok="submitTask"
      @cancel="cancelTask"
    >
      <a-form :model="editingTask" ref="taskForm" label-align="left">
        <a-form-item field="title" label="标题" required>
          <a-input v-model="editingTask.title" placeholder="请输入任务标题" />
        </a-form-item>
        <a-form-item field="type" label="任务类型" required>
          <a-select v-model="editingTask.type" placeholder="请选择任务类型">
            <a-option :value="1">Bug</a-option>
            <a-option :value="2">新需求</a-option>
            <a-option :value="3">优化</a-option>
            <a-option :value="4">文档</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="content" label="内容描述">
          <a-textarea v-model="editingTask.content" placeholder="请输入任务描述" />
        </a-form-item>
        <a-form-item field="priority" label="优先级">
          <a-select v-model="editingTask.priority" placeholder="请选择优先级">
            <a-option :value="1">低</a-option>
            <a-option :value="2">中</a-option>
            <a-option :value="3">高</a-option>
            <a-option :value="4">紧急</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="status" label="状态">
          <a-select v-model="editingTask.status" placeholder="请选择状态">
            <a-option :value="1">待处理</a-option>
            <a-option :value="2">进行中</a-option>
            <a-option :value="3">已完成</a-option>
            <a-option :value="4">已关闭</a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="assigneeId" label="负责人">
          <a-select v-model="editingTask.assigneeId" placeholder="请选择负责人">
            <a-option v-for="member in projectMembers" :key="member.id" :value="member.id">
              {{ member.name }}
            </a-option>
          </a-select>
        </a-form-item>
        <a-form-item field="dateRange" label="时间范围">
          <a-range-picker
            v-model="dateRange"
            style="width: 100%"
            show-time
            format="YYYY-MM-DD HH:mm:ss"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { getProject } from '@/api/project'
import { getTaskList, createTask, updateTask, deleteTask } from '@/api/task'

const route = useRoute()
const project = ref({}) // 确保初始化为空对象而不是null
const tasks = ref([])
const loading = ref(false)
const taskModalVisible = ref(false)
const projectMembers = ref([])
const dateRange = ref([])

const editingTask = ref({
  title: '',
  content: '',
  projectId: 0,
  assigneeId: null,
  priority: 2,
  status: 1,
  type: 1,
  startDate: '',
  endDate: ''
})

const taskForm = ref(null)

const fetchProjectDetail = async () => {
  try {
    const projectId = route.params.id
    if (!projectId) {
      console.error('项目ID不存在')
      Message.error('项目ID不存在')
      return
    }
    
    const res = await getProject(projectId)
    if (res && res.data) {
      project.value = res.data
      projectMembers.value = res.data.members || []
    } else {
      project.value = { name: '未知项目' }
      projectMembers.value = []
    }
  } catch (error) {
    console.error('获取项目详情失败', error)
    Message.error('获取项目详情失败')
    // 设置默认值，防止模板渲染错误
    project.value = { name: '加载失败' }
    projectMembers.value = []
  }
}

const fetchTasks = async () => {
  loading.value = true
  try {
    const projectId = route.params.id
    const res = await getTaskList({ projectId })
    
    // 确保获取的数据中优先级和状态是数字类型
    if (res.data && Array.isArray(res.data)) {
      tasks.value = res.data.map(task => ({
        ...task,
        priority: task.priority !== null && task.priority !== undefined 
          ? Number(task.priority) 
          : 0,
        status: task.status !== null && task.status !== undefined 
          ? Number(task.status) 
          : 0,
        type: task.type !== null && task.type !== undefined 
          ? Number(task.type) 
          : 0
      }))
    } else {
      tasks.value = []
    }
    
    console.log('处理后的任务数据:', tasks.value)
  } catch (error) {
    console.error('获取任务列表失败', error)
    Message.error('获取任务列表失败')
    tasks.value = []
  } finally {
    loading.value = false
  }
}

const showCreateTaskModal = () => {
  editingTask.value = {
    title: '',
    content: '',
    projectId: parseInt(route.params.id),
    assigneeId: null,
    priority: 2,
    status: 1,
    type: 1,
    startDate: '',
    endDate: ''
  }
  dateRange.value = []
  taskModalVisible.value = true
}

const editTask = (task) => {
  editingTask.value = { ...task }
  dateRange.value = task.startDate && task.endDate ? [task.startDate, task.endDate] : []
  taskModalVisible.value = true
}

const submitTask = async () => {
  if (!editingTask.value.title) {
    Message.error('请输入任务标题')
    return
  }

  if (dateRange.value && dateRange.value.length === 2) {
    editingTask.value.startDate = dateRange.value[0]
    editingTask.value.endDate = dateRange.value[1]
  }

  try {
    if (editingTask.value.id) {
      await updateTask(editingTask.value)
      Message.success('更新任务成功')
    } else {
      await createTask(editingTask.value)
      Message.success('创建任务成功')
    }
    taskModalVisible.value = false
    fetchTasks()
  } catch (error) {
    console.error('保存任务失败', error)
    Message.error('保存任务失败')
  }
}

const cancelTask = () => {
  taskModalVisible.value = false
}

const confirmDelete = async (id) => {
  try {
    await deleteTask(id)
    Message.success('删除任务成功')
    fetchTasks()
  } catch (error) {
    console.error('删除任务失败', error)
    Message.error('删除任务失败')
  }
}

const viewTaskDetail = (task) => {
  console.log('查看任务详情', task)
}

const formatDate = (dateStr) => {
  if (!dateStr) return '未设置';
  try {
    const date = new Date(dateStr);
    return date.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    });
  } catch (error) {
    console.error('日期格式化错误', error);
    return dateStr;
  }
};

const getMemberName = (memberId) => {
  const member = projectMembers.value.find(m => m.id === memberId);
  return member ? member.name : '未知成员';
};

const getTaskTypeName = (type) => {
  const types = {
    0: '一般任务',
    1: 'Bug',
    2: '新需求',
    3: '优化',
    4: '文档'
  }
  return types[type] !== undefined ? types[type] : '未知';
};

const getTaskTypeColor = (type) => {
  const colors = {
    0: 'gray',
    1: 'red',
    2: 'blue',
    3: 'green',
    4: 'purple'
  }
  return colors[type] !== undefined ? colors[type] : 'gray';
};

const getPriorityName = (priority) => {
  console.log('原始优先级值:', priority, 'typeof:', typeof priority)
  const priorityNum = Number(priority)
  console.log('转换后优先级值:', priorityNum, 'typeof:', typeof priorityNum)
  
  const priorities = {
    0: '普通',
    1: '低',
    2: '中',
    3: '高',
    4: '紧急'
  }
  
  return priorities[priorityNum] !== undefined ? priorities[priorityNum] : '未知(' + priority + ')'
};

const getPriorityColor = (priority) => {
  const priorityNum = Number(priority)
  
  const colors = {
    0: 'gray',
    1: 'gray',
    2: 'blue',
    3: 'orange',
    4: 'red'
  }
  
  return colors[priorityNum] !== undefined ? colors[priorityNum] : 'gray'
};

const getStatusName = (status) => {
  const statuses = {
    0: '待处理',
    1: '待处理',
    2: '进行中',
    3: '已完成',
    4: '已关闭'
  }
  return statuses[status] !== undefined ? statuses[status] : '未知';
};

const getStatusColor = (status) => {
  const colors = {
    0: 'gray',
    1: 'gray',
    2: 'blue',
    3: 'green',
    4: 'red'
  }
  return colors[status] !== undefined ? colors[status] : 'gray';
};

onMounted(() => {
  fetchProjectDetail()
  fetchTasks()
})

watch(() => route.params.id, () => {
  fetchProjectDetail()
  fetchTasks()
})
</script>

<style lang="less" scoped>
.project-detail {
  padding: 20px;
  
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 24px;
    color: #333;
  }
}
</style>
