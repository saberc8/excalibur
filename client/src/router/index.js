import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'home',
    component: () => import('@/views/home/index.vue')
  },
  {
    path: '/page2',
    name: 'page2',
    component: () => import('@/views/page2/index.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
