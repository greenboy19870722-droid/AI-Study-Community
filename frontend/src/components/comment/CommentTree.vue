<!--
  CommentTree.vue - 评论树组件
  显示帖子的所有评论（树形结构）
-->
<script setup>
import { ref, computed, onMounted } from 'vue'
import { getCommentTree } from '@/api/comment.js'
import CommentItem from './CommentItem.vue'

const props = defineProps({
  postId: {
    type: [Number, String],
    required: true
  },
  pageSize: {
    type: Number,
    default: 20
  }
})

const emit = defineEmits(['reply', 'delete', 'load-more'])

const comments = ref([])
const total = ref(0)
const loading = ref(false)
const error = ref(null)
const currentPage = ref(1)

// 计算总页数
const totalPages = computed(() => {
  return Math.ceil(total.value / props.pageSize) || 1
})

// 加载评论树
const loadComments = async () => {
  loading.value = true
  error.value = null
  
  try {
    const data = await getCommentTree({
      postId: props.postId,
      page: currentPage.value,
      pageSize: props.pageSize
    })
    comments.value = data.list || []
    total.value = data.total || 0
  } catch (err) {
    error.value = err.message || '加载评论失败'
    console.error('加载评论失败:', err)
  } finally {
    loading.value = false
  }
}

// 处理回复事件
const handleReply = (comment) => {
  emit('reply', comment)
}

// 处理删除事件
const handleDelete = (comment) => {
  emit('delete', comment)
}

// 翻页
const changePage = (newPage) => {
  if (newPage < 1 || newPage > totalPages.value) return
  currentPage.value = newPage
  loadComments()
}

// 刷新评论列表
const refresh = () => {
  currentPage.value = 1
  loadComments()
}

// 暴露方法给父组件
defineExpose({
  refresh,
  loadComments
})

// 组件挂载时加载数据
onMounted(() => {
  loadComments()
})
</script>

<template>
  <div class="comment-tree">
    <!-- 加载状态 -->
    <div v-if="loading && comments.length === 0" class="comment-loading">
      <div class="loading-spinner"></div>
      <span>加载评论中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error" class="comment-error-state">
      <p class="error-message">{{ error }}</p>
      <button class="btn-retry" @click="loadComments">重试</button>
    </div>

    <!-- 评论列表 -->
    <div v-else-if="comments.length > 0" class="comment-list">
      <CommentItem
        v-for="comment in comments"
        :key="comment.id"
        :comment="comment"
        @reply="handleReply"
        @delete="handleDelete"
      />
    </div>

    <!-- 空状态 -->
    <div v-else class="comment-empty">
      <p>暂无评论，来发表第一篇评论吧</p>
    </div>

    <!-- 分页 -->
    <div v-if="!loading && comments.length > 0" class="comment-pagination">
      <button
        class="page-btn"
        :disabled="currentPage <= 1"
        @click="changePage(currentPage - 1)"
      >
        上一页
      </button>
      <span class="page-info">第 {{ currentPage }} / {{ totalPages }} 页</span>
      <button
        class="page-btn"
        :disabled="currentPage >= totalPages"
        @click="changePage(currentPage + 1)"
      >
        下一页
      </button>
    </div>
  </div>
</template>

<style scoped>
.comment-tree {
  width: 100%;
}

.comment-loading,
.comment-error-state {
  text-align: center;
  padding: 32px;
  color: #6b7280;
}

.loading-spinner {
  width: 32px;
  height: 32px;
  border: 3px solid #e5e7eb;
  border-top-color: #3b82f6;
  border-radius: 50%;
  margin: 0 auto 12px;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.error-message {
  color: #ef4444;
  margin-bottom: 12px;
}

.btn-retry {
  padding: 8px 20px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
}

.btn-retry:hover {
  background-color: #2563eb;
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.comment-empty {
  text-align: center;
  padding: 40px 20px;
  color: #9ca3af;
  font-size: 14px;
}

.comment-pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  margin-top: 24px;
  padding-top: 16px;
  border-top: 1px solid #e5e7eb;
}

.page-info {
  font-size: 14px;
  color: #6b7280;
}

.page-btn {
  padding: 6px 16px;
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
</style>
