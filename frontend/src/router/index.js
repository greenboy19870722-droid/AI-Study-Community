import { createRouter, createWebHistory } from 'vue-router'
import PostList from '../views/post/PostList.vue'
import PostDetail from '../views/post/PostDetail.vue'
import PostCreate from '../views/post/PostCreate.vue'
import PostEdit from '../views/post/PostEdit.vue'
import { isLoggedIn } from '../utils/auth.js'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/post/list'
    },
    {
      path: '/post/list',
      name: 'PostList',
      component: PostList,
      meta: { title: '帖子列表' }
    },
    {
      path: '/post/detail/:id',
      name: 'PostDetail',
      component: PostDetail,
      meta: { title: '帖子详情', requiresComment: true }
    },
    {
      path: '/post/create',
      name: 'PostCreate',
      component: PostCreate,
      meta: { title: '发布帖子', requiresAuth: true }
    },
    {
      path: '/post/edit/:id',
      name: 'PostEdit',
      component: PostEdit,
      meta: { title: '编辑帖子', requiresAuth: true }
    }
  ]
})

// 路由守卫 - 全局前置守卫
router.beforeEach((to, from, next) => {
  // 更新页面标题
  if (to.meta.title) {
    document.title = to.meta.title
  }

  // 检查需要登录的路由
  if (to.meta.requiresAuth && !isLoggedIn()) {
    // TODO: 跳转到登录页（登录页待实现）
    // next({ name: 'Login', query: { redirect: to.fullPath } })
    console.warn('[Route Guard] 此操作需要登录')
    alert('请先登录后再进行操作')
    // 暂时允许继续，但评论功能会检查登录状态
    next()
    return
  }

  next()
})

// 路由守卫 - 评论权限检查（导出供组件使用）
export function checkCommentPermission() {
  if (!isLoggedIn()) {
    alert('请先登录后再评论')
    return false
  }
  return true
}

// 路由守卫 - 评论删除权限检查
export function checkCommentDeletePermission(authorId) {
  if (!isLoggedIn()) {
    alert('请先登录后再操作')
    return false
  }
  const currentUserId = parseInt(localStorage.getItem('auth_user') ? JSON.parse(localStorage.getItem('auth_user')).id : 0)
  if (currentUserId !== authorId) {
    alert('只能删除自己的评论')
    return false
  }
  return true
}

export default router
