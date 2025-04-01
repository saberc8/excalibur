import request from '@/utils/request'

// 创建菜单
export function addMenu(data) {
  return request({
    url: '/menu/create',
    method: 'post',
    data: data
  })
}

// 获取单个菜单信息
export function getMenu(menuId) {
  return request({
    url: '/menu',
    method: 'get',
    params: { id: menuId }
  })
}

// 获取菜单列表
export function menuList(query) {
  return request({
    url: '/menu/list',
    method: 'get',
    params: query
  })
}

// 更新菜单信息
export function updateMenu(data) {
  return request({
    url: '/menu/update',
    method: 'post',
    data: data
  })
}

// 删除菜单
export function delMenu(menuId) {
  return request({
    url: '/menu/delete',
    method: 'post',
    data: { id: menuId }
  })
}

// 获取拥有该菜单访问权限的角色列表
export function getMenuRoles(menuId) {
  return request({
    url: '/menu/roles',
    method: 'get',
    params: { id: menuId }
  })
}

// 获取菜单的树形结构
export function getMenuTree() {
  return request({
    url: '/menu/tree',
    method: 'get'
  })
}

// 查询菜单下拉树结构
export function treeselect() {
  return request({
    url: '/menu/treeselect',
    method: 'get',
  })
}

// 根据角色ID查询菜单下拉树结构
export function roleMenuTreeselect(roleId) {
  return request({
    url: '/menu/roleMenuTreeselect/' + roleId,
    method: 'get',
  })
}

// 根据角色ID获取菜单树
export function getMenuTreeByRoleId(roleId) {
  return request({
    url: '/menu/role-menu-tree?roleId=' + roleId,
    method: 'get'
  })
}

// 获取用户菜单
export const getUserMenus = () => {
  return request({
    url: '/user/menus',
    method: 'get',
  })
}

