<script setup lang="ts">
import login from "../components/login.vue";
import myShow from "../components/show.vue";
import header1 from '../components/header1.vue'
</script>

<template>
    <div class="wrapper">
        <div id="login">
            <login></login>
        </div>
        <div id="show">
            <myShow></myShow>
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

#login {
    position: absolute;
}

#show {
    position: relative;
    margin-left: 600px;
}
</style>