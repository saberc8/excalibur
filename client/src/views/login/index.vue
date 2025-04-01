<template>
  <div class="login-container">
    <div class="login-box">
      <a-typography-title :heading="2" style="margin-bottom: 40px">
        登录系统
      </a-typography-title>
      <a-form
        :model="loginForm"
        @submit="handleSubmit"
        style="width: 100%"
      >
        <a-form-item field="username" label="用户名">
          <a-input v-model="loginForm.username" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item field="password" label="密码">
          <a-input-password v-model="loginForm.password" placeholder="请输入密码" />
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" long :loading="loading">
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'
import { login } from '@/api/auth'
import { setToken, setUserInfo } from '@/utils/auth'

const router = useRouter()
const loading = ref(false)
const loginForm = reactive({
  username: '',
  password: ''
})

const handleSubmit = async () => {
  loading.value = true
  try {
    const res = await login(loginForm)
    setToken(res.token)
    setUserInfo(res.userInfo)
    Message.success('登录成功')
    router.push('/')
  } catch (error) {
    Message.error('登录失败：' + (error.message || '未知错误'))
  } finally {
    loading.value = false
  }
}
</script>

<style lang="less" scoped>
.login-container {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-neutral-1);
}

.login-box {
  width: 360px;
  padding: 40px;
  border-radius: 4px;
  background-color: var(--color-bg-2);
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}
</style>
