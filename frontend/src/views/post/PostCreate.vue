<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { createPost } from '@/api/post.js'

const router = useRouter()

// 表单数据
const form = ref({
  title: '',
  content: '',
  authorId: 1, // TODO: 后续从登录用户上下文获取
  tags: '',
  coverImage: '',
})

// 表单验证
const formRules = {
  title: [
    { required: true, message: '请输入帖子标题', trigger: 'blur' },
    { min: 3, max: 100, message: '标题长度在 3 到 100 个字符', trigger: 'blur' },
  ],
  content: [
    { required: true, message: '请输入帖子内容', trigger: 'blur' },
    { min: 10, message: '内容至少 10 个字符', trigger: 'blur' },
  ],
  authorId: [
    { required: true, message: '请输入作者ID', trigger: 'blur' },
    { type: 'number', message: '作者ID必须为数字', trigger: 'blur' },
  ],
}

// 提交状态
const submitting = ref(false)
const error = ref(null)
const validationErrors = ref({})

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

// 提交表单
const handleSubmit = async () => {
  // 清除之前的错误
  error.value = null
  
  // 验证表单
  if (!validateForm()) {
    return
  }
  
  submitting.value = true
  
  try {
    const submitData = {
      title: form.value.title.trim(),
      content: form.value.content.trim(),
      authorId: Number(form.value.authorId),
      tags: form.value.tags.trim() || undefined,
      coverImage: form.value.coverImage.trim() || undefined,
    }
    
    const result = await createPost(submitData)
    
    // 发布成功，跳转到列表页
    if (result && result.id) {
      router.push(`/post/detail/${result.id}`)
    } else {
      // 如果没有返回id，也跳转到列表页
      router.push('/post/list')
    }
  } catch (err) {
    error.value = err.message || '发布帖子失败，请稍后重试'
    console.error('发布帖子失败:', err)
  } finally {
    submitting.value = false
  }
}

// 重置表单
const handleReset = () => {
  form.value = {
    title: '',
    content: '',
    authorId: 1,
    tags: '',
    coverImage: '',
  }
  validationErrors.value = {}
  error.value = null
}

// 返回列表页
const goBack = () => {
  router.push('/post/list')
}
</script>

<template>
  <div class="post-create-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <button class="btn-back" @click="goBack">
        &larr; 返回列表
      </button>
      <h1 class="page-title">发布帖子</h1>
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="error-banner">
      <span class="error-icon">!</span>
      <span class="error-text">{{ error }}</span>
      <button class="error-close" @click="error = null">&times;</button>
    </div>

    <!-- 发布表单 -->
    <div class="post-form">
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

      <!-- 作者ID -->
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
          @blur="validateForm"
        />
        <span v-if="validationErrors.authorId" class="field-error">
          {{ validationErrors.authorId }}
        </span>
        <span class="form-hint">请输入正确的用户ID</span>
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
          {{ submitting ? '发布中...' : '发布帖子' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.post-create-container {
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
  background-color: #3b82f6;
  border: 1px solid #3b82f6;
  color: #ffffff;
  display: flex;
  align-items: center;
  gap: 8px;
}

.btn-submit:hover:not(:disabled) {
  background-color: #2563eb;
  border-color: #2563eb;
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
