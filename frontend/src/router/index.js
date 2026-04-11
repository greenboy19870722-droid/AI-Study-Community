import { createRouter, createWebHistory } from 'vue-router'
import { getCurrentUserId, isLoggedIn } from '../utils/auth.js'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/post/list'
    },
    {
      path: '/post',
      name: 'post',
      component: import('@/views/Post.vue'),
      children: [
        {
          path: 'list',
          name: 'PostList',
          component: import('@/views/post/PostList.vue'),
          meta: { title: '帖子列表' }
        },
        {
          path: 'detail/:id',
          name: 'PostDetail',
          component: import('@/views/post/PostDetail.vue'),
          meta: { title: '帖子详情', requiresComment: true }
        },
        {
          path: 'create',
          name: 'PostCreate',
          component: import('@/views/post/PostCreate.vue'),
          meta: { title: '发布帖子', requiresAuth: true }
        },
        {
          path: 'edit/:id',
          name: 'PostEdit',
          component: import('@/views/post/PostEdit.vue'),
          meta: { title: '编辑帖子', requiresAuth: true }
        },
      ]
    },
    {
      path: '/login',
      name: 'Login',
      component: import('@/views/Login.vue'),
      meta: { title: '登录', requiresAuth: false }
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

    return next(false)
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

  const currentUserId = getCurrentUserId();

  if (currentUserId !== authorId) {
    alert('只能删除自己的评论')
    return false
  }
  return true
}

export default router
