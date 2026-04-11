<template>
  <div class="login-page" style="background: var(--el-bg-color-page)">
    <el-card class="login-card" shadow="hover">
      <h2 class="login-title">系统登录</h2>

      <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="80px"
        class="login-form"
      >
        <el-form-item label="昵称" prop="username">
          <el-input v-model="form.username" placeholder="请输入昵称" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" class="login-btn" @click="handleLogin">
            登 录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { ElMessage } from "element-plus";
import router from "@/router";
import { addUser } from "@/api/user";
// 表单数据
const form = ref({
  username: "",
});

// 表单验证规则
const rules = {
  username: [{ required: true, message: "请输入账号", trigger: "blur" }],
};

const handleLogin = async () => {
  const { id } = await addUser(form.value.username);
  localStorage.setItem("authorId", id);
  localStorage.setItem("username", form.value.username);
  ElMessage.success("登录成功！");
  router.back();
};
</script>

<style scoped>
.login-page {
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.login-card {
  width: 420px;
  padding: 30px;
}

.login-title {
  text-align: center;
  margin-bottom: 30px;
  font-size: 24px;
  color: var(--el-text-color-primary);
}

.login-btn {
  width: 100%;
}

.flex-between {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
</style>