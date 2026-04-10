<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getPostDetail } from '@/api/post.js'
import { getCommentTree, createComment, replyComment } from '@/api/comment.js'

// 获取路由信息
const route = useRoute()
const router = useRouter()

// 帖子数据
const post = ref(null)

// 评论数据
const comments = ref([])
const totalComments = ref(0)
const commentPage = ref(1)
const commentPageSize = ref(10)

// 评论加载状态
const commentLoading = ref(false)
const commentError = ref(null)

// 加载状态
const loading = ref(false)
const error = ref(null)

// 评论输入
const commentContent = ref('')
const replyingTo = ref(null) // { id, authorName } 回复目标
const commentSubmitting = ref(false)
const commentSubmitError = ref(null)

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

// 评论总页数
const commentTotalPages = computed(() => {
  return Math.ceil(totalComments.value / commentPageSize.value) || 1
})

// 加载帖子详情
const loadPostDetail = async () => {
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

// 加载评论列表
const loadComments = async () => {
  const postId = route.params.id || route.query.id

  if (!postId) return

  commentLoading.value = true
  commentError.value = null

  try {
    const params = {
      postId,
      page: commentPage.value,
      pageSize: commentPageSize.value,
    }
    const data = await getCommentTree(params)
    comments.value = data.list || []
    totalComments.value = data.total || 0
  } catch (err) {
    commentError.value = err.message || '加载评论失败'
    console.error('加载评论失败:', err)
  } finally {
    commentLoading.value = false
  }
}

// 提交评论
const submitComment = async () => {
  if (!commentContent.value.trim()) {
    commentSubmitError.value = '评论内容不能为空'
    return
  }

  const postId = route.params.id || route.query.id
  if (!postId) return

  commentSubmitting.value = true
  commentSubmitError.value = null

  try {
    if (replyingTo.value) {
      // 回复评论
      await replyComment({
        postId,
        parentId: replyingTo.value.id,
        authorId: 1, // TODO: 从用户状态获取
        content: commentContent.value.trim(),
      })
    } else {
      // 新增评论
      await createComment({
        postId,
        authorId: 1, // TODO: 从用户状态获取
        content: commentContent.value.trim(),
      })
    }

    // 清空输入
    commentContent.value = ''
    replyingTo.value = null

    // 刷新评论列表
    commentPage.value = 1
    await loadComments()
  } catch (err) {
    commentSubmitError.value = err.message || '提交评论失败'
    console.error('提交评论失败:', err)
  } finally {
    commentSubmitting.value = false
  }
}

// 取消回复
const cancelReply = () => {
  replyingTo.value = null
  commentContent.value = ''
  commentSubmitError.value = null
}

// 点击回复按钮
const handleReply = (comment) => {
  replyingTo.value = {
    id: comment.id,
    authorName: comment.authorName || `用户${comment.authorId}`,
  }
  // 滚动到评论输入框
  document.querySelector('.comment-input-area')?.scrollIntoView({ behavior: 'smooth' })
}

// 评论翻页
const handleCommentPageChange = (newPage) => {
  if (newPage < 1 || newPage > commentTotalPages.value) return
  commentPage.value = newPage
  loadComments()
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
  loadComments()
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

    <!-- 评论区 -->
    <div v-if="post && post.id" class="comment-section">
      <!-- 评论头部 -->
      <div class="comment-header">
        <h3 class="comment-title">评论 ({{ totalComments }})</h3>
      </div>

      <!-- 评论输入区 -->
      <div class="comment-input-area">
        <!-- 回复提示 -->
        <div v-if="replyingTo" class="reply-hint">
          <span>回复 @{{ replyingTo.authorName }}：</span>
          <button class="btn-cancel-reply" @click="cancelReply">取消回复</button>
        </div>

        <!-- 输入框 -->
        <div class="comment-input-wrapper">
          <textarea
            v-model="commentContent"
            class="comment-textarea"
            placeholder="写下你的评论..."
            rows="3"
            :disabled="commentSubmitting"
          ></textarea>
        </div>

        <!-- 错误提示 -->
        <div v-if="commentSubmitError" class="comment-error">
          {{ commentSubmitError }}
        </div>

        <!-- 提交按钮 -->
        <div class="comment-actions">
          <button
            class="btn-submit-comment"
            :disabled="commentSubmitting || !commentContent.trim()"
            @click="submitComment"
          >
            {{ commentSubmitting ? '提交中...' : (replyingTo ? '回复' : '发表评论') }}
          </button>
        </div>
      </div>

      <!-- 评论加载状态 -->
      <div v-if="commentLoading" class="comment-loading">
        <div class="loading-spinner small"></div>
        <span>加载评论中...</span>
      </div>

      <!-- 评论错误状态 -->
      <div v-else-if="commentError" class="comment-error-state">
        <p class="error-message">{{ commentError }}</p>
        <button class="btn-retry small" @click="loadComments">重试</button>
      </div>

      <!-- 评论列表 -->
      <div v-else-if="comments.length > 0" class="comment-list">
        <div
          v-for="comment in comments"
          :key="comment.id"
          class="comment-item"
        >
          <div class="comment-main">
            <div class="comment-author">
              <span class="author-name">{{ comment.authorName || `用户${comment.authorId}` }}</span>
              <span class="comment-time">{{ formatDate(comment.createdAt) }}</span>
            </div>
            <div class="comment-body">{{ comment.content }}</div>
            <div class="comment-actions-row">
              <button class="btn-reply" @click="handleReply(comment)">回复</button>
            </div>
          </div>

          <!-- 子评论 -->
          <div v-if="comment.children && comment.children.length > 0" class="comment-children">
            <div
              v-for="child in comment.children"
              :key="child.id"
              class="comment-child"
            >
              <div class="child-author">
                <span class="author-name">{{ child.authorName || `用户${child.authorId}` }}</span>
                <span v-if="child.replyToAuthor" class="reply-arrow">回复</span>
                <span v-if="child.replyToAuthor" class="reply-target">@{{ child.replyToAuthor }}</span>
                <span class="comment-time">{{ formatDate(child.createdAt) }}</span>
              </div>
              <div class="comment-body">{{ child.content }}</div>
              <div class="comment-actions-row">
                <button class="btn-reply" @click="handleReply(child)">回复</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 评论空状态 -->
      <div v-else-if="!commentLoading" class="comment-empty">
        <p>暂无评论，来发表第一篇评论吧</p>
      </div>

      <!-- 评论分页 -->
      <div v-if="!commentLoading && comments.length > 0" class="comment-pagination">
        <button
          class="page-btn"
          :disabled="commentPage <= 1"
          @click="handleCommentPageChange(commentPage - 1)"
        >
          上一页
        </button>
        <span class="page-info">第 {{ commentPage }} / {{ commentTotalPages }} 页</span>
        <button
          class="page-btn"
          :disabled="commentPage >= commentTotalPages"
          @click="handleCommentPageChange(commentPage + 1)"
        >
          下一页
        </button>
      </div>
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

.loading-spinner.small {
  width: 24px;
  height: 24px;
  border-width: 2px;
  margin: 0 auto;
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

.btn-retry.small {
  padding: 6px 16px;
  font-size: 13px;
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

/* 评论区样式 */
.comment-section {
  margin-top: 32px;
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 24px;
}

.comment-header {
  margin-bottom: 20px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e5e7eb;
}

.comment-title {
  font-size: 18px;
  font-weight: 600;
  color: #1f2937;
  margin: 0;
}

.comment-input-area {
  margin-bottom: 24px;
}

.reply-hint {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
  padding: 8px 12px;
  background-color: #f3f4f6;
  border-radius: 6px;
  font-size: 14px;
  color: #6b7280;
}

.btn-cancel-reply {
  padding: 4px 12px;
  background-color: #ffffff;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 12px;
  color: #6b7280;
  cursor: pointer;
}

.btn-cancel-reply:hover {
  background-color: #fee2e2;
  border-color: #ef4444;
  color: #ef4444;
}

.comment-input-wrapper {
  margin-bottom: 12px;
}

.comment-textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 14px;
  line-height: 1.6;
  resize: vertical;
  font-family: inherit;
  box-sizing: border-box;
}

.comment-textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.comment-textarea:disabled {
  background-color: #f9fafb;
  cursor: not-allowed;
}

.comment-error {
  color: #ef4444;
  font-size: 13px;
  margin-bottom: 12px;
}

.comment-actions {
  display: flex;
  justify-content: flex-end;
}

.btn-submit-comment {
  padding: 10px 24px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-submit-comment:hover:not(:disabled) {
  background-color: #2563eb;
}

.btn-submit-comment:disabled {
  background-color: #9ca3af;
  cursor: not-allowed;
}

.comment-loading,
.comment-error-state {
  text-align: center;
  padding: 24px;
  color: #6b7280;
}

.comment-error-state .error-message {
  margin-bottom: 12px;
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.comment-item {
  padding-bottom: 20px;
  border-bottom: 1px solid #f3f4f6;
}

.comment-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.comment-main {
  margin-bottom: 12px;
}

.comment-author {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.author-name {
  font-weight: 600;
  color: #1f2937;
  font-size: 14px;
}

.comment-time {
  color: #9ca3af;
  font-size: 12px;
}

.comment-body {
  color: #374151;
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 8px;
  white-space: pre-wrap;
  word-break: break-word;
}

.comment-actions-row {
  display: flex;
  gap: 12px;
}

.btn-reply {
  padding: 4px 12px;
  background-color: transparent;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 12px;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-reply:hover {
  background-color: #eff6ff;
  border-color: #3b82f6;
  color: #3b82f6;
}

/* 子评论 */
.comment-children {
  margin-left: 24px;
  padding-left: 16px;
  border-left: 2px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.comment-child {
  padding: 12px;
  background-color: #f9fafb;
  border-radius: 6px;
}

.child-author {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
  font-size: 13px;
}

.reply-arrow {
  color: #9ca3af;
  font-size: 12px;
}

.reply-target {
  color: #3b82f6;
  font-size: 12px;
}

.comment-empty {
  text-align: center;
  padding: 32px;
  color: #9ca3af;
  font-size: 14px;
}

/* 评论分页 */
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
</style>
