<template>
    <el-carousel :interval="5000" class="show" type="card">
        <el-carousel-item v-for="item in imageList" :key="item.id" class="el-carousel__item">
            <img :src="item.url" alt="无图片">
        </el-carousel-item>
    </el-carousel>
</template>

<script setup lang="ts">
import { ElCarousel, ElCarouselItem } from 'element-plus';
import { ref } from "vue";

const imageList = ref([
    { id: 1, url: '/sample1.jpg' },
    { id: 2, url: '/sample2.png' },
    { id: 3, url: '/sample3.jpg' },
    { id: 4, url: '/sample4.png' }
]); 
</script>

<style scoped>
.show {
    width: 800px;
    max-height: 600px;
    z-index: 50;
    margin-top: 50px;
}

.el-carousel__item {
    opacity: 0.95;
    line-height: 300px;
    margin: 0;
    text-align: center;
}

.el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
    background-color: #d3dce6;
}

img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}
</style>