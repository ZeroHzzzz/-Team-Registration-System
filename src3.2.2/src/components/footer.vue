<template>
    <div class="footer">
        <el-link type="info" :underline="false" class="link" href="https://github.com/JustaGreenhand/-Team-Registration-System" target="_blank">Github</el-link>
        <el-divider class="divider" direction="vertical" />
        <el-tooltip class="box-item" effect="dark" placement="right-start">
            <template #content>前端：目猫/salt<br />后端：黑白</template>
            <el-button link type="info" class="about" @click="about">关于我们</el-button>
        </el-tooltip>
    </div>
</template>

<script setup lang="ts">
import { ElMessageBox } from 'element-plus';

const about = () => {
  ElMessageBox.alert('详情请见左侧源码仓库内readme介绍>-<', '系统提示', {
    confirmButtonText: 'OK',
  });
}
</script>
  
<style scoped>
.footer {
    position: fixed;
    bottom: 0;
    height: 80px;
    left: 0px;
    right: 0px;
    background-color: #394467;
    z-index: 80;
}
.link{
    font-size: large;
    margin-top: 30px;
    margin-left: 45%;
}
.about{
    font-size: large;
    margin-left: 15px;
    margin-top: 30px;
}
.divider{
    z-index: 100;
    position: fixed;
    color:aliceblue;
    bottom: 30px;
}
</style>