<template>
	<el-config-provider :locale="zhCn">
		<router-view />
	</el-config-provider>
</template>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import useSettingsStore from '@/store/modules/settings'
import { handleThemeStyle } from '@/utils/theme'

onMounted(() => {
	nextTick(() => {
		// 初始化主题样式
		handleThemeStyle(useSettingsStore().theme)
	})
})
defineOptions({
  name: 'App'
})

</script>
