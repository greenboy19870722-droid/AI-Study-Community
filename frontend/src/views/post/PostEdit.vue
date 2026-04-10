<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { getPostDetail, updatePost } from '@/api/post.js'

// 获取路由信息
const route = useRoute()
const router = useRouter()

// 帖子ID
const postId = ref(null)

// 表单数据
const form = ref({
  title: '',
  content: '',
  authorId: 1,
  tags: '',
  coverImage: '',
})

// 原始数据（用于检测变更）
const originalData = ref(null)

// 加载状态
const loading = ref(false)
const submitting = ref(false)
const error = ref(null)
const validationErrors = ref({})

// 页面初始化
onMounted(() => {
  postId.value = route.params.id || route.query.id

  if (!postId.value) {
    error.value = '帖子ID不存在'
    return
  }

  loadPostData()
})

// 加载帖子数据
const loadPostData = async () => {
  loading.value = true
  error.value = null

  try {
    const data = await getPostDetail(postId.value)

    if (!data || !data.id) {
      error.value = '帖子不存在或已被删除'
      return
    }

    // 填充表单数据
    form.value = {
      title: data.title || '',
      content: data.content || '',
      authorId: data.authorId || 1,
      tags: data.tags || '',
      coverImage: data.coverImage || '',
    }

    // 保存原始数据用于变更检测
    originalData.value = { ...form.value }

    // 设置页面标题
    document.title = `编辑帖子 - ${data.title}`
  } catch (err) {
    error.value = err.message || '加载帖子数据失败'
    console.error('加载帖子数据失败:', err)
  } finally {
    loading.value = false
  }
}

// 验证表单
const validateForm = () => {
  validationErrors.value = {}

  if (!form.value.title || form.value.title.trim().length < 3) {
    validationErrors.value.title = '标题长度在 3 到 100 个字符'
  }

  if (!form.value.content || form.value.content.trim().length < 10) {
    validationErrors.value.content = '内容至少 10 个字符'
  }

  if (!form.value.authorId) {
    validationErrors.value.authorId = '请输入作者ID'
  }

  return Object.keys(validationErrors.value).length === 0
}

// 检测表单是否有变更
const hasChanges = () => {
  if (!originalData.value) return false

  return (
    form.value.title !== originalData.value.title ||
    form.value.content !== originalData.value.content ||
    form.value.tags !== originalData.value.tags ||
    form.value.coverImage !== originalData.value.coverImage
  )
}

// 提交表单
const handleSubmit = async () => {
  // 清除之前的错误
  error.value = null

  // 验证表单
  if (!validateForm()) {
    return
  }

  // 检查是否有变更
  if (!hasChanges()) {
    error.value = '没有检测到任何变更，请修改后再提交'
    return
  }

  submitting.value = true

  try {
    const submitData = {
      id: Number(postId.value),
      title: form.value.title.trim(),
      content: form.value.content.trim(),
      authorId: Number(form.value.authorId),
      tags: form.value.tags.trim() || undefined,
      coverImage: form.value.coverImage.trim() || undefined,
    }

    const result = await updatePost(submitData)

    // 更新成功，跳转到详情页
    if (result && result.id) {
      router.push(`/post/detail/${result.id}`)
    } else {
      // 如果没有返回id，也跳转到详情页
      router.push(`/post/detail/${postId.value}`)
    }
  } catch (err) {
    error.value = err.message || '更新帖子失败，请稍后重试'
    console.error('更新帖子失败:', err)
  } finally {
    submitting.value = false
  }
}

// 重置表单（恢复到原始数据）
const handleReset = () => {
  if (originalData.value) {
    form.value = { ...originalData.value }
  }
  validationErrors.value = {}
  error.value = null
}

// 返回详情页
const goBack = () => {
  router.push(`/post/detail/${postId.value}`)
}

// 返回列表页
const goToList = () => {
  router.push('/post/list')
}
</script>

