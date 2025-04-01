import { createApp } from 'vue'

import Cookies from 'js-cookie'

// import ElementPlus from 'element-plus'
// import zhCn from 'element-plus/dist/locale/zh-cn.mjs'

import '@/assets/styles/index.scss' // global css
import '@/assets/styles/variables.module.scss' // global css
import 'element-plus/theme-chalk/index.css' // global css

import App from './App'
import store from './store'
import router from './router'
// 注册指令
import plugins from './plugins' // plugins
import { download } from '@/utils/request'

import './permission' // permission control

import {
	parseTime,
	resetForm,
	addDateRange,
	handleTree
} from '@/utils/tools'

import VxeTable from 'vxe-table'
import 'vxe-table/lib/style.css'

import VxeUI from 'vxe-pc-ui'
import 'vxe-pc-ui/lib/style.css'

import * as echarts from 'echarts'

const app = createApp(App)
app.config.globalProperties.$echarts = echarts // 全局使用echarts
app.config.globalProperties.download = download
app.config.globalProperties.parseTime = parseTime
app.config.globalProperties.resetForm = resetForm
app.config.globalProperties.handleTree = handleTree
app.config.globalProperties.addDateRange = addDateRange
app.use(router)
app.use(store)
app.use(plugins)

app.use(VxeUI)
app.use(VxeTable)

app.mount('#app')
