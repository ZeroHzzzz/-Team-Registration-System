<script setup lang="ts">
import operator from "../components/operator.vue";
import header1 from '../components/header1.vue'
</script>

<template>
    <div class="wrapper">
        <div id="login">
            <operator></operator>
        </div>
        <header1></header1>
        <Footer1></Footer1>
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
}

#show {
    position: relative;
    margin-left: 600px;
}
</style>
