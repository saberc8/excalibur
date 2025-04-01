import request from '@/utils/request'

export const getWxuserAdminList = () => {
  return request({
    url: '/wxuser/admin/list',
    method: 'get',
  })
}
