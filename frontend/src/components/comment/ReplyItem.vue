<!--
  ReplyItem.vue - 回复项组件
  显示二级回复内容，简洁样式，嵌套显示
-->
<script setup>
import { computed } from 'vue'

const props = defineProps({
  reply: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['reply'])

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

// 回复者名称
const authorName = computed(() => {
  return props.reply.authorName || `用户${props.reply.authorId}`
})

// 是否显示回复目标
const showReplyTo = computed(() => {
  return props.reply.replyToAuthor && props.reply.replyToUserId
})

// 处理回复
const handleReply = () => {
  emit('reply', props.reply)
}
</script>

<template>
  <div class="reply-item">
    <div class="reply-author">
      <span class="author-name">{{ authorName }}</span>
      <span v-if="showReplyTo" class="reply-arrow">回复</span>
      <span v-if="showReplyTo" class="reply-target">@{{ reply.replyToAuthor }}</span>
      <span class="reply-time">{{ formatDate(reply.createdAt) }}</span>
    </div>
    <div class="reply-body">{{ reply.content }}</div>
    <div class="reply-actions">
      <button class="btn-reply" @click="handleReply">回复</button>
    </div>
  </div>
</template>

<style scoped>
.reply-item {
  padding: 12px;
  background-color: #f9fafb;
  border-radius: 6px;
}

.reply-author {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 6px;
  font-size: 13px;
}

.author-name {
  font-weight: 600;
  color: #1f2937;
}

.reply-arrow {
  color: #9ca3af;
  font-size: 12px;
}

.reply-target {
  color: #3b82f6;
  font-size: 12px;
}

.reply-time {
  color: #9ca3af;
  font-size: 12px;
}

.reply-body {
  color: #374151;
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 6px;
  white-space: pre-wrap;
  word-break: break-word;
}

.reply-actions {
  display: flex;
  gap: 8px;
}

.btn-reply {
  padding: 2px 8px;
  background-color: transparent;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 11px;
  color: #6b7280;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-reply:hover {
  background-color: #eff6ff;
  border-color: #3b82f6;
  color: #3b82f6;
}
</style>
