<script setup lang="ts">
import { RouterView } from 'vue-router';
localStorage.setItem("login", String(true));
import myAside from './components/aside.vue';
</script>

<template>
  <div class="container">
    <el-aside width="150px" class="el-aside">
      <my-aside></my-aside>
    </el-aside>
    <el-container class="view">
      <el-main><router-view /></el-main>
    </el-container>
  </div>
</template>

<style scoped>
.container {
  display: flex;
}

.view {
  place-items: center
}

.el-aside {
  order: -1;
  left: 0;
  top: 0;
  height: 100%;
  width: 200px;
}
</style>
