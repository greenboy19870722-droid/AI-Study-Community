<script setup>
import { ref, onMounted, computed } from 'vue'
import { getPostList } from '@/api/post.js'

// 帖子列表数据
const postList = ref([])
const total = ref(0)

// 分页参数
const pagination = ref({
  page: 1,
  pageSize: 10,
})

// 加载状态
const loading = ref(false)
const error = ref(null)

// 计算总页数
const totalPages = computed(() => {
  return Math.ceil(total.value / pagination.value.pageSize) || 1
})

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

// 加载帖子列表
const loadPostList = async () => {
  loading.value = true
  error.value = null
  try {
    const params = {
      page: pagination.value.page,
      pageSize: pagination.value.pageSize,
    }
    const data = await getPostList(params)
    postList.value = data.list || []
    total.value = data.total || 0
  } catch (err) {
    error.value = err.message || '加载帖子列表失败'
    console.error('加载帖子列表失败:', err)
  } finally {
    loading.value = false
  }
}

// 翻页
const handlePageChange = (newPage) => {
  if (newPage < 1 || newPage > totalPages.value) return
  pagination.value.page = newPage
  loadPostList()
}

// 每页数量改变
const handlePageSizeChange = (newSize) => {
  pagination.value.pageSize = newSize
  pagination.value.page = 1
  loadPostList()
}

// 跳转到帖子详情
const goToDetail = (postId) => {
  window.location.href = `/post/detail/${postId}`
}

// 跳转到发布帖子页
const goToCreate = () => {
  window.location.href = '/post/create'
}

// 组件挂载时加载数据
onMounted(() => {
  loadPostList()
})
</script>

<template>
  <div class="post-list-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">帖子列表</h1>
      <button class="btn-primary" @click="goToCreate">发布帖子</button>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <span>加载中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="error-state">
      <p class="error-message">{{ error }}</p>
      <button class="btn-retry" @click="loadPostList">重试</button>
    </div>

    <!-- 空状态 -->
    <div v-else-if="postList.length === 0" class="empty-state">
      <p>暂无帖子</p>
      <button class="btn-primary" @click="goToCreate">发布第一个帖子</button>
    </div>

    <!-- 帖子列表 -->
    <div v-else class="post-list">
      <div
        v-for="post in postList"
        :key="post.id"
        class="post-item"
        @click="goToDetail(post.id)"
      >
        <div class="post-header">
          <h2 class="post-title">{{ post.title }}</h2>
          <span v-if="post.status === 2" class="post-badge top">置顶</span>
          <span v-if="post.status === 3" class="post-badge recommend">推荐</span>
        </div>
        <p class="post-content">{{ post.content }}</p>
        <div class="post-meta">
          <span class="meta-item">
            <span class="meta-label">作者ID:</span>
            <span class="meta-value">{{ post.authorId }}</span>
          </span>
          <span class="meta-item">
            <span class="meta-label">发布时间:</span>
            <span class="meta-value">{{ formatDate(post.createdAt) }}</span>
          </span>
          <span v-if="post.tags" class="meta-item">
            <span class="meta-label">标签:</span>
            <span class="meta-value tags">{{ post.tags }}</span>
          </span>
        </div>
      </div>
    </div>

    <!-- 分页组件 -->
    <div v-if="!loading && !error && postList.length > 0" class="pagination">
      <div class="pagination-info">
        共 {{ total }} 条记录，每页
        <select
          v-model.number="pagination.pageSize"
          class="page-size-select"
          @change="handlePageSizeChange(pagination.pageSize)"
        >
          <option :value="5">5</option>
          <option :value="10">10</option>
          <option :value="20">20</option>
          <option :value="50">50</option>
        </select>
        条，当前第 {{ pagination.page }} / {{ totalPages }} 页
      </div>
      <div class="pagination-controls">
        <button
          class="page-btn"
          :disabled="pagination.page <= 1"
          @click="handlePageChange(1)"
        >
          首页
        </button>
        <button
          class="page-btn"
          :disabled="pagination.page <= 1"
          @click="handlePageChange(pagination.page - 1)"
        >
          上一页
        </button>
        <button
          v-for="p in Math.min(5, totalPages)"
          :key="p"
          class="page-btn"
          :class="{ active: p === pagination.page }"
          @click="handlePageChange(p)"
        >
          {{ p }}
        </button>
        <button
          class="page-btn"
          :disabled="pagination.page >= totalPages"
          @click="handlePageChange(pagination.page + 1)"
        >
          下一页
        </button>
        <button
          class="page-btn"
          :disabled="pagination.page >= totalPages"
          @click="handlePageChange(totalPages)"
        >
          尾页
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.post-list-container {
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

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #1f2937;
  margin: 0;
}

.btn-primary {
  padding: 8px 16px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-primary:hover {
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

.btn-retry {
  padding: 8px 20px;
  background-color: #6b7280;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.post-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.post-item {
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 20px;
  cursor: pointer;
  transition: box-shadow 0.2s, border-color 0.2s;
}

.post-item:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: #3b82f6;
}

.post-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 12px;
}

.post-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  margin: 0;
  flex: 1;
}

.post-badge {
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
}

.post-badge.top {
  background-color: #fef3c7;
  color: #d97706;
}

.post-badge.recommend {
  background-color: #dbeafe;
  color: #2563eb;
}

.post-content {
  color: #4b5563;
  font-size: 14px;
  line-height: 1.6;
  margin: 0 0 12px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.post-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
  font-size: 13px;
  color: #6b7280;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.meta-label {
  color: #9ca3af;
}

.meta-value.tags {
  color: #3b82f6;
}

.pagination {
  margin-top: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.pagination-info {
  font-size: 14px;
  color: #6b7280;
}

.page-size-select {
  padding: 4px 8px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
}

.pagination-controls {
  display: flex;
  gap: 8px;
}

.page-btn {
  padding: 6px 12px;
  border: 1px solid #d1d5db;
  background-color: #ffffff;
  color: #374151;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.page-btn:hover:not(:disabled) {
  background-color: #f3f4f6;
  border-color: #3b82f6;
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-btn.active {
  background-color: #3b82f6;
  border-color: #3b82f6;
  color: #ffffff;
}
</style>
