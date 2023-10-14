<script setup lang="ts">
import { RouterView } from 'vue-router';
import myAside from './components/aside.vue';
import Footer from './components/footer.vue';
import Header from './components/header2.vue'

localStorage.setItem("login", String(true));
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
    <Header></Header>
    <Footer></Footer>
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

#background {
  background: url("/background.png");
  opacity: 0.7;
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
