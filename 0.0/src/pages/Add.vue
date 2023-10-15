<script setup lang="ts">
import add from "../components/add.vue";
</script>


<template>
    <div class="wrapper">
        <div id="add">
            <add></add>
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

#add {
    position: absolute;
    margin-left: 24%;
    top: 40%;
}
</style>