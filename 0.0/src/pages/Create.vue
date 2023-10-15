<script setup lang="ts">
import create from "../components/create.vue";
</script>

<template>
    <div class="wrapper">
        <div id="create">
            <create></create>
        </div>
    </div>
</template>

<style scoped>
.wrapper {
    display: flex;
}

.wrapper>div {
    flex: 1;
}

#create {
    position: absolute;
    margin-left: 24%;
    top: 40%;
}
</style>