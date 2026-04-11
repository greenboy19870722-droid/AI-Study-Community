<template>
  <el-config-provider>
    <el-container>
      <el-header>
        <el-row justify="end">
          <el-space size="medium">
            <el-text tag="b">主题切换</el-text>
            <el-switch v-model="isDark" inline-prompt @change="toggleDark" />
            <el-link
              v-if="!isLoggedIn"
              type="primary"
              @click="$router.push('/login')"
              >登录</el-link
            >

            <el-dropdown v-else>
              <span class="el-dropdown-link">
                {{ username }}
                <el-icon class="el-icon--right">
                  <arrow-down />
                </el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="handleLogout"
                    >点击退出登录</el-dropdown-item
                  >
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </el-space>
        </el-row>
      </el-header>

      <el-row>
        <el-aside>
          <el-menu router default-active="1" class="el-menu-vertical-demo">
            <el-menu-item index="/post/list">
              <i class="el-icon-house"></i>
              <span>首页</span>
            </el-menu-item>
            <el-menu-item index="/">
              <i class="el-icon-info"></i>
              <span>关于</span>
            </el-menu-item>
          </el-menu>
        </el-aside>

        <el-main>
          <el-row justify="space-between">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item :to="{ path: '/post/list' }"
                >首页</el-breadcrumb-item
              >
              <el-breadcrumb-item :to="{ path: '/post/list' }"
                ><a href="/">帖子列表</a></el-breadcrumb-item
              >
            </el-breadcrumb>
          </el-row>
          <router-view />
        </el-main>
      </el-row>
    </el-container>
  </el-config-provider>
</template>

<script setup>
import { useDark, useToggle } from "@vueuse/core";
import { computed } from "vue";

const isDark = useDark();
const toggleDark = useToggle(isDark);
const isLoggedIn = computed(() => !!localStorage.getItem("username"));
const username = computed(() => localStorage.getItem("username") || "用户");
const handleLogout = () => {
  localStorage.removeItem("username");
  window.location.reload();
};
</script>

<style >
.el-popper > span:hover {
  cursor: pointer;
}
</style>>
