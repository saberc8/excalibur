<template>
	<div class="login">
		<div class="container">
			<div class="wavy">
				<span></span>
				<span></span>
				<span></span>
			</div>
			<div class="content">
				<el-form
					ref="loginRef"
					:model="loginForm"
					:rules="loginRules"
					class="login-form"
				>
					<h3 class="title">evomix管理系统</h3>
					<el-form-item prop="username">
						<el-input
							v-model="loginForm.username"
							type="text"
							size="large"
							auto-complete="off"
							placeholder="账号"
						>
							<template #prefix
								><svg-icon icon-class="solar:user-circle-bold-duotone" class="el-input__icon input-icon"
							/></template>
						</el-input>
					</el-form-item>
					<el-form-item prop="password">
						<el-input
							v-model="loginForm.password"
							type="password"
							size="large"
							auto-complete="off"
							placeholder="密码"
							@keyup.enter="handleLogin"
						>
							<template #prefix
								><svg-icon
									icon-class="solar:lock-password-bold-duotone"
									class="el-input__icon input-icon"
							/></template>
						</el-input>
					</el-form-item>
					<el-form-item prop="code" v-if="captchaEnabled">
						<el-input
							v-model="loginForm.code"
							size="large"
							clearable
							auto-complete="off"
							placeholder="验证码"
							style="width: 63%"
							@keyup.enter="handleLogin"
						>
							<template #prefix
								><svg-icon
									icon-class="solar:shield-keyhole-bold-duotone"
									class="el-input__icon input-icon"
							/></template>
						</el-input>
						<div
							class="login-code login-code-img"
							v-html="codeUrl"
							@click="getCode"
						></div>
					</el-form-item>
					<el-checkbox
						v-model="loginForm.rememberMe"
						style="margin: 0px 0px 25px 0px"
						>记住密码</el-checkbox
					>
					<el-form-item style="width: 100%">
						<el-button
							:loading="loading"
							size="large"
							type="primary"
							style="width: 100%"
							@click.prevent="handleLogin"
						>
							<span v-if="!loading">登 录</span>
							<span v-else>登 录 中...</span>
						</el-button>
						<el-button
							style="width: 100%; margin: 10px 0;"
							size="large"
							type="warning"
							@click="handleInitDb"
						>
							初始化数据库
						</el-button>
						<div style="float: right" v-if="register">
							<router-link class="link-type" :to="'/register'"
								>立即注册</router-link
							>
						</div>
					</el-form-item>
				</el-form>
			</div>
		</div>
		<!--  底部  -->
		<div class="el-login-footer">
			<span>evomix管理系统</span>
		</div>
	</div>
</template>

<script setup>
import { getCodeImg } from '@/api/login'
import { initDb } from '@/api/system/init_db'
import Cookies from 'js-cookie'
import { encrypt, decrypt } from '@/utils/jsencrypt'
import useUserStore from '@/store/modules/user'
import SvgIcon from '@/components/SvgIcon/index.vue'

const userStore = useUserStore()
const route = useRoute()
const router = useRouter()
const { proxy } = getCurrentInstance()

const loginForm = ref({
	username: 'admin',
	password: '123456',
	rememberMe: false,
	code: '',
	captchaKey: '',
})

const loginRules = {
	username: [{ required: true, trigger: 'blur', message: '请输入您的账号' }],
	password: [{ required: true, trigger: 'blur', message: '请输入您的密码' }],
	code: [{ required: true, trigger: 'change', message: '请输入验证码' }],
}

const codeUrl = ref('')
const loading = ref(false)
// 验证码开关
const captchaEnabled = ref(true)
// 注册开关
const register = ref(false)
const redirect = ref(undefined)

watch(
	route,
	(newRoute) => {
		redirect.value = newRoute.query && newRoute.query.redirect
	},
	{ immediate: true }
)

function handleLogin() {
	proxy.$refs.loginRef.validate((valid) => {
		if (valid) {
			loading.value = true
			console.log(loginForm.value, 'loginForm.value');
			userStore
				.login(loginForm.value)
				.then(() => {
					const query = route.query
					const otherQueryParams = Object.keys(query).reduce((acc, cur) => {
						if (cur !== 'redirect') {
							acc[cur] = query[cur]
						}
						return acc
					}, {})
					router.push({ path: redirect.value || '/', query: otherQueryParams })
				})
				.catch(() => {
					loading.value = false
				})
		}
	})
}

function getCode() {
	getCodeImg().then((res) => {
		captchaEnabled.value =
			res.captchaEnabled === undefined ? true : res.captchaEnabled
		if (captchaEnabled.value) {
			codeUrl.value = res.captchaSvg
			loginForm.value.captchaKey = res.captchaKey
			console.log(loginForm.value);
		}
	})
}

const handleInitDb = () => {
	proxy.$modal.confirm('确认要初始化数据库吗？').then(() => {
		initDb().then(res => {
			proxy.$modal.msgSuccess('初始化成功')
		})
	}).catch(() => {})
}

getCode()
</script>

<style lang="scss" scoped>
.login {
	display: flex;
	justify-content: center;
	align-items: center;
	height: 100%;
	background-size: cover;
	background-image: linear-gradient(to top, #accbee 0%, #e7f0fd 100%);
}

.login-code {
	width: 33%;
	height: 40px;
	float: right;
	img {
		cursor: pointer;
		vertical-align: middle;
	}
}
.el-login-footer {
	height: 40px;
	line-height: 40px;
	position: fixed;
	bottom: 0;
	width: 100%;
	text-align: center;
	color: #fff;
	font-family: Arial;
	font-size: 12px;
	letter-spacing: 1px;
}
.login-code-img {
	height: 40px;
	padding-left: 12px;
}

.container {
	position: relative;
	height: 100vh;
	width: 100%;
	overflow: hidden;
	display: flex;
	align-items: center;
	justify-content: center;
}

.container .wavy {
	position: absolute;
	left: 0;
	width: 100%;
	height: 100%;
	background: #4973ff;
}
.container .wavy span {
	position: absolute;
	width: 295vh;
	height: 295vh;
	top: 0;
	left: 50%;
	transform: translate(-50%, -75%);
	background: #000000;
}
.container .wavy span:nth-child(1) {
	animation: animate 45s linear infinite;
	border-radius: 45%;
	background: rgba(20, 20, 20, 1);
}
.container .wavy span:nth-child(2) {
	animation: animate 50s linear infinite;
	border-radius: 40%;
	background: rgba(20, 20, 20, 0.5);
}
.container .wavy span:nth-child(3) {
	animation: animate 45s linear infinite;
	border-radius: 42.5%;
	background: rgba(20, 20, 20, 0.5);
}

@keyframes animate {
	0% {
		transform: translate(-50%, -75%) rotate(0deg);
	}
	100% {
		transform: translate(-50%, -75%) rotate(360deg);
	}
}
.content {
	position: relative;
	z-index: 1;
	margin-top: -180px;
	.login-form {
		position: relative;
		border-radius: 6px;
		background: #ffffff;
		width: 400px;
		padding: 25px 25px 5px 25px;
		box-shadow: 0px 0px 10px 2px rgba(255, 255, 255);
		.title {
			margin: 0px auto 30px auto;
			text-align: center;
			color: #707070;
		}
		.el-input {
			height: 40px;
			input {
				height: 40px;
			}
		}
		.input-icon {
			height: 39px;
			width: 16px;
			margin-left: 0px;
		}
	}
}
</style>
