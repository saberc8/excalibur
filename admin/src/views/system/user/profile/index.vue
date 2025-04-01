<template>
	<div class="app-container">
		<el-row :gutter="20">
			<el-col :span="18" :xs="24">
				<el-card>
					<template v-slot:header>
						<div class="clearfix">
							<span>基本资料</span>
						</div>
					</template>
					<el-tabs v-model="activeTab">
						<el-tab-pane label="基本资料" name="userinfo">
							<userInfo :user="state.user" />
						</el-tab-pane>
						<el-tab-pane label="修改密码" name="resetPwd">
							<resetPwd />
						</el-tab-pane>
					</el-tabs>
				</el-card>
			</el-col>
		</el-row>
	</div>
</template>

<script setup name="Profile">
import userAvatar from './userAvatar'
import userInfo from './userInfo'
import resetPwd from './resetPwd'

const activeTab = ref('userinfo')
const state = reactive({
	user: {},
	roleGroup: {},
	postGroup: {},
})

</script>