<template>
  <div class="post-edit-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <button class="btn-back" @click="goBack">
        &larr; 返回详情
      </button>
      <h1 class="page-title">编辑帖子</h1>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <span>加载帖子数据中...</span>
    </div>

    <!-- 错误状态 -->
    <div v-else-if="error && !form.title" class="error-state">
      <p class="error-message">{{ error }}</p>
      <div class="error-actions">
        <button class="btn-retry" @click="loadPostData">重试</button>
        <button class="btn-back-error" @click="goToList">返回列表</button>
      </div>
    </div>

    <!-- 编辑表单 -->
    <div v-else class="post-form">
      <!-- 错误提示 -->
      <div v-if="error" class="error-banner">
        <span class="error-icon">!</span>
        <span class="error-text">{{ error }}</span>
        <button class="error-close" @click="error = null">&times;</button>
      </div>

      <!-- 标题 -->
      <div class="form-group">
        <label class="form-label">
          帖子标题 <span class="required">*</span>
        </label>
        <input
          v-model="form.title"
          type="text"
          class="form-input"
          :class="{ 'input-error': validationErrors.title }"
          placeholder="请输入帖子标题（3-100个字符）"
          maxlength="100"
          @blur="validateForm"
        />
        <span v-if="validationErrors.title" class="field-error">
          {{ validationErrors.title }}
        </span>
        <span class="char-count">{{ form.title.length }}/100</span>
      </div>

      <!-- 内容 -->
      <div class="form-group">
        <label class="form-label">
          帖子内容 <span class="required">*</span>
        </label>
        <textarea
          v-model="form.content"
          class="form-textarea"
          :class="{ 'input-error': validationErrors.content }"
          placeholder="请输入帖子内容（至少10个字符）"
          rows="10"
          @blur="validateForm"
        ></textarea>
        <span v-if="validationErrors.content" class="field-error">
          {{ validationErrors.content }}
        </span>
        <span class="char-count">{{ form.content.length }} 字符</span>
      </div>

      <!-- 标签 -->
      <div class="form-group">
        <label class="form-label">标签</label>
        <input
          v-model="form.tags"
          type="text"
          class="form-input"
          placeholder="请输入标签，多个标签用逗号分隔（可选）"
        />
        <span class="form-hint">多个标签用逗号分隔，例如：前端,Vue3,技术分享</span>
      </div>

      <!-- 封面图片 -->
      <div class="form-group">
        <label class="form-label">封面图片</label>
        <input
          v-model="form.coverImage"
          type="text"
          class="form-input"
          placeholder="请输入封面图片URL（可选）"
        />
        <span class="form-hint">直接粘贴图片链接地址</span>
        <!-- 封面预览 -->
        <div v-if="form.coverImage" class="cover-preview">
          <img :src="form.coverImage" alt="封面预览" @error="form.coverImage = ''" />
        </div>
      </div>

      <!-- 作者ID（禁用） -->
      <div class="form-group">
        <label class="form-label">
          作者ID <span class="required">*</span>
        </label>
        <input
          v-model.number="form.authorId"
          type="number"
          class="form-input"
          :class="{ 'input-error': validationErrors.authorId }"
          placeholder="请输入作者ID"
          min="1"
          disabled
          @blur="validateForm"
        />
        <span v-if="validationErrors.authorId" class="field-error">
          {{ validationErrors.authorId }}
        </span>
        <span class="form-hint">作者ID不可修改</span>
      </div>

      <!-- 按钮组 -->
      <div class="form-actions">
        <button
          type="button"
          class="btn-cancel"
          @click="goBack"
        >
          取消
        </button>
        <button
          type="button"
          class="btn-reset"
          @click="handleReset"
        >
          重置
        </button>
        <button
          type="button"
          class="btn-submit"
          :disabled="submitting"
          @click="handleSubmit"
        >
          <span v-if="submitting" class="btn-spinner"></span>
          {{ submitting ? '保存中...' : '保存修改' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.post-edit-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.page-header {
  display: flex;
  align-items: center;
  gap: 16px;
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

.page-title {
  font-size: 24px;
  font-weight: 600;
  color: #1f2937;
  margin: 0;
}

/* 加载状态 */
.loading-state {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid #e5e7eb;
  border-top-color: #3b82f6;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* 错误状态 */
.error-state {
  text-align: center;
  padding: 60px 20px;
  color: #6b7280;
}

.error-message {
  color: #ef4444;
  margin-bottom: 16px;
  font-size: 16px;
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

/* 错误提示横幅 */
.error-banner {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background-color: #fef2f2;
  border: 1px solid #fecaca;
  border-radius: 6px;
  margin-bottom: 24px;
}

.error-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  background-color: #ef4444;
  color: #ffffff;
  border-radius: 50%;
  font-size: 12px;
  font-weight: bold;
}

.error-text {
  flex: 1;
  color: #dc2626;
  font-size: 14px;
}

.error-close {
  padding: 4px 8px;
  background: none;
  border: none;
  color: #9ca3af;
  font-size: 18px;
  cursor: pointer;
}

.error-close:hover {
  color: #6b7280;
}

/* 表单样式 */
.post-form {
  background: #ffffff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 32px;
}

.form-group {
  margin-bottom: 24px;
  position: relative;
}

.form-label {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  margin-bottom: 8px;
}

.required {
  color: #ef4444;
}

.form-input,
.form-textarea {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 14px;
  color: #1f2937;
  transition: border-color 0.2s, box-shadow 0.2s;
  box-sizing: border-box;
  background-color: #ffffff;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-input::placeholder,
.form-textarea::placeholder {
  color: #9ca3af;
}

.form-input:disabled {
  background-color: #f3f4f6;
  color: #9ca3af;
  cursor: not-allowed;
}

.form-textarea {
  resize: vertical;
  min-height: 200px;
  font-family: inherit;
  line-height: 1.6;
}

.input-error {
  border-color: #ef4444 !important;
}

.input-error:focus {
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1) !important;
}

.field-error {
  display: block;
  color: #ef4444;
  font-size: 12px;
  margin-top: 4px;
}

.char-count {
  position: absolute;
  right: 0;
  top: 0;
  font-size: 12px;
  color: #9ca3af;
}

.form-hint {
  display: block;
  color: #9ca3af;
  font-size: 12px;
  margin-top: 4px;
}

/* 封面预览 */
.cover-preview {
  margin-top: 12px;
  border-radius: 6px;
  overflow: hidden;
  border: 1px solid #e5e7eb;
}

.cover-preview img {
  width: 100%;
  max-height: 200px;
  object-fit: cover;
  display: block;
}

/* 按钮组 */
.form-actions {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid #e5e7eb;
}

.btn-cancel,
.btn-reset,
.btn-submit {
  padding: 10px 24px;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-cancel {
  background-color: #ffffff;
  border: 1px solid #d1d5db;
  color: #374151;
}

.btn-cancel:hover {
  background-color: #f3f4f6;
  border-color: #9ca3af;
}

.btn-reset {
  background-color: #ffffff;
  border: 1px solid #d1d5db;
  color: #6b7280;
}

.btn-reset:hover {
  background-color: #f3f4f6;
  border-color: #9ca3af;
  color: #374151;
}

.btn-submit {
  background-color: #10b981;
  border: 1px solid #10b981;
  color: #ffffff;
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-submit:hover:not(:disabled) {
  background-color: #059669;
  border-color: #059669;
}

.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-spinner {
  width: 14px;
  height: 14px;
  border: 2px solid #ffffff;
  border-top-color: transparent;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
