<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { createPost } from "@/api/post.js";
import { ElMessage } from "element-plus";
import { getCurrentUserId } from "../../utils/auth";

const router = useRouter();

// 表单数据
const form = ref({
  title: "",
  content: "",
  authorId: getCurrentUserId(),
  tags: "",
  coverImage: "",
});

// 表单验证规则
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
};

// 表单引用
const formRef = ref(null);

// 提交状态
const submitting = ref(false);

// 提交表单
const handleSubmit = async () => {
  if (!formRef.value) return;

  try {
    await formRef.value.validate();
  } catch {
    return;
  }

  submitting.value = true;

  try {
    const submitData = {
      title: form.value.title.trim(),
      content: form.value.content.trim(),
      authorId: Number(form.value.authorId),
      tags: form.value.tags.trim() || undefined,
      coverImage: form.value.coverImage.trim() || undefined,
    };

    const result = await createPost(submitData);

    if (result && result.id) {
      ElMessage.success("发布成功");
      router.push(`/post/detail/${result.id}`);
    } else {
      ElMessage.success("发布成功");
      router.push("/post/list");
    }
  } catch (err) {
    ElMessage.error(err.message || "发布帖子失败");
  } finally {
    submitting.value = false;
  }
};

// 重置表单
const handleReset = () => {
  formRef.value?.resetFields();
};

// 返回列表页
const goBack = () => {
  router.push("/post/list");
};
</script>

<template>
  <div class="post-create-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <el-button @click="goBack"> <span>&larr;</span> 返回列表 </el-button>
      <el-text size="large" tag="b">发布帖子</el-text>
    </div>

    <!-- 发布表单 -->
    <el-card class="form-card">
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
          <!-- 封面预览 -->
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

        <!-- 按钮组 -->
        <el-form-item class="form-actions">
          <el-button @click="goBack">取消</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="primary" :loading="submitting" @click="handleSubmit">
            {{ submitting ? "发布中..." : "发布帖子" }}
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<style scoped>
.post-create-container {
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

.form-actions {
  margin-top: 24px;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>
