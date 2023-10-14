<script setup lang="ts">
import userCondition from '../stores/userStore';
import userService from '../apis/begin/userService';
import { reactive, ref, h } from 'vue';
import { ElNotification } from 'element-plus';

const NewUserStore = userCondition()
const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign,
    AvataraUrl: NewUserStore.userSession.AvataraUrl,
})

let visible = ref(true)
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
    visible.value = !visible.value //切换显示修改后数据
}
</script>

<template>
<div>
    <el-card class="card" shadow="always">
        <el-row class="head" :span="4">
            <el-avatar id="avatar" shape="circle" :size="207" fit="filled" src=Personcard.AvataraUrl />
            <el-row id="name">
                <p id="nametext" v-if="visible">{{ PersonCard.UserName }}</p>
                <el-input id="input" v-model="PersonCard.UserName" v-if="!visible" />
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
.card{
    width: 1280px;
    height: 640px;
    margin-top: 10%;
    background-color:rgb(252, 252, 252);
}
.head{
    position: fixed;
    margin-top: 25px;
}
#name{
    margin-left: 40px;
}
#nametext{
    font-size: 45px;
    font-weight: 514;
}
#input{
    position: flex;
    height: 50px;
    margin-top: 25px;
    z-index: 50;
}
#avatar{
    margin-top: 25px;
}
.change{
    margin-left: 95%;
}
</style>