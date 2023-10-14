<script setup lang="ts">
import operator from "../components/operator.vue";
</script>

<template>
    <div class="wrapper">
        <div id="operator">
            <operator></operator>
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

#operator{
    position: absolute;
}
</style>
