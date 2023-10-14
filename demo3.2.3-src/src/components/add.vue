<script setup lang="ts">
import { reactive, h } from "vue";
import member from '../apis/team/member';
import userCondition from '../stores/userStore';
import { ElNotification } from "element-plus";


const newUserStore = userCondition()

const form = reactive({
    Email: newUserStore.userSession.Email,
    Username: newUserStore.userSession.Username,
    TeamID: "",
    TeamName: "",
    Password: "",
    teamCondition: false
});

const onSubmit = async () => {
    if (form.TeamID !== "") {
        const res = await member.add(form);

        if (res.data.code === 200 && res.data.msg === "OK") {
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
    form.Password = ""
    form.TeamName = ""
}


</script>

<template>
    <div class="div">
        <div style="text-align: center ; margin-left: 120px">
            <h3>(加入)团队</h3>
        </div>
        <el-form :model="form" label-width="120px">
            <el-form-item label="团队加人">
                <el-input v-model="form.TeamID" placeholder="请输入团队id" />
                <el-input v-model="form.Password" placeholder="请输入邀请码"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit" class="button">加入</el-button>
                <el-button class="button" @click="clear">清空</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<style></style>