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

#login {
    position: absolute;
    margin-top: 3%;
    margin-left: -5%;
}

#show {
    position: relative;
    margin-top: 3%;
    margin-left: 514px;
}
</style>