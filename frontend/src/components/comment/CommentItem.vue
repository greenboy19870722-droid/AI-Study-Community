<!--
  CommentItem.vue - 单条评论组件
  显示单条评论内容（作者、时间、内容）
-->
<script setup>
import { computed } from 'vue'
import ReplyItem from './ReplyItem.vue'

const props = defineProps({
  comment: {
    type: Object,
    required: true
  },
  currentUserId: {
    type: Number,
    default: null
  }
})

const emit = defineEmits(['reply', 'delete'])

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

// 作者名称
const authorName = computed(() => {
  return props.comment.authorName || `用户${props.comment.authorId}`
})

// 是否有子评论
const hasChildren = computed(() => {
  return props.comment.children && props.comment.children.length > 0
})

// 是否可以删除（作者本人）
const canDelete = computed(() => {
  return props.currentUserId && props.comment.authorId === props.currentUserId
})

// 处理回复
const handleReply = () => {
  emit('reply', props.comment)
}

// 处理删除
const handleDelete = () => {
  emit('delete', props.comment)
}
</script>

<template>
  <div class="comment-item">
    <!-- 主评论 -->
    <div class="comment-main">
      <div class="comment-author">
        <span class="author-name">{{ authorName }}</span>
        <span class="comment-time">{{ formatDate(comment.createdAt) }}</span>
      </div>
      <div class="comment-body">{{ comment.content }}</div>
      <div class="comment-actions-row">
        <button class="btn-reply" @click="handleReply">回复</button>
        <button
          v-if="canDelete"
          class="btn-delete"
          @click="handleDelete"
        >
          删除
        </button>
      </div>
    </div>

    <!-- 子评论区域 -->
    <div v-if="hasChildren" class="comment-children">
      <ReplyItem
        v-for="child in comment.children"
        :key="child.id"
        :reply="child"
        @reply="() => emit('reply', child)"
      />
    </div>
  </div>
</template>

<style scoped>
.comment-item {
  padding-bottom: 16px;
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

.btn-delete {
  padding: 4px 12px;
  background-color: transparent;
  border: 1px solid #fca5a5;
  border-radius: 4px;
  font-size: 12px;
  color: #ef4444;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-delete:hover {
  background-color: #fef2f2;
  border-color: #ef4444;
}

.comment-children {
  margin-left: 24px;
  padding-left: 16px;
  border-left: 2px solid #e5e7eb;
  display: flex;
  flex-direction: column;
  gap: 12px;
}
</style>
