<script setup lang="ts">
import router from '../routers';

const addClick = () => {
    router.push("/Add")
}
const createClick = () => {
    router.push("/Create")
}
</script>

<template>
    <div>
        <el-button type="primary" @click="addClick">创建自己的团队</el-button>
        <br>
        <el-button type="success" @click="createClick">加入别人的团队</el-button>
    </div>
</template>