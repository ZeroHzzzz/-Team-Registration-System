<script setup lang="ts">
import { ElMessageBox } from "element-plus";
import { h, reactive } from "vue";
import userCondition from '../stores/userStore'


const NewUserStore = userCondition()
const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign,
    AvataraUrl: NewUserStore.userSession.AvataraUrl,
})

const emit = defineEmits(["openPerson"])
//弹窗卡片
const openPerson = () => {
    ElMessageBox({
        title: '个人信息',
        message: h('p', null, [
            h('img', { attrs: { src: PersonCard.AvataraUrl } }),
            h('h1', null, PersonCard.UserName),
            h('h3', null, PersonCard.Sign),
            h('h3', null, PersonCard.Description)
        ]),
        showCancelButton: true,
        confirmButtonClass: "返回",
    })
    emit("openPerson")
}


</script>


<template>
    <button @click="openPerson"></button>
</template>