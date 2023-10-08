<!--在登录和注册的时候使用这一个页头-->
<!--不带任何功能的header，只是纯展示性质-->
<script setup lang="ts">
import { ElMessage, ElMessageBox } from 'element-plus'

const vivo50 = () => {
  ElMessageBox.alert('队内负责页面布局的那个前端太菜了，还没有办法放团队打赏码', '系统提示', {
    confirmButtonText: 'OK',
    callback: () => {
      ElMessage({
        type: 'success',
        message: `(T^T)水平足够之后下次再v`,
      })
    },
  });
}
</script>

<template>
  <div class="header">
    <el-popover :width="300"
      popper-style="boxshadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
      class="popover">
      <template #reference>
        <el-avatar class="test" size="large" src="/team.jpg"></el-avatar>
        <!--这个是页头内的头像-->
      </template>
      <template #default>
        <div class="conent" style="display: flex; gap: 16px; flex-direction: column;">
          <el-avatar size="large" src="/team.jpg"></el-avatar>
          <div>
            <p class="content_title" style="margin-top: 0 font-weight:500;"><strong>健康作息小组</strong></p>
            <p class="content_ntitle" style="margin: 0 font-size:14px; color:var(--el-color-info)">健康と休息の時間グループ</p>
          </div>
          <p class="content_sub" style="margin: 0;">逸一时，误一世。以健康的作息，造就奢华的人生体验——<br />欸，等下，我趣怎么凌晨三点半了？？</p>
          <el-button text type="primary" @click="vivo50">点击此处v我们50支持一下</el-button>
        </div>
      </template>
    </el-popover>
    <header>
      <h1 class="title">&nbsp;&nbsp;团队报名系统</h1>
      <h2 class="title">&nbsp;&nbsp;&nbsp;&nbsp;//--v.1.14.514</h2>
    </header>
  </div>
</template>

<style scoped>
.test {
  position: absolute;
  margin-top: 75px;
  margin-left: 90%;
}

.header {
  position: fixed;
  top: 0px;
  left: 0px;
  right: 0px;
  background-color: #a3a3a3;
  z-index: 50;
}

.title {
  color: #ededed;
}
</style>