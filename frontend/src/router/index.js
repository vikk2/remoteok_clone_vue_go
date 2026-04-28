import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import LoginPage from '../views/LoginPage.vue'
import PostJob from '../views/PostJob.vue'

const routes = [
  { path: '/', name: 'home', component: Home },
  { path: '/login', name: 'login', component: LoginPage },
  {
    path: '/post-job',
    name: 'post-job',
    component: PostJob,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
