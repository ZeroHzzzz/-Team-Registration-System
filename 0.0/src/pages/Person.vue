<script setup lang="ts">
import person from "../components/person.vue";
</script>

<template>
    <div class="wrapper">
        <div id="space">
            <person></person>
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

#space {
    position: absolute;
    margin-left: 24%;
    top: 15%;
}
</style>