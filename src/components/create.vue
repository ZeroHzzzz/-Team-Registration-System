<script setup lang="ts">
import { reactive, h } from "vue";
import member from '../apis/team/member';
import userCondition from '../stores/userStore';
import { ElNotification } from "element-plus";


const newUserStore = userCondition()

const form = reactive({
    user_id: newUserStore.userSession.id,
    team_id: 0,
    teamName: "",
    teamPassword: "",
    teamCondition: false
});

const onSubmit = async () => {
    if (form.teamName !== "") {
        const res = await member.add(form);
        if (res.data.code === 200 && res.data.msg === "OK") {
            console.log(form);
            ElNotification({
                title: "成功",
                message: h("i", { style: "color:teal" }, "创建团队成功！"),
            });
            clear();
        }
        else if (res.data.code === 2005 && res.data.msg === "已被注册") {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color:teal" }, "创建团队失败，团队名称已被注册")
            });
        }
        else {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color:teal" }, "创建团队失败，网络错误"),
            });
        }
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "创建，请输入团队名称！"),
        });
    }
};
const clear = () => {
    form.teamName = "";
    form.teamPassword = ""
}


</script>

<template>
    <div class="div">
        <div style="text-align: center ; margin-left: 120px">
            <h3>(创建)团队</h3>
        </div>
        <el-form :model="form" label-width="120px">
            <el-form-item label="团队创建">
                <el-input v-model="form.team_id" placeholder="请输入团队id" />
                <el-input v-model="form.teamPassword" placeholder="请输入邀请码"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit" class="button">加入</el-button>
                <el-button class="button" @click="clear">清空</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<style scoped></style>