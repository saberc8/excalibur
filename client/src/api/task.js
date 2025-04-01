import request from '@/utils/request'

export function getTaskList(params) {
  return request({
    url: '/task/list',
    method: 'get',
    params
  })
}

export function getTask(id) {
  return request({
    url: '/task',
    method: 'get',
    params: { id }
  })
}

export function createTask(data) {
  // 确保数据类型正确
  const taskData = { ...data }
  
  // 确保优先级是数字类型
  if (taskData.priority !== undefined) {
    taskData.priority = Number(taskData.priority)
  }
  
  // 确保状态是数字类型
  if (taskData.status !== undefined) {
    taskData.status = Number(taskData.status)
  }
  
  return request({
    url: '/task/create',
    method: 'post',
    data: taskData
  })
}

export function updateTask(data) {
  // 确保数据类型正确
  const taskData = { ...data }
  
  // 确保优先级是数字类型
  if (taskData.priority !== undefined) {
    taskData.priority = Number(taskData.priority)
  }
  
  // 确保状态是数字类型
  if (taskData.status !== undefined) {
    taskData.status = Number(taskData.status)
  }
  
  return request({
    url: '/task/update',
    method: 'post',
    data: taskData
  })
}

export function deleteTask(id) {
  return request({
    url: '/task/delete',
    method: 'post',
    data: { id }
  })
}
