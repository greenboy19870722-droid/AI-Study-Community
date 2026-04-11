<script setup>
import { ref, onMounted } from "vue";
import { getPostList } from "@/api/post.js";
import { ElMessage } from "element-plus";

// 帖子列表数据
const postList = ref([]);
const total = ref(0);

// 分页参数
const pagination = ref({
  page: 1,
  pageSize: 10,
});

// 加载状态
const loading = ref(false);

// 格式化日期
const formatDate = (dateString) => {
  if (!dateString) return "";
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");
  return `${year}-${month}-${day} ${hours}:${minutes}`;
};

// 获取状态标签类型
const getStatusType = (status) => {
  const types = {
    2: "warning", // 置顶
    3: "primary", // 推荐
  };
  return types[status] || "info";
};

// 获取状态标签文本
const getStatusText = (status) => {
  const texts = {
    2: "置顶",
    3: "推荐",
  };
  return texts[status] || "";
};

// 加载帖子列表
const loadPostList = async () => {
  loading.value = true;
  try {
    const params = {
      page: pagination.value.page,
      pageSize: pagination.value.pageSize,
    };
    const data = await getPostList(params);
    postList.value = data.list || [];
    total.value = data.total || 0;
  } catch (err) {
    ElMessage.error(err.message || "加载帖子列表失败");
  } finally {
    loading.value = false;
  }
};

// 翻页
const handlePageChange = (newPage) => {
  pagination.value.page = newPage;
  loadPostList();
};

// 每页数量改变
const handleSizeChange = (newSize) => {
  pagination.value.pageSize = newSize;
  pagination.value.page = 1;
  loadPostList();
};

// 跳转到帖子详情
const goToDetail = (postId) => {
  window.location.href = `/post/detail/${postId}`;
};

// 跳转到发布帖子页
const goToCreate = () => {
  window.location.href = "/post/create";
};

// 组件挂载时加载数据
onMounted(() => {
  loadPostList();
});
</script>

<template>
  <div class="post-list-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <h1 class="page-title">帖子列表</h1>
      <el-button type="primary" @click="goToCreate">
        <span class="btn-icon">+</span> 发布帖子
      </el-button>
    </div>

    <!-- 加载状态 -->
    <div v-loading="loading" class="loading-container">
      <!-- 空状态 -->
      <el-empty v-if="!loading && postList.length === 0" description="暂无帖子">
        <el-button type="primary" @click="goToCreate">发布第一个帖子</el-button>
      </el-empty>

      <!-- 帖子列表 -->
      <div v-else class="post-list">
        <el-card
          v-for="post in postList"
          :key="post.id"
          class="post-card"
          shadow="hover"
          @click="goToDetail(post.id)"
        >
          <template #header>
            <div class="card-header">
              <div class="title-row">
                <h3 class="post-title">{{ post.title }}</h3>
                <el-tag
                  v-if="post.status"
                  :type="getStatusType(post.status)"
                  size="small"
                >
                  {{ getStatusText(post.status) }}
                </el-tag>
              </div>
            </div>
          </template>

          <div class="card-body">
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
                <el-tag size="small" type="info">{{ post.tags }}</el-tag>
              </span>
            </div>
          </div>
        </el-card>
      </div>

      <!-- 分页组件 -->
      <div v-if="postList.length > 0" class="pagination-container">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="total"
          :page-sizes="[5, 10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          background
          @current-change="handlePageChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.post-list-container {
  margin: 0 auto;
  padding: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid var(--border);
}

.page-title {
  font-size: 28px;
  font-weight: 600;
  margin: 0;
  background: linear-gradient(135deg, #08060d 0%, #aa3bff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.btn-icon {
  margin-right: 4px;
  font-weight: bold;
}

.loading-container {
  min-height: 300px;
}

.post-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.post-card {
  cursor: pointer;
  transition: all 0.3s;
  border-radius: 12px;
}

.post-card:hover {
  transform: translateY(-4px);
}

.card-header {
  padding: 0;
}

.title-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.post-title {
  font-size: 18px;
  font-weight: 600;
  color: var(--text-h);
  margin: 0;
  flex: 1;
}

.card-body {
  padding-top: 0;
}

.post-content {
  color: var(--text);
  font-size: 14px;
  line-height: 1.6;
  margin: 0 0 16px;
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

.meta-value {
  color: var(--text);
}

.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 32px;
  padding-top: 24px;
  border-top: 1px solid var(--border);
}
</style>
