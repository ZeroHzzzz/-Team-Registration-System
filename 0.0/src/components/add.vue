<script setup lang="ts">
import { reactive, h } from "vue";
import member from '../apis/team/member';
import userCondition from '../stores/userStore';
import { ElNotification } from "element-plus";


const newUserStore = userCondition()

const form = reactive({
    Email: newUserStore.userSession.Email,
    TeamID: "",
    TeamPassword: "",
});

const onSubmit = async () => {
    if (form.TeamID !== "") {
        const res = await member.add(form);

        if (res.data.code === 200 && res.data.msg === "Accepted") {
            console.log(form);
            ElNotification({
                title: "成功",
                message: h("i", { style: "color:teal" }, "加入团队成功！"),
            });
            clear();
        }
        else {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color:teal" }, "加入团队，网络错误"),
            });
        }
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "加入团队，请输入正确的团队号！"),
        });
    }
};
const clear = () => {
    form.TeamID = "";
    form.TeamPassword = ""
}


</script>
<template>
    <div class="div">
        <div class="card">
            <div class="title">
                <h2>加入团队</h2>
            </div>
            <el-form :model="form" label-width="120px">
                <el-form-item label="团队ID">
                    <el-input v-model="form.TeamID" placeholder="请输入团队ID" />
                </el-form-item>
                <el-form-item label="邀请码">
                    <el-input v-model="form.TeamPassword" placeholder="请输入团队邀请码" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit" class="button">加入</el-button>
                    <el-button class="button" @click="clear">清空</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<style scoped>
.card {
    width: 400px;
    height: 250px;
    border: 1px solid transparent;
    border-radius: 15px;
    padding-right: 30px;
    background-color: rgba(255, 255, 255, 0.5);
    box-shadow: 5px 5px 0 0 rgba(0, 0, 0, 0.2);
    position: relative;
    padding: 1em;
}

.title {
    margin-left: 65px;
    color: #313855;
}
</style>