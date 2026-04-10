<!--
  CommentPagination.vue - 评论分页组件
  分页控件：上一页/下一页
-->
<script setup>
import { computed } from 'vue'

const props = defineProps({
  currentPage: {
    type: Number,
    default: 1
  },
  totalPages: {
    type: Number,
    default: 1
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['page-change'])

// 是否有上一页
const hasPrev = computed(() => {
  return props.currentPage > 1
})

// 是否有下一页
const hasNext = computed(() => {
  return props.currentPage < props.totalPages
})

// 上一页
const prevPage = () => {
  if (hasPrev.value && !props.loading) {
    emit('page-change', props.currentPage - 1)
  }
}

// 下一页
const nextPage = () => {
  if (hasNext.value && !props.loading) {
    emit('page-change', props.currentPage + 1)
  }
}
</script>

<template>
  <div class="comment-pagination" :class="{ loading }">
    <button
      class="page-btn prev"
      :disabled="!hasPrev || loading"
      @click="prevPage"
    >
      <span v-if="loading && currentPage > 1" class="btn-loading">加载中...</span>
      <span v-else>上一页</span>
    </button>

    <div class="page-info">
      <span class="current">{{ currentPage }}</span>
      <span class="separator">/</span>
      <span class="total">{{ totalPages }}</span>
    </div>

    <button
      class="page-btn next"
      :disabled="!hasNext || loading"
      @click="nextPage"
    >
      <span v-if="loading && currentPage < totalPages" class="btn-loading">加载中...</span>
      <span v-else>下一页</span>
    </button>
  </div>
</template>

<style scoped>
.comment-pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
  padding: 16px 0;
}

.comment-pagination.loading {
  opacity: 0.7;
  pointer-events: none;
}

.page-btn {
  padding: 8px 16px;
  border: 1px solid #d1d5db;
  background-color: #ffffff;
  color: #374151;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
  min-width: 80px;
}

.page-btn:hover:not(:disabled) {
  background-color: #f3f4f6;
  border-color: #3b82f6;
  color: #3b82f6;
}

.page-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.page-btn.prev {
  margin-right: auto;
}

.page-btn.next {
  margin-left: auto;
}

.page-info {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: #6b7280;
}

.page-info .current {
  font-weight: 600;
  color: #1f2937;
}

.page-info .separator {
  color: #d1d5db;
}

.page-info .total {
  color: #6b7280;
}

.btn-loading {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}
</style>
