import request from '@/utils/request'

// 创建用户
export function addUser(data) {
  return request({
    url: '/user/create',
    method: 'post',
    data: data
  })
}

// 获取单个用户信息
export function getUser(userId) {
  return request({
    url: '/user',
    method: 'get',
    params: { id: userId }
  })
}

// 获取用户列表
export function userList(query) {
  return request({
    url: '/user/list',
    method: 'get',
    params: query
  })
}

// 更新用户信息
export function updateUser(data) {
  return request({
    url: '/user/update',
    method: 'post',
    data: data
  })
}

// 删除用户
export function delUser(userId) {
  return request({
    url: '/user/delete', 
    method: 'post',
    data: { id: userId }
  })
}

// 为用户绑定角色
export function bindUserRole(data) {
  return request({
    url: '/user/bindRole',
    method: 'post',
    data: data
  })
}

// 为用户解绑角色
export function unbindUserRole(data) {
  return request({
    url: '/user/unbindRole',
    method: 'post',
    data: data
  })
}

// 获取用户拥有的角色列表
export function getUserRoles(userId) {
  return request({
    url: '/user/roles',
    method: 'get',
    params: { id: userId }
  })
}

// 获取用户可访问的菜单列表
export function getUserMenus() {
  return request({
    url: '/user/menus',
    method: 'get'
  })
}

// POST /api/user/changePassword # 修改用户密码
export function changePassword(data) {
	return request({
		url: '/user/changePassword',
		method: 'post',
		data: data
	})
}

// 修改用户角色关联关系
export function updateUserRoles(data) {
  return request({
    url: '/user/bindRole',
    method: 'post',
    data: data
  })
}

// 获取所有角色列表
export function getAllRoles() {
  return request({
    url: '/role/list',
    method: 'get'
  })
}