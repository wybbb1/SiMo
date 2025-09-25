import { createRouter, createWebHistory } from 'vue-router'
import login from '@/views/login.vue'
import home from '@/views/home.vue'
import write from '@/views/write.vue'
import whiteBoard from '@/views/whiteBoard.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
      path: '/home',
      name: 'home',
      component: home
    },
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/write',
      name: 'write',
      component: write
    },
    {
      path: '/whiteBoard',
      name: 'whiteBoard',
      component: whiteBoard

    }

  ],
})

export default router
