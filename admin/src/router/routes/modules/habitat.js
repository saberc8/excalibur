import { LAYOUT } from '@/router/constant'

const qun = {
  path: '/qun',
  name: 'qun',
  component: LAYOUT,
  redirect: '/qun/list',
  meta: {
    orderNo: 10,
    icon: 'ant-design:cluster-outlined',
    title: '微信管理',
  },
  children: [
    {
      path: 'qunlist',
      name: 'qunList',
      component: () => import('@/views/qun/qun/index.vue'),
      meta: {
        title: '管理',
      },
    },
    {
      path: 'bannerlist',
      name: 'BannerList',
      component: () => import('@/views/qun/banner/index.vue'),
      meta: {
        title: 'banner管理',
      },
    },
  ],
}

export default qun
