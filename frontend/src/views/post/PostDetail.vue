<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getPostDetail } from "@/api/post.js";
import { getCommentTree, createComment, deleteComment } from "@/api/comment.js";
import { getCurrentUserId, isLoggedIn } from "@/utils/auth.js";
import { ElMessage, ElMessageBox } from "element-plus";

const route = useRoute();
const router = useRouter();

const post = ref(null);
const comments = ref([]);
const totalComments = ref(0);
const commentPage = ref(1);
const commentPageSize = ref(10);
const commentLoading = ref(false);
const loading = ref(false);

// 评论输入
const commentContent = ref("");
const replyingTo = ref(null);
const commentSubmitting = ref(false);

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

// 获取状态文本
const getStatusText = (status) => {
  const statusMap = { 0: "已删除", 1: "正常", 2: "置顶", 3: "推荐" };
  return statusMap[status] || "未知";
};

// 获取状态类型
const getStatusType = (status) => {
  const typeMap = { 0: "info", 1: "success", 2: "warning", 3: "primary" };
  return typeMap[status] || "info";
};

// 当前登录用户
const currentUserId = () => getCurrentUserId();
const loggedIn = () => isLoggedIn();

// 加载帖子详情
const loadPostDetail = async () => {
  const postId = route.params.id || route.query.id;
  if (!postId) {
    ElMessage.error("帖子ID不存在");
    return;
  }

  loading.value = true;
  try {
    const data = await getPostDetail(postId);
    post.value = data;
  } catch (err) {
    ElMessage.error(err.message || "加载帖子详情失败");
  } finally {
    loading.value = false;
  }
};

// 加载评论列表
const loadComments = async () => {
  const postId = route.params.id || route.query.id;
  if (!postId) return;

  commentLoading.value = true;
  try {
    const data = await getCommentTree({
      postId,
      page: commentPage.value,
      pageSize: commentPageSize.value,
    });
    comments.value = data.list || [];
    totalComments.value = data.total || 0;
  } catch (err) {
    ElMessage.error(err.message || "加载评论失败");
  } finally {
    commentLoading.value = false;
  }
};

// 提交评论
const submitComment = async () => {
  if (!loggedIn()) {
    ElMessage.warning("请先登录后再评论");
    return;
  }

  if (!commentContent.value.trim()) {
    ElMessage.warning("评论内容不能为空");
    return;
  }

  const postId = route.params.id || route.query.id;
  if (!postId) return;

  commentSubmitting.value = true;
  try {
    await createComment({
      postId,
      authorId: currentUserId(),
      content: commentContent.value.trim(),
    });

    ElMessage.success("评论成功");
    commentContent.value = "";
    replyingTo.value = null;
    commentPage.value = 1;
    await loadComments();
  } catch (err) {
    ElMessage.error(err.message || "提交评论失败");
  } finally {
    commentSubmitting.value = false;
  }
};

// 回复评论
const handleReply = (comment) => {
  replyingTo.value = {
    id: comment.id,
    authorName: comment.authorName || `用户${comment.authorId}`,
  };
};

// 取消回复
const cancelReply = () => {
  replyingTo.value = null;
  commentContent.value = "";
};

// 提交回复
const submitReply = async () => {
  if (!loggedIn()) {
    ElMessage.warning("请先登录后再评论");
    return;
  }

  if (!commentContent.value.trim()) {
    ElMessage.warning("评论内容不能为空");
    return;
  }

  const postId = route.params.id || route.query.id;
  if (!postId || !replyingTo.value) return;

  commentSubmitting.value = true;
  try {
    await createComment({
      postId,
      parentId: replyingTo.value.id,
      authorId: currentUserId(),
      content: commentContent.value.trim(),
    });

    ElMessage.success("回复成功");
    commentContent.value = "";
    replyingTo.value = null;
    await loadComments();
  } catch (err) {
    ElMessage.error(err.message || "提交回复失败");
  } finally {
    commentSubmitting.value = false;
  }
};

