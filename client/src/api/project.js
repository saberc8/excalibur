import request from '@/utils/request'

export function getProjectList(params) {
  return request({
    url: '/project/list',
    method: 'get',
    params
  })
}

export function getProject(id) {
  return request({
    url: '/project',
    method: 'get',
    params: { id }
  })
}

export function createProject(data) {
  return request({
    url: '/project/create',
    method: 'post',
    data
  })
}

export function updateProject(data) {
  return request({
    url: '/project/update',
    method: 'post',
    data
  })
}

export function deleteProject(id) {
  return request({
    url: '/project/delete',
    method: 'post',
    data: { id }
  })
}
