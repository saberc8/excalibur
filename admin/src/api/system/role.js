import request from '@/utils/request'

// 创建角色
export function addRole(data) {
  return request({
    url: '/role/create',
    method: 'post',
    data: data
  })
}

// 获取单个角色信息
export function getRole(roleId) {
  return request({
    url: '/role',
    method: 'get',
    params: { id: roleId }
  })
}

// 获取角色列表
export function roleList(query) {
  return request({
    url: '/role/list',
    method: 'get',
    params: query
  })
}

// 更新角色信息
export function updateRole(data) {
  return request({
    url: '/role/update',
    method: 'post',
    data: data
  })
}

// 删除角色
export function delRole(roleId) {
  return request({
    url: '/role/delete',
    method: 'post',
    data: { id: roleId }
  })
}

// 为角色绑定菜单
export function bindMenu(data) {
  return request({
    url: '/role/bindMenu',
    method: 'post',
    data: data
  })
}

// 为角色解绑菜单
export function unbindMenu(data) {
  return request({
    url: '/role/unbindMenu',
    method: 'post',
    data: data
  })
}

// 获取角色拥有的菜单列表
export function getRoleMenus(roleId) {
  return request({
    url: '/role/menus',
    method: 'get',
    params: { id: roleId }
  })
}

// 获取拥有该角色的用户列表
export function getRoleUsers(roleId) {
  return request({
    url: '/role/users',
    method: 'get',
    params: { id: roleId }
  })
}