// 删除评论
const handleDeleteComment = async (comment) => {
  if (!loggedIn() || currentUserId() !== comment.authorId) {
    ElMessage.warning("只能删除自己的评论");
    return;
  }

  try {
    await ElMessageBox.confirm("确定要删除这条评论吗？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    });

    await deleteComment({ id: comment.id, authorId: comment.authorId });
    ElMessage.success("删除成功");
    await loadComments();
  } catch {
    // 用户取消
  }
};

// 评论翻页
const handleCommentPageChange = (newPage) => {
  commentPage.value = newPage;
  loadComments();
};

// 返回列表页
const goBack = () => {
  router.push("/post/list");
};

// 跳转到编辑页
const goToEdit = () => {
  router.push(`/post/edit/${post.value.id}`);
};

onMounted(() => {
  loadPostDetail();
  loadComments();
});
</script>

<template>
  <div class="post-detail-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="goBack"> <span>&larr;</span> 返回列表 </el-button>
      <el-button v-if="post && post.id" type="primary" @click="goToEdit">
        编辑帖子
      </el-button>
    </div>

    <!-- 加载状态 -->
    <div v-loading="loading" class="loading-container">
      <!-- 帖子详情 -->
      <el-card v-if="post && post.id" class="post-card">
        <template #header>
          <div class="post-header">
            <div class="post-badges">
              <el-tag
                v-if="post.status !== 1"
                :type="getStatusType(post.status)"
              >
                {{ getStatusText(post.status) }}
              </el-tag>
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
              <span v-if="post.updatedAt !== post.createdAt" class="meta-item">
                <span class="meta-label">更新时间:</span>
                <span class="meta-value">{{ formatDate(post.updatedAt) }}</span>
              </span>
            </div>
          </div>
        </template>

        <!-- 标签 -->
        <div v-if="post.tags" class="post-tags">
          <el-tag
            v-for="(tag, index) in post.tags.split(',')"
            :key="index"
            size="small"
            effect="plain"
          >
            {{ tag.trim() }}
          </el-tag>
        </div>

        <!-- 封面图片 -->
        <div v-if="post.coverImage" class="post-cover">
          <el-image :src="post.coverImage" fit="cover" />
        </div>

        <!-- 帖子内容 -->
        <div class="post-content">
          <p>{{ post.content }}</p>
        </div>

        <!-- 帖子底部 -->
        <div class="post-footer">
          <el-button type="primary" plain>
            <span class="btn-icon">♥</span> 点赞
          </el-button>
          <el-button type="info" plain>分享</el-button>
        </div>
      </el-card>

      <!-- 空状态 -->
      <el-empty
        v-if="!loading && (!post || !post.id)"
        description="帖子不存在或已被删除"
      >
        <el-button type="primary" @click="goBack">返回列表</el-button>
      </el-empty>
    </div>

    <!-- 评论区 -->
    <el-card v-if="post && post.id" class="comment-card">
      <template #header>
        <div class="comment-header">
          <h3 class="comment-title">评论 ({{ totalComments }})</h3>
        </div>
      </template>

      <!-- 评论输入区 -->
      <div class="comment-input-area" v-if="!replyingTo">
        <!-- 输入框 -->
        <el-input
          v-model="commentContent"
          type="textarea"
          :placeholder="'写下你的评论...'"
          :rows="3"
          :disabled="commentSubmitting"
        />

        <!-- 提交按钮 -->
        <div class="comment-actions">
          <el-button
            type="primary"
            :loading="commentSubmitting"
            :disabled="!commentContent.trim()"
            @click="replyingTo ? submitReply() : submitComment()"
            >发表评论
          </el-button>
        </div>
      </div>

      <!-- 评论加载状态 -->
      <div v-loading="commentLoading" class="comment-loading">
        <el-empty
          v-if="!commentLoading && comments.length === 0"
          description="暂无评论"
        >
          <span class="comment-empty-text">来发表第一篇评论吧</span>
        </el-empty>

        <!-- 评论列表 -->
        <div v-else class="comment-list">
          <div
            v-for="comment in comments"
            :key="comment.id"
            class="comment-item"
          >
            <div class="comment-main">
              <div class="comment-author">
                <span class="author-name">{{
                  comment.authorName || `用户${comment.authorId}`
                }}</span>
                <span class="comment-time">{{
                  formatDate(comment.createdAt)
                }}</span>
              </div>
              <div class="comment-body">{{ comment.content }}</div>
              <div class="comment-actions-row">
                <div
                  v-if="replyingTo && comment.id === replyingTo.id"
                  class="reply-hint"
                >
                  <div></div>
                  <p>回复 @{{ replyingTo.authorName }}：</p>
                  <div class="comment-input-area">
                    <!-- 输入框 -->
                    <el-input
                      v-model="commentContent"
                      type="textarea"
                      :placeholder="'写下你的回复...'"
                      :rows="3"
                      :disabled="commentSubmitting"
                    />

                    <!-- 提交按钮 -->
                    <div class="comment-actions">
                      <el-button @click="cancelReply">取消回复</el-button>
                      <el-button
                        type="primary"
                        :loading="commentSubmitting"
                        :disabled="!commentContent.trim()"
                        @click="replyingTo ? submitReply() : submitComment()"
                      >
                        {{
                          commentSubmitting
                            ? "提交中..."
                            : replyingTo
                            ? "回复"
                            : "发表评论"
                        }}
                      </el-button>
                    </div>
                  </div>
                </div>
                <el-button v-else size="small" @click="handleReply(comment)"
                  >回复</el-button
                >

                <el-button
                  v-if="loggedIn() && currentUserId() === comment.authorId"
                  size="small"
                  type="danger"
                  plain
                  @click="handleDeleteComment(comment)"
                >
                  删除
                </el-button>
              </div>
            </div>

            <!-- 子评论 -->
            <div v-if="comment.children?.length > 0" class="comment-children">
              <div
                v-for="child in comment.children"
                :key="child.id"
                class="comment-child"
              >
                <div class="child-author">
                  <span class="author-name">{{
                    child.authorName || `用户${child.authorId}`
                  }}</span>
                  <span v-if="child.replyToAuthor" class="reply-arrow"
                    >回复</span
                  >
                  <span v-if="child.replyToAuthor" class="reply-target"
                    >@{{ child.replyToAuthor }}</span
                  >
                  <span class="comment-time">{{
                    formatDate(child.createdAt)
                  }}</span>
                </div>
                <div class="comment-body">{{ child.content }}</div>
                <div class="comment-actions-row">
                  <el-button
                    v-if="loggedIn() && currentUserId() === child.authorId"
                    size="small"
                    type="danger"
                    plain
                    @click="handleDeleteComment(child)"
                  >
                    删除
                  </el-button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 评论分页 -->
        <div v-if="comments.length > 0" class="comment-pagination">
          <el-pagination
            v-model:current-page="commentPage"
            :total="totalComments"
            :page-size="commentPageSize"
            layout="prev, pager, next"
            background
            small
            @current-change="handleCommentPageChange"
          />
        </div>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.post-detail-container {
  margin: 0 auto;
  padding: 24px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.loading-container {
  min-height: 200px;
}

.post-card {
  border-radius: 12px;
  margin-bottom: 24px;
}

.post-header {
  padding: 0;
}

.post-badges {
  display: flex;
  gap: 8px;
  margin-bottom: 12px;
}

.post-title {
  font-size: 28px;
  font-weight: 700;
  color: var(--text-h);
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
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border);
}

.post-cover {
  margin-bottom: 20px;
  border-radius: 8px;
  overflow: hidden;
}

.post-cover .el-image {
  width: 100%;
  max-height: 400px;
}

.post-content {
  color: var(--text);
  font-size: 16px;
  line-height: 1.8;
  margin-bottom: 24px;
}

.post-content p {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-word;
}

.post-footer {
  display: flex;
  gap: 12px;
  padding-top: 20px;
  border-top: 1px solid var(--border);
}

.btn-icon {
  margin-right: 4px;
}

/* 评论区 */
.comment-card {
  border-radius: 12px;
}

.comment-header {
  padding: 0;
}

.comment-title {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
}

.comment-input-area {
  margin-bottom: 24px;
  width: 100%;
}

.reply-hint {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 12px;
  margin-bottom: 12px;
  padding: 8px 12px;
  background: var(--accent-bg);
  border-radius: 6px;
  font-size: 14px;
  color: var(--accent);
}

.comment-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 12px;
}

.comment-loading {
  min-height: 100px;
}

.comment-empty-text {
  color: #9ca3af;
  font-size: 14px;
}

.comment-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.comment-item {
  padding-bottom: 20px;
  border-bottom: 1px solid var(--border);
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
  color: var(--text-h);
  font-size: 14px;
}

.comment-time {
  color: #9ca3af;
  font-size: 12px;
}

.comment-body {
  color: var(--text);
  font-size: 14px;
  line-height: 1.6;
  margin-bottom: 8px;
  white-space: pre-wrap;
  word-break: break-word;
}

.comment-actions-row {
  display: flex;
  gap: 8px;
}

/* 子评论 */
.comment-children {
  margin-left: 24px;
  padding-left: 16px;
  border-left: 2px solid var(--border);
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.comment-child {
  padding: 12px;
  background: var(--social-bg);
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
  color: var(--accent);
  font-size: 12px;
}

.comment-pagination {
  display: flex;
  justify-content: center;
  margin-top: 24px;
  padding-top: 16px;
  border-top: 1px solid var(--border);
}
</style>
