<script setup lang="ts">
import userCondition from '../stores/userStore';
import userService from '../apis/begin/userService';
import { reactive, ref, h, nextTick } from 'vue';
import { ElNotification } from 'element-plus';

const NewUserStore = userCondition()
const PersonCard = reactive({
    Username: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Description,
    Email: NewUserStore.userSession.Email,
    AvataraUrl: NewUserStore.userSession.AvataraUrl,
    TelePhone: NewUserStore.userSession.TelePhone
})
const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

let visible = ref(true)
const submit = async () => {
    console.log("1")
    const res = await userService.update(PersonCard)

    if (res.data.msg === "Accepted" && res.data.code === 200) {
        ElNotification({
            title: "系统提示",
            message: h("i", { style: "color: teal" }, "修改成功")
        })
        NewUserStore.setUserInfo({
            Email: PersonCard.Email,
            Username: PersonCard.Username,
            Type: NewUserStore.userSession.Type,
            TeamID: NewUserStore.userSession.TeamID,
            Sign: PersonCard.Sign,
            Description: PersonCard.Description,
            AvataraUrl: NewUserStore.userSession.AvataraUrl,
            TelePhone: PersonCard.TelePhone
        });
        forceRender()
    }
    console.log("提交成功")
    visible.value = !visible.value //切换显示修改后数据
}
</script>

<template>
    <div>
        <el-card class="card" shadow="always">
            <el-row class="head" :span="4">
                <el-avatar id="avatar" shape="circle" :size="207" fit="filled" :src=PersonCard.AvataraUrl />
                <el-row id="name">
                    <p id="nametext" v-if="visible">{{ PersonCard.Username }}</p>
                    <el-input id="input" v-model="PersonCard.Username" v-if="!visible" />
                </el-row>
                <el-row id="sign">
                    <p id="signtext" v-if="visible">{{ PersonCard.Sign }}</p>
                    <el-input id="input" v-model="PersonCard.Sign" v-if="!visible" />
                </el-row>
                <el-row id="description">
                    <p id="descriptiontext" v-if="visible">{{ PersonCard.Description }}</p>
                    <el-input id="input" v-model="PersonCard.Description" v-if="!visible" />
                </el-row>
            </el-row>
            <div class="change">
                <el-button text @click="visible = !visible" v-if="visible">编辑</el-button>
                <el-button text @click="submit" v-if="!visible">提交</el-button>
            </div>
        </el-card>
    </div>
</template>

<style scoped>
.card {
    width: 1280px;
    height: 640px;
    margin-top: 10%;
    background-color: rgb(252, 252, 252);
    margin-left: -200px;
}

.head {
    position: fixed;
    margin-top: 25px;
}

#name {
    margin-left: 40px;
}

#nametext {
    font-size: 45px;
    font-weight: 514;
}

#input {
    position: flex;
    height: 50px;
    margin-top: 25px;
    z-index: 50;
}

#avatar {
    margin-top: 25px;
}

.change {
    margin-left: 95%;
}
</style>