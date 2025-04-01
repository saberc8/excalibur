// api/init/db
import request from '@/utils/request'

// 初始化数据库
export function initDb() {
  return request({
    url: '/init/db',
    method: 'post',
  })
}
