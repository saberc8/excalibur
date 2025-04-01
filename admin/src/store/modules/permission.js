import auth from '@/plugins/auth'
import { constantRoutes } from '@/router'
import { getUserMenus } from '@/api/system/menu'
import Layout from '@/layout/index'
import InnerLink from '@/layout/components/InnerLink'

// 匹配views里面所有的.vue文件
const modules = import.meta.glob('./../../views/**/*.vue')

const usePermissionStore = defineStore('permission', {
	state: () => ({
		// routes: [],
		sidebarRouters: [],
	}),
	actions: {
		// setRoutes(routes) {
		// 	this.routes = constantRoutes.concat(routes)
		// },
		setSidebarRouters(routes) {
			this.sidebarRouters = routes
		},
		generateRoutes() {
			return new Promise((resolve) => {
				// 向后端请求路由数据
				getUserMenus().then((res) => {
					console.log(res)
					const sdata = JSON.parse(JSON.stringify(res))
					const sidebarRoutes = filterAsyncRouter(sdata)
					sidebarRoutes.sort((a, b) => {
						return a.orderNum - b.orderNum
					})
					sidebarRoutes.forEach((item) => {
						if (item.children && item.children.length) {
							item.children.sort((a, b) => {
								return a.orderNum - b.orderNum
							})
						}
					})
					sidebarRoutes.forEach((route) => {
						route.meta = {
							title: route.name,
							icon: route.icon,
							noCache: route.isCatch === 0 ? false : true,
						}
						if (route.children && route.children.length) {
							route.children.forEach((child) => {
								child.meta = {
									title: child.name,
									icon: child.icon,
									noCache: route.isCatch === 0 ? false : true,
								}
							})
						}
					})
					console.log(sidebarRoutes, 'sidebarRoutes')
					this.setSidebarRouters(constantRoutes.concat(sidebarRoutes))
					resolve(sidebarRoutes)
				})
			})
		},
	},
})

// 遍历后台传来的路由字符串，转换为组件对象
// q: type的作用
// a: type为true时，children不会被删除
function filterAsyncRouter(asyncRouterMap, type = false) {
	const data = JSON.parse(JSON.stringify(asyncRouterMap))
	return data.filter((route) => {
		if (type && route.children) {
			route.children = filterChildren(route.children)
		}
		if (route.component) {
			if (route.component === 'Layout') {
				route.component = Layout
			} else if (route.component === 'InnerLink') {
				route.component = InnerLink
			} else {
				route.component = loadView(route.component)
			}
		}
		if (route.children != null && route.children && route.children.length) {
			route.children = filterAsyncRouter(route.children, type)
		} else {
			delete route['children']
			delete route['redirect']
		}
		return true
	})
}

function filterChildren(childrenMap, lastRouter = false) {
	var children = []
	childrenMap.forEach((el, index) => {
		if (el.children && el.children.length) {
			if (el.component === 'ParentView' && !lastRouter) {
				el.children.forEach((c) => {
					c.path = el.path + '/' + c.path
					if (c.children && c.children.length) {
						children = children.concat(filterChildren(c.children, c))
						return
					}
					children.push(c)
				})
				return
			}
		}
		if (lastRouter) {
			el.path = lastRouter.path + '/' + el.path
		}
		children = children.concat(el)
	})
	return children
}

// 动态路由遍历，验证是否具备权限
export function filterDynamicRoutes(routes) {
	const res = []
	routes.forEach((route) => {
		if (route.permissions) {
			if (auth.hasPermiOr(route.permissions)) {
				res.push(route)
			}
		} else if (route.roles) {
			if (auth.hasRoleOr(route.roles)) {
				res.push(route)
			}
		}
	})
	return res
}

export const loadView = (view) => {
	let res
	for (const path in modules) {
		const dir = path.split('views/')[1].split('.vue')[0]
		if (dir === view) {
			res = () => modules[path]()
		}
	}
	return res
}

export default usePermissionStore
