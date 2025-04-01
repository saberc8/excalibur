import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '@/layouts/default-layout.vue'

const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login/index.vue'),
  },
  {
    path: '/:spacePath',
    component: DefaultLayout,
    children: [
      {
        path: 'home',
        name: 'home',
        component: () => import('@/views/home/index.vue'),
      },
      {
        path: 'page1',
        name: 'page1',
        component: () => import('@/views/page1/index.vue'),
      },
      {
        path: 'page2',
        name: 'page2',
        component: () => import('@/views/page2/index.vue'),
      },
      {
        path: 'projects',
        name: 'projects',
        component: () => import('@/views/projects/index.vue'),
      },
      {
        path: 'tasks',
        name: 'tasks',
        component: () => import('@/views/tasks/index.vue'),
      },
      {
        path: 'project/:id',
        name: 'project-detail',
        component: () => import('@/views/projects/detail.vue'),
      },
    ],
  },
  {
    path: '/',
    redirect: to => {
      // 从 localStorage 获取最后访问的空间路径，如果没有则默认使用 'avalon'
      const lastSpacePath = localStorage.getItem('lastSpacePath') || 'avalon'
      return `/${lastSpacePath}/home`
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
