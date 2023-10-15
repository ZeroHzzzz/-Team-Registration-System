<script setup lang="ts">
import { ref, h, nextTick } from "vue";
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

//初始化
localStorage.setItem("login", String(false));


const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

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
        const message = responseData.returnfile.Username + "，欢迎回来！"
        ElNotification({
            title: "登录成功",
            message: h("i", { style: "color: teal" }, message),
            position: "bottom-right",
        });

        nowLoginStore.setlog(true);
        console.log("登录状态" + nowLoginStore.loginSession);
        newUserStore.setUserInfo({
            Email: responseData.returnfile.Email,
            Username: responseData.returnfile.Username,
            Password: responseData.returnfile.Password,
            Type: responseData.returnfile.Type,
            TeamID: responseData.returnfile.TeamID,
            Sign: responseData.returnfile.Sign,
            Description: responseData.returnfile.Description,
            AvataraUrl: responseData.returnfile.AvataraUrl,
            TelePhone: responseData.returnfile.TelePhone
        });
        localStorage.setItem("login", String(true));
        localStorage.setItem("token", responseData.token);
        console.log(res.data)
        forceRender()
        if (newUserStore.userSession.TeamID = null) {
            router.push("/Choice")
        }
        else {
            router.push("/Person")
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
    <div class="card">
        <div class="div">
            <el_cotainer class="title">
                <el_header>登录</el_header>
            </el_cotainer>
            <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane name="first">
                    <template #label><span class="test">用户登录</span></template>
                    <el-container>
                        <label class="label">Email&nbsp;&nbsp;&nbsp;</label>
                        <el-input v-model="Email" placeholder="请输入您的邮箱" class="input" />
                    </el-container>
                    <br />
                    <el-container>
                        <label class="label">Password&nbsp;&nbsp;</label>
                        <el-input v-model="Password" type="Password" placeholder="请输入您的密码" show-Password class="input" />
                    </el-container>
                </el-tab-pane>
                <el-tab-pane name="second">
                    <template #label><span class="test">管理员登录</span></template>
                    <el-container>
                        <label class="label">Email&nbsp;&nbsp;&nbsp;</label>
                        <el-input v-model="Email" placeholder="请输入您的邮箱" class="input" />
                    </el-container>
                    <br />
                    <el-container>
                        <label class="label">Password&nbsp;&nbsp;</label>
                        <el-input v-model="Password" type="Password" placeholder="请输入您的密码" show-Password class="input" />
                    </el-container>
                    <br />
                    <el-container>
                        <label class="label">key&nbsp;&nbsp;</label>
                        <el-input v-model="key" placeholder="请输入管理员密钥" show-Password class="input" />
                    </el-container>
                </el-tab-pane>
            </el-tabs>
            <br />
            <el-button type="primary" @click="login" class="button">登录</el-button>
            <el-button @click="clear" class="button">清空</el-button>
            <el-button type="plain" @click="push" class="register">注册</el-button>
        </div>
    </div>
</template>
<style scoped>
.card {
    width: 100%;
    height: 310px;
    border: 1px solid transparent;
    border-radius: 15px;
    padding-right: 30px;
    background-color: rgba(255, 255, 255, 0.5);
    /* 设置box-shadow使其有立体感 */
    box-shadow: 5px 5px 0 0 rgba(0, 0, 0, 0.2);
    position: flex;
    padding: 1em;
}

.test {
    font-weight: bold;
    color: #313855;
}

.div {
    margin-left: 5%;
    height: 100%;
    width: 320px;
    z-index: 0;
    color: #313855;
    font-size: 16px;
    font-weight: bold;
}

.title {
    font-size: 35px;
    font-weight: bold;
    color: #313855;
}
</style>