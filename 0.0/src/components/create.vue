<script setup lang="ts">
import { reactive, h } from "vue";
import leader from '../apis/team/leader';
import userCondition from '../stores/userStore';
import { ElNotification } from "element-plus";
import router from "../routers";


const newUserStore = userCondition()

const form = reactive({
    LeaderID: newUserStore.userSession.Email,
    TeamDescription: "",
    TeamName: "",
    Password: "",
});

const onSubmit = async () => {
    if (form.TeamName !== "") {
        const res = await leader.create(form);
        if (res.data.code === 200 && res.data.msg === "Accepted") {
            console.log(form);
            ElNotification({
                title: "成功",
                message: h("i", { style: "color:teal" }, "创建团队成功！"),
            });
            clear();
            router.push("/team")
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
    form.TeamName = "";
    form.Password = "";
    form.TeamDescription = "";
}


</script>
<template>
    <div class="div">
        <div class="card">
            <div class="title">
                <h2>创建团队</h2>
            </div>
            <el-form :model="form" label-width="120px">
                <el-form-item label="团队名称">
                    <el-input v-model="form.TeamName" placeholder="请输入团队名称" />
                </el-form-item>
                <el-form-item label="邀请码">
                    <el-input v-model="form.Password" placeholder="请输入团队邀请码" />
                </el-form-item>
                <el-form-item label="团队简介">
                    <el-input v-model="form.TeamDescription" placeholder="为团队写个简介吧" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="onSubmit" class="button">创建</el-button>
                    <el-button class="button" @click="clear">清空</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<style scoped>
.card {
    width: 400px;
    height: 270px;
    border: 1px solid transparent;
    border-radius: 15px;
    padding-right: 30px;
    background-color: rgba(255, 255, 255, 0.5);
    box-shadow: 5px 5px 0 0 rgba(0, 0, 0, 0.2);
    position: relative;
    padding: 1em;
}

.title {
    margin-left: 50px;
    color: #313855;
}
</style>