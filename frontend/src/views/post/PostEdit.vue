<script setup>
import { ref, onMounted } from "vue";
import { useRoute, useRouter } from "vue-router";
import { getPostDetail, updatePost } from "@/api/post.js";
import { ElMessage } from "element-plus";

const route = useRoute();
const router = useRouter();

const postId = ref(null);
const form = ref({
  title: "",
  content: "",
  authorId: 1,
  tags: "",
  coverImage: "",
});
const originalData = ref(null);
const loading = ref(false);
const submitting = ref(false);

const formRef = ref(null);

const formRules = {
  title: [
    { required: true, message: "请输入帖子标题", trigger: "blur" },
    {
      min: 3,
      max: 100,
      message: "标题长度在 3 到 100 个字符",
      trigger: "blur",
    },
  ],
  content: [
    { required: true, message: "请输入帖子内容", trigger: "blur" },
    { min: 10, message: "内容至少 10 个字符", trigger: "blur" },
  ],
  authorId: [{ required: true, message: "请输入作者ID", trigger: "blur" }],
};

// 加载帖子数据
const loadPostData = async () => {
  postId.value = route.params.id || route.query.id;

  if (!postId.value) {
    ElMessage.error("帖子ID不存在");
    return;
  }

  loading.value = true;

  try {
    const data = await getPostDetail(postId.value);

    if (!data || !data.id) {
      ElMessage.error("帖子不存在或已被删除");
      return;
    }

    form.value = {
      title: data.title || "",
      content: data.content || "",
      authorId: data.authorId || 1,
      tags: data.tags || "",
      coverImage: data.coverImage || "",
    };

    originalData.value = { ...form.value };
    document.title = `编辑帖子 - ${data.title}`;
  } catch (err) {
    ElMessage.error(err.message || "加载帖子数据失败");
  } finally {
    loading.value = false;
  }
};

// 检测表单是否有变更
const hasChanges = () => {
  if (!originalData.value) return false;
  return (
    form.value.title !== originalData.value.title ||
    form.value.content !== originalData.value.content ||
    form.value.tags !== originalData.value.tags ||
    form.value.coverImage !== originalData.value.coverImage
  );
};

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return;

  try {
    await formRef.value.validate();
  } catch {
    return;
  }

  if (!hasChanges()) {
    ElMessage.warning("没有检测到任何变更");
    return;
  }

  submitting.value = true;

  try {
    const submitData = {
      id: Number(postId.value),
      title: form.value.title.trim(),
      content: form.value.content.trim(),
      authorId: Number(form.value.authorId),
      tags: form.value.tags.trim() || undefined,
      coverImage: form.value.coverImage.trim() || undefined,
    };

    await updatePost(submitData);
    ElMessage.success("更新成功");
    router.push(`/post/detail/${postId.value}`);
  } catch (err) {
    ElMessage.error(err.message || "更新帖子失败");
  } finally {
    submitting.value = false;
  }
};

// 重置表单
const handleReset = () => {
  if (originalData.value) {
    form.value = { ...originalData.value };
  }
  formRef.value?.clearValidate();
};

// 返回详情页
const goBack = () => {
  router.push(`/post/detail/${postId.value}`);
};

// 返回列表页
const goToList = () => {
  router.push("/post/list");
};

onMounted(() => {
  loadPostData();
});
</script>

<template>
  <div class="post-edit-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="goBack"> <span>&larr;</span> 返回详情 </el-button>
      <h1 class="page-title">编辑帖子</h1>
    </div>

    <!-- 加载状态 -->
    <div v-loading="loading" class="loading-container">
      <el-card v-if="!loading && postId" class="form-card">
        <el-form
          ref="formRef"
          :model="form"
          :rules="formRules"
          label-position="top"
          class="post-form"
        >
          <!-- 标题 -->
          <el-form-item label="帖子标题" prop="title">
            <el-input
              v-model="form.title"
              placeholder="请输入帖子标题（3-100个字符）"
              maxlength="100"
              show-word-limit
              clearable
            />
          </el-form-item>

          <!-- 内容 -->
          <el-form-item label="帖子内容" prop="content">
            <el-input
              v-model="form.content"
              type="textarea"
              placeholder="请输入帖子内容（至少10个字符）"
              :rows="10"
              show-word-limit
            />
          </el-form-item>

          <!-- 标签 -->
          <el-form-item label="标签" prop="tags">
            <el-input
              v-model="form.tags"
              placeholder="多个标签用逗号分隔，例如：前端,Vue3,技术分享"
              clearable
            />
          </el-form-item>

          <!-- 封面图片 -->
          <el-form-item label="封面图片" prop="coverImage">
            <el-input
              v-model="form.coverImage"
              placeholder="请输入封面图片URL（可选）"
              clearable
            />
            <div v-if="form.coverImage" class="cover-preview">
              <el-image
                :src="form.coverImage"
                fit="cover"
                @error="form.coverImage = ''"
              >
                <template #error>
                  <div class="image-error">
                    <span>图片加载失败</span>
                  </div>
                </template>
              </el-image>
            </div>
          </el-form-item>

          <!-- 作者ID（禁用） -->
          <el-form-item label="作者ID" prop="authorId">
            <el-input-number
              v-model="form.authorId"
              :min="1"
              controls-position="right"
              disabled
            />
            <span class="form-tip">作者ID不可修改</span>
          </el-form-item>

          <!-- 按钮组 -->
          <el-form-item class="form-actions">
            <el-button @click="goBack">取消</el-button>
            <el-button @click="handleReset">重置</el-button>
            <el-button
              type="primary"
              :loading="submitting"
              @click="handleSubmit"
            >
              {{ submitting ? "保存中..." : "保存修改" }}
            </el-button>
          </el-form-item>
        </el-form>
      </el-card>

      <!-- 加载失败状态 -->
      <el-result
        v-if="!loading && !postId"
        icon="error"
        title="加载失败"
        sub-title="帖子不存在或已被删除"
      >
        <template #extra>
          <el-button type="primary" @click="goToList">返回列表</el-button>
        </template>
      </el-result>
    </div>
  </div>
</template>

<style scoped>
.post-edit-container {
  margin: 0 auto;
  padding: 24px;
}

.page-header {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.page-title {
  font-size: 24px;
  font-weight: 600;
  margin: 0;
  background: linear-gradient(135deg, #08060d 0%, #aa3bff 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.loading-container {
  min-height: 300px;
}

.form-card {
  border-radius: 12px;
}

.post-form {
  max-width: 100%;
}

.cover-preview {
  margin-top: 12px;
  border-radius: 8px;
  overflow: hidden;
  border: 1px solid var(--border);
  max-height: 200px;
}

.cover-preview .el-image {
  width: 100%;
  height: 200px;
}

.image-error {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  background: #f5f5f5;
  color: #999;
  font-size: 14px;
}

.form-tip {
  display: block;
  margin-top: 4px;
  font-size: 12px;
  color: #9ca3af;
}

.form-actions {
  margin-top: 24px;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
