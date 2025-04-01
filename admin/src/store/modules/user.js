import { login, logout, getInfo } from '@/api/login'
import { getToken, setToken, removeToken } from '@/utils/auth'
import defAva from '@/assets/images/profile.jpeg'

const useUserStore = defineStore('user', {
	state: () => ({
		token: getToken(),
		userInfo: {},
		name: '',
		nickname: '',
		phonenumber: '',
		avatar: '',
		roles: [],
		permissions: [],
		dept: null,
	}),
	actions: {
		// 登录
		login(userInfo) {
			const username = userInfo.username.trim()
			const password = userInfo.password
			const captchaKey = userInfo.captchaKey
			const code = userInfo.code
			return new Promise((resolve, reject) => {
				login(username, password, code, captchaKey)
					.then((res) => {
						console.log(res, '登录')
						setToken(res.token)
						this.token = res.token
						const user = res.userInfo
						this.name = user.userName
						this.nickname = user.nickname
						this.avatar = user.avatar || defAva
						this.dept = ''
						this.phonenumber = ''
						resolve()
					})
					.catch((error) => {
						reject(error)
					})
			})
		},
		// 获取用户信息
		getInfo() {
			return new Promise((resolve, reject) => {
				getInfo()
					.then((res) => {
						console.log(res, '用户信息')
						const user = res
						const avatar =
							user.avatar == '' || user.avatar == null
								? defAva
								: import.meta.env.VITE_APP_BASE_API + user.avatar

						if (res.roles && res.roles.length > 0) {
							// 验证返回的roles是否是一个非空数组
							this.roles = res.roles
							this.permissions = res.permissions
						} else {
							this.roles = ['ROLE_DEFAULT']
						}
						this.name = user.userName
						this.nickname = user.nickname
						this.avatar = avatar
						this.dept = user.dept
						this.phonenumber = user.phonenumber || ''
						resolve(res)
					})
					.catch((error) => {
						reject(error)
					})
			})
		},
		// 退出系统
		logOut() {
			return new Promise((resolve, reject) => {
				logout(this.token)
					.then(() => {
						this.token = ''
						this.roles = []
						this.permissions = []
						removeToken()
						resolve()
					})
					.catch((error) => {
						reject(error)
					})
			})
		},
	},
})

export default useUserStore
