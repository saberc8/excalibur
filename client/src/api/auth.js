import request from '@/utils/request'

export function getCaptcha() {
  return request({
    url: '/common/captcha',
    method: 'get'
  })
}

export function login(data) {
  return request({
    url: '/common/login',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/common/logout',
    method: 'post'
  })
}

export function register(data) {
  return request({
    url: '/common/register',
    method: 'post',
    data
  })
}
