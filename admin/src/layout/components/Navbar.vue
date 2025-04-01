<template>
	<div class="navbar">
		<hamburger
			id="hamburger-container"
			:is-active="appStore.sidebar.opened"
			class="hamburger-container"
			@toggleClick="toggleSideBar"
		/>
		<breadcrumb id="breadcrumb-container" class="breadcrumb-container" />
		<div class="right-menu">
			<template v-if="appStore.device !== 'mobile'">
				<screenfull id="screenfull" class="right-menu-item hover-effect" />
			</template>
			<div class="avatar-container">
				<el-dropdown
					@command="handleCommand"
					class="right-menu-item hover-effect"
					trigger="click"
				>
					<div class="avatar-wrapper">
						<img :src="userStore.avatar" class="user-avatar" />
						 <svg-icon icon-class="solar:alt-arrow-down-bold-duotone" />
					</div>
					<template #dropdown>
						<el-dropdown-menu>
							<router-link to="/user/profile">
								<el-dropdown-item>个人中心</el-dropdown-item>
							</router-link>
							<el-dropdown-item
								command="setLayout"
								v-if="settingsStore.showSettings"
							>
								<span>布局设置</span>
							</el-dropdown-item>
							<el-dropdown-item divided command="logout">
								<span>退出登录</span>
							</el-dropdown-item>
						</el-dropdown-menu>
					</template>
				</el-dropdown>
			</div>
		</div>
	</div>
</template>

<script setup>
import { ElMessageBox } from 'element-plus'
import Breadcrumb from './Breadcrumb'
import Hamburger from './Hamburger'
import Screenfull from './Screenfull'
import SvgIcon from '@/components/SvgIcon/index.vue'
import useAppStore from '@/store/modules/app'
import useUserStore from '@/store/modules/user'
import useSettingsStore from '@/store/modules/settings'

const appStore = useAppStore()
const userStore = useUserStore()
const settingsStore = useSettingsStore()

function toggleSideBar() {
	appStore.toggleSideBar()
}

function handleCommand(command) {
	switch (command) {
		case 'setLayout':
			setLayout()
			break
		case 'logout':
			logout()
			break
		default:
			break
	}
}

function logout() {
	ElMessageBox.confirm('确定注销并退出系统吗？', '提示', {
		confirmButtonText: '确定',
		cancelButtonText: '取消',
		type: 'warning',
	})
		.then(() => {
			userStore.logOut().then(() => {
				location.reload()
			})
		})
		.catch(() => {})
}

const emits = defineEmits(['setLayout'])
function setLayout() {
	emits('setLayout')
}
</script>

<style lang="scss" scoped>
.navbar {
	height: 50px;
	overflow: hidden;
	position: relative;
	background: #fff;
	box-shadow: 0 1px 4px rgba(41, 0, 0, 0.04);

	.hamburger-container {
		line-height: 46px;
		height: 100%;
		float: left;
		cursor: pointer;
		transition: background 0.3s;
		-webkit-tap-highlight-color: transparent;

		&:hover {
			background: rgba(0, 0, 0, 0.025);
		}
	}

	.breadcrumb-container {
		float: left;
	}

	.topmenu-container {
		position: absolute;
		left: 50px;
	}

	.errLog-container {
		display: inline-block;
		vertical-align: top;
	}

	.right-menu {
		float: right;
		height: 100%;
		display: flex;
		align-items: center;

		&:focus {
			outline: none;
		}

		.right-menu-item {
			font-size: 12px;
			color: #5a5e66;
			padding: 4px;
			margin-right: 10px;

			&.hover-effect {
				cursor: pointer;
				transition: background 0.3s;

				&:hover {
					background: rgba(0, 0, 0, 0.1);
					padding: 4px;
					border-radius: 4px;
				}
			}
		}

		.avatar-container {

			.avatar-wrapper {
				position: relative;
				display: flex;
				align-items: center;

				.user-avatar {
					cursor: pointer;
					width: 30px;
					height: 30px;
					border-radius: 6px;
				}
				.svg-icon {
					width: 16px;
					height: 16px;
					margin-left: 6px;
					cursor: pointer;
				}
			}
		}
	}
}
</style>
