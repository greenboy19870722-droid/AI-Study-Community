<template>
  <el-config-provider>
    <el-container>
      <el-header> </el-header>

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
              <el-breadcrumb-item :to="{ path: '/home' }"
                >首页</el-breadcrumb-item
              >
              <el-breadcrumb-item><a href="/">帖子列表</a></el-breadcrumb-item>
              <el-breadcrumb-item>活动详情</el-breadcrumb-item>
            </el-breadcrumb>

            <el-space size="medium">
              <el-text tag="b">主题切换</el-text>
              <el-switch v-model="isDark" inline-prompt @change="toggleDark" />
            </el-space>
          </el-row>
          <router-view />
        </el-main>
      </el-row>
    </el-container>
  </el-config-provider>
</template>

<script setup>
import { useDark, useToggle } from "@vueuse/core";

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>

<style>
</style>