<script setup lang="ts">
import team from '../components/team.vue';
import header2 from '../components/header2.vue'
</script>
<template>
    <header2 />
    <el-container>
        <el-col :span="100">
            <team />
        </el-col>
    </el-container>
</template>

<style scoped></style>