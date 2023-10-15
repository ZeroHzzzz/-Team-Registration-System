<script setup lang="ts">
import register from "../components/register.vue";
import header1 from '../components/header1.vue'
</script>

<template>
    <div class="wrapper">
        <div id="register">
            <register></register>
        </div>
        <header1></header1>
    </div>
</template>

<style scoped>
.wrapper {
    display: flex;
}

.wrapper>div {
    flex: 1;
}

#register {
    position: absolute;
    margin-left: 24%;
    bottom: 20%;
}

#show {
    position: relative;
    margin-left: 600px;
}
</style>