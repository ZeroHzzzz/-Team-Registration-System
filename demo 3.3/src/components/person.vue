<script setup lang="ts">
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, h, reactive } from "vue";
import userCondition from '../stores/userStore'
const NewUserStore = userCondition()
const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign,
    AvataraUrl: NewUserStore.userSession.AvataraUrl
})
//弹窗卡片
const open = () => {
    ElMessageBox({
        title: '个人信息',
        message: h('p', null, [
            h('img', { attrs: { src: PersonCard.AvataraUrl } }),
            h('h1', null, PersonCard.UserName),
            h('h3', null, PersonCard.Sign),
            h('h3', null, PersonCard.Description)
        ]),

    })
}

</script>

<template></template>