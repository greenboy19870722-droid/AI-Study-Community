<!--
  CommentInput.vue - 评论输入组件
  文本输入框，提交按钮
-->
<script setup>
import { ref, computed } from 'vue'
import { createComment, replyComment } from '@/api/comment.js'

const props = defineProps({
  postId: {
    type: [Number, String],
    required: true
  },
  replyingTo: {
    type: Object,
    default: null // { id, authorName }
  },
  authorId: {
    type: Number,
    default: 1 // TODO: 从用户状态获取
  }
})

const emit = defineEmits(['submit', 'cancel'])

// 输入内容
const content = ref('')
const submitting = ref(false)
const error = ref(null)

// 是否是回复模式
const isReplyMode = computed(() => {
  return props.replyingTo !== null
})

// 是否可以提交
const canSubmit = computed(() => {
  return content.value.trim() && !submitting.value
})

// 提交评论
const submit = async () => {
  if (!canSubmit.value) return

  error.value = null
  submitting.value = true

  try {
    if (isReplyMode.value) {
      // 回复评论
      await replyComment({
        postId: props.postId,
        parentId: props.replyingTo.id,
        authorId: props.authorId,
        content: content.value.trim()
      })
    } else {
      // 新增评论
      await createComment({
        postId: props.postId,
        authorId: props.authorId,
        content: content.value.trim()
      })
    }

    // 清空输入
    content.value = ''
    emit('submit')
  } catch (err) {
    error.value = err.message || '提交评论失败'
    console.error('提交评论失败:', err)
  } finally {
    submitting.value = false
  }
}

// 取消回复
const cancel = () => {
  content.value = ''
  error.value = null
  emit('cancel')
}
</script>

<template>
  <div class="comment-input">
    <!-- 回复提示 -->
    <div v-if="isReplyMode" class="reply-hint">
      <span>回复 @{{ replyingTo.authorName }}：</span>
      <button class="btn-cancel-reply" @click="cancel">取消回复</button>
    </div>

    <!-- 输入框 -->
    <div class="input-wrapper">
      <textarea
        v-model="content"
        class="comment-textarea"
        :placeholder="isReplyMode ? '写下你的回复...' : '写下你的评论...'"
        rows="3"
        :disabled="submitting"
        @keydown.ctrl.enter="submit"
        @keydown.meta.enter="submit"
      ></textarea>
    </div>

    <!-- 错误提示 -->
    <div v-if="error" class="input-error">
      {{ error }}
    </div>

    <!-- 操作按钮 -->
    <div class="input-actions">
      <span class="input-hint">Ctrl+Enter 提交</span>
      <button
        class="btn-submit"
        :disabled="!canSubmit"
        @click="submit"
      >
        {{ submitting ? '提交中...' : (isReplyMode ? '回复' : '发表评论') }}
      </button>
    </div>
  </div>
</template>

<style scoped>
.comment-input {
  width: 100%;
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
  transition: all 0.2s;
}

.btn-cancel-reply:hover {
  background-color: #fee2e2;
  border-color: #ef4444;
  color: #ef4444;
}

.input-wrapper {
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
  transition: border-color 0.2s, box-shadow 0.2s;
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

.input-error {
  color: #ef4444;
  font-size: 13px;
  margin-bottom: 12px;
}

.input-actions {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
}

.input-hint {
  font-size: 12px;
  color: #9ca3af;
}

.btn-submit {
  padding: 10px 24px;
  background-color: #3b82f6;
  color: #ffffff;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.btn-submit:hover:not(:disabled) {
  background-color: #2563eb;
}

.btn-submit:disabled {
  background-color: #9ca3af;
  cursor: not-allowed;
}
</style>
