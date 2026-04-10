<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getPostDetail } from '@/api/post.js'

// 获取路由信息
const route = useRoute()
const router = useRouter()

// 帖子数据
const post = ref(null)

// 加载状态
const loading = ref(false)
const error = ref(null)

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  return `${year}-${month}-${day} ${hours}:${minutes}`
}

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = {
    0: '已删除',
    1: '正常',
    2: '置顶',
    3: '推荐',
  }
  return statusMap[status] || '未知'
}

// 获取状态类名
const getStatusClass = (status) => {
  const classMap = {
    0: 'deleted',
    1: 'normal',
    2: 'top',
    3: 'recommend',
  }
  return classMap[status] || ''
}

// 加载帖子详情
const loadPostDetail = async () => {
  // 从路由参数获取帖子ID
  const postId = route.params.id || route.query.id

  if (!postId) {
    error.value = '帖子ID不存在'
    return
  }

  loading.value = true
  error.value = null

  try {
    const data = await getPostDetail(postId)
    post.value = data
  } catch (err) {
    error.value = err.message || '加载帖子详情失败'
    console.error('加载帖子详情失败:', err)
  } finally {
    loading.value = false
  }
}

// 返回列表页
const goBack = () => {
  router.push('/post/list')
}

// 跳转到编辑页
const goToEdit = () => {
  router.push(`/post/edit/${post.value.id}`)
}

// 组件挂载时加载数据
onMounted(() => {
  loadPostDetail()
})
</script>

<template>
  <div class="post-detail-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <button class="btn-back" @click="goBack">
        &larr; 返回列表
      </button>
      <button
        v-if="post && post.id"
        class="btn-edit"
        @click="goToEdit"
      >
        编辑帖子
      </button>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <span>加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="error-state">
      <p class="error-message">{{ error }}</p>
      <div class="error-actions">
        <button class="btn-retry" @click="loadPostDetail">重试</button>
        <button class="btn-back-error" @click="goBack">返回列表</button>
      </div>
    </div>

    <!-- 帖子详情 -->
    <div v-else-if="post && post.id" class="post-detail">
      <!-- 帖子头部 -->
      <div class="post-header">
        <div class="post-badges">
          <span
            v-if="post.status !== 1"
            class="post-badge"
            :class="getStatusClass(post.status)"
          >
            {{ getStatusText(post.status) }}
          </span>
        </div>
        <h1 class="post-title">{{ post.title }}</h1>
        <div class="post-meta">
          <span class="meta-item">
            <span class="meta-label">作者ID:</span>
            <span class="meta-value">{{ post.authorId }}</span>
          </span>
          <span class="meta-item">
            <span class="meta-label">发布时间:</span>
            <span class="meta-value">{{ formatDate(post.createdAt) }}</span>
          </span>
          <span v-if="post.updatedAt && post.updatedAt !== post.createdAt" class="meta-item">
            <span class="meta-label">更新时间:</span>
            <span class="meta-value">{{ formatDate(post.updatedAt) }}</span>
          </span>
        </div>
      </div>

      <!-- 标签 -->
      <div v-if="post.tags" class="post-tags">
        <span
          v-for="(tag, index) in post.tags.split(',')"
          :key="index"
          class="tag-item"
        >
          {{ tag.trim() }}
        </span>
      </div>

      <!-- 封面图片 -->
      <div v-if="post.coverImage" class="post-cover">
        <img :src="post.coverImage" :alt="post.title" />
      </div>

      <!-- 帖子内容 -->
      <div class="post-content">
        <p>{{ post.content }}</p>
      </div>

      <!-- 帖子底部 -->
      <div class="post-footer">
        <button class="btn-like">
          ♥ 点赞
        </button>
        <button class="btn-share">
          分享
        </button>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-else class="empty-state">
      <p>帖子不存在或已被删除</p>
      <button class="btn-primary" @click="goBack">返回列表</button>
    </div>
  </div>
</template>

<style scoped>
.post-detail-container {
  max-width: 900px;
  margin: 0 auto;
  padding: 20px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e5e7eb;
}

.btn-back {
  padding: 8px 16px;
  background-color: #6b7280;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-back:hover {
  background-color: #4b5563;
}

.btn-edit {
  padding: 8px 16px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-edit:hover {
  background-color: #2563eb;
}

.loading-state,
.error-state,
.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #e5e7eb;
  border-top-color: #3b82f6;
  border-radius: 50%;
  margin: 0 auto 16px;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.error-message {
  color: #ef4444;
  margin-bottom: 16px;
}

.error-actions {
  display: flex;
  gap: 12px;
  justify-content: center;
}

.btn-retry {
  padding: 8px 20px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.btn-back-error {
  padding: 8px 20px;
  background-color: #6b7280;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.btn-primary {
  padding: 8px 20px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.post-detail {
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 32px;
}

.post-header {
  margin-bottom: 24px;
}

.post-badges {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
}

.post-badge {
  padding: 4px 12px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.post-badge.deleted {
  background-color: #fee2e2;
  color: #dc2626;
}

.post-badge.normal {
  background-color: #d1fae5;
  color: #059669;
}

.post-badge.top {
  background-color: #fef3c7;
  color: #d97706;
}

.post-badge.recommend {
  background-color: #dbeafe;
  color: #2563eb;
}

.post-title {
  font-size: 28px;
  font-weight: 700;
  color: #1f2937;
  margin: 0 0 16px;
  line-height: 1.4;
}

.post-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  font-size: 14px;
  color: #6b7280;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.meta-label {
  color: #9ca3af;
}

.post-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
  padding-bottom: 24px;
  border-bottom: 1px solid #e5e7eb;
}

.tag-item {
  padding: 4px 12px;
  background-color: #eff6ff;
  color: #3b82f6;
  border-radius: 16px;
  font-size: 13px;
}

.post-cover {
  margin-bottom: 24px;
  border-radius: 8px;
  overflow: hidden;
}

.post-cover img {
  width: 100%;
  max-height: 400px;
  object-fit: cover;
}

.post-content {
  color: #374151;
  font-size: 16px;
  line-height: 1.8;
  margin-bottom: 32px;
}

.post-content p {
  margin: 0 0 16px;
  white-space: pre-wrap;
  word-break: break-word;
}

.post-content p:last-child {
  margin-bottom: 0;
}

.post-footer {
  display: flex;
  gap: 12px;
  padding-top: 24px;
  border-top: 1px solid #e5e7eb;
}

.btn-like,
.btn-share {
  padding: 8px 20px;
  border: 1px solid #d1d5db;
  background-color: #ffffff;
  color: #374151;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-like:hover,
.btn-share:hover {
  background-color: #f3f4f6;
  border-color: #3b82f6;
  color: #3b82f6;
}
</style>
