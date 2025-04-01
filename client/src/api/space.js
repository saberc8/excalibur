import request from '@/utils/request'

export function getSpaceList() {
  return request({
    url: '/space/list',
    method: 'get'
  })
}

export function getSpace(id) {
  return request({
    url: '/space',
    method: 'get',
    params: { id }
  })
}

export function createSpace(data) {
  return request({
    url: '/space/create',
    method: 'post',
    data
  })
}

export function updateSpace(data) {
  return request({
    url: '/space/update',
    method: 'post',
    data
  })
}

export function deleteSpace(id) {
  return request({
    url: '/space/delete',
    method: 'post',
    data: { id }
  })
}
