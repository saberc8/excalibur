import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
import axios from 'axios'

const app = createApp(App)
app.config.globalProperties.$axios = axios
app.use(ArcoVue)
app.use(router)
app.mount('#app')
