<script setup lang="ts">
import router from '../routers';

const addClick = () => {
    router.push("/Create")
}
const createClick = () => {
    router.push("/Add")
}
</script>

<template>
    <div class="create">
        <el-button type="primary" @click="addClick">创建自己的团队</el-button>

    </div>
    <div class="join">
        <el-button type="success" @click="createClick">加入别人的团队</el-button>
    </div>
</template>

<style scoped></style>