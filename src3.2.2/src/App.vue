<script setup lang="ts">
import { RouterView } from 'vue-router';
localStorage.setItem("login", String(true));
import myAside from './components/aside.vue';
import myFooter from './components/footer.vue';
</script>

<template>
  <div class="container">
    <div id="background"></div>
    <el-aside width="150px" class="el-aside">
      <my-aside></my-aside>
    </el-aside>
    <el-container class="view">
      <el-main><router-view /></el-main>
    </el-container>
    <my-footer></my-footer>
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
#background{
  background: url("/background.png");
  opacity: 0.75;
  position: fixed;
  width: 100%;			
  height: 100%;			
  position: fixed;
  background-size: cover;
  z-index: -5;
  bottom: 80px;
  margin-left: 0;
}
</style>
