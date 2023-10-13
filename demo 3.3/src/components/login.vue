<script setup lang="ts">
import { ref, h } from "vue";
import loginStore from '../stores/loginStore';
import { ElNotification } from "element-plus";
import userCondition from '../stores/userStore'
import userService from '../apis/begin/userService'
import router from '../routers'
import type { TabsPaneContext } from "element-plus";

const Email = ref("");
const Password = ref("")
const nowLoginStore = loginStore()
const newUserStore = userCondition();
const key = ref("")
const activeName = ref('first');

const login = async () => {
    if (Email.value === "" || Password.value === "") {
        ElNotification({
            title: "登录失败",
            message: h("i", { style: "color: teal" }, "你的账号或密码为空"),
            position: "bottom-right",
        });
        return;
    }
    const loginInfo = ref({
        Email: Email.value,
        Password: Password.value,
    });
    const res = await userService.login(loginInfo.value);

    if (res.data.msg === "Accepted" && res.data.code === 200) {
        const responseData = res.data.data;
        const message = responseData.account + "，欢迎回来！"
        ElNotification({
            title: "登录成功",
            message: h("i", { style: "color: teal" }, message),
            position: "bottom-right",
        });

        nowLoginStore.setlog(true);
        console.log("登录状态" + nowLoginStore.loginSession);
        newUserStore.setUserInfo({
            Email: responseData.Email,
            Username: responseData.Username,
            Password: responseData.Password,
            Type: responseData.Type,
            TeamID: responseData.TeamID,
            Sign: responseData.Sign,
            Description: responseData.Description,
            AvataraUrl: responseData.AvataraUrl
            TelePhone: responseData.TelePhone
        });
        router.push("/Team")
        console.log(res.data)
        if (res.data.data.TeamID != "") {
            router.push("/Team")
        }
        else {
            router.push("/Choice")
        }
    }
    else if (res.data.msg === "Wrong:UserID" && res.data.code === 401) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "用户不存在"),
            type: "error"
        });
        return;
    }
    else if (res.data.msg === "Wrong:Password" && res.data.code === 402) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "密码错误"),
        });
        return;
    }
};
const clear = () => {
    Email.value = "";
    Password.value = "";
    key.value = "";
};

const push = () => {
    router.push("/register")
}

const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}

</script>

<template>
    <div class="div">
        <el_cotainer class="title">
            <el_header>登录</el_header>
        </el_cotainer>
        <br />
        <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="用户登录" name="first">
                <el-container>
                    <label class="label">Email&nbsp;&nbsp;&nbsp;</label>
                    <el-input v-model="Email" placeholder="请输入您的邮箱" class="input" />
                </el-container>
                <br />
                <el-container>
                    <label class="label">Password</label>
                    <el-input v-model="Password" type="Password" placeholder="请输入您的密码" show-Password class="input" />
                </el-container>
            </el-tab-pane>
            <el-tab-pane label="管理员登录" name="second">
                <el-container>
                    <label class="label">input&nbsp;&nbsp;&nbsp;</label>
                    <el-input v-model="Email" placeholder="请输入您的邮箱" class="input" />
                </el-container>
                <br />
                <el-container>
                    <label class="label">Password</label>
                    <el-input v-model="Password" type="Password" placeholder="请输入您的密码" show-Password class="input" />
                </el-container>
                <br />
                <el-container>
                    <label class="label">key</label>
                    <el-input v-model="key" placeholder="请输入管理员密钥" show-Password class="input" />
                </el-container>
            </el-tab-pane>
        </el-tabs>
        <br />
        <el-button type="primary" @click="login" class="button">登录</el-button>
        <el-button @click="clear" class="button">清空</el-button>
        <el-button type="plain" @click="push" class="register">注册</el-button>
    </div>
</template>
<style scoped>
.test {
    font-weight: bold;
    color: #313855;
}

.div {
    height: 100%;
    width: 300px;
    z-index: 0;
    color: #313855;
    font-size: 16px;
    font-weight: bold;
    margin-top: 50px;
}

.title {
    font-size: x-large;
    font-weight: bold;
    color: #313855;
}
</style>