<script setup lang="ts">
import { ElMessageBox, ElNotification } from "element-plus";
import { h, reactive, ref, nextTick } from "vue";
import userCondition from '../stores/userStore';
import userService from '../apis/begin/userService';

const NewUserStore = userCondition()
const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign,
    AvataraUrl: NewUserStore.userSession.AvataraUrl,
})

const cardVisible = ref(false)
const formVisible = ref(false)
const emit = defineEmits(["openPerson"])
//弹窗卡片
/* const openPerson = () => {
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
} */
const openPerson = () => {
    cardVisible.value = !cardVisible.value
}

const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

const submit = async() => {
    console.log("1")
    const updateInfo = reactive(PersonCard)
    const res = await userService.update(updateInfo)

    if (res.data.msg === "Accepted" && res.data.code === 200) {
        ElNotification({
            title:"系统提示",
            message: h("i", { style :"color: teal" },"修改成功")
        })
    }
    console.log("提交成功")
    forceRender();
}
const editPerson = () => {
    formVisible.value = !formVisible.value
}
</script>


<template>
    <el-button type="primary" @click="openPerson">测测你的</el-button>
    <el-dialog v-model="cardVisible">
        <el-avatar id="avatar" shape="circle" :size="207" fit="filled" src=Personcard.AvataraUrl />
        <p id="nametext">{{ PersonCard.UserName }}</p>
        <p id="signtext">{{ PersonCard.Sign }}</p>
        <p id="descriptiontext">{{ PersonCard.Description }}</p>
        <el-button type="primary" @click="editPerson">让我改改</el-button>
        <el-dialog v-model="formVisible" title="个人资料" append-to-body>
            <el-form v-model="formVisible">
                <el-form-item label="用户名">
                    <el-input v-model="PersonCard.UserName" autocomplete="off" />
                </el-form-item>
                <el-form-item label="个性签名">
                    <el-input v-model="PersonCard.Sign" autocomplete="off" />
                </el-form-item>
                <el-form-item label="个人简介">
                    <el-input v-model="PersonCard.Description" autocomplete="off" />
                </el-form-item>
                <span class="dialog-footer">
                    <el-button @click="formVisible=false">我靠等下</el-button>
                    <el-button type="primary" @click="submit">tmd,交!</el-button>
                </span>
            </el-form>
        </el-dialog>
    </el-dialog>
</template>