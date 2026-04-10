import { createRouter, createWebHistory } from 'vue-router'
import PostList from '../views/post/PostList.vue'
import PostDetail from '../views/post/PostDetail.vue'
import PostCreate from '../views/post/PostCreate.vue'
import PostEdit from '../views/post/PostEdit.vue'

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
      meta: { title: '帖子详情' }
    },
    {
      path: '/post/create',
      name: 'PostCreate',
      component: PostCreate,
      meta: { title: '发布帖子' }
    },
    {
      path: '/post/edit/:id',
      name: 'PostEdit',
      component: PostEdit,
      meta: { title: '编辑帖子' }
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

export default router
