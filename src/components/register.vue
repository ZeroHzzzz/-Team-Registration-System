<script setup lang="ts">
import { ref, h, computed, reactive } from "vue";
import loginStore from '../stores/loginStore';
import { ElNotification } from "element-plus";
import userCondition from '../stores/userStore'
import userService from '../apis/begin/userService'
import router from '../routers'

const input = ref("");
const password = ref("")
const key = ref("")
const email = ref("")
const nowLoginStore = loginStore()
const newUserStore = userCondition();

const RegisterForm = reactive({
    email: "",
    input: "",
    password: "",
    checkPassword: "",
    key: "",
})
const register = async () => {
    if (input.value === "" || password.value === "") {
        ElNotification({
            title: "注册失败",
            message: h("i", { style: "color: teal" }, "你的账号或密码为空"),
            position: "bottom-right",
        });
        return;
    }
    const loginInfo = ref({
        account: input.value,
        password: password.value,
    });
    const res = await userService.register(loginInfo.value);

    if (res.data.msg === "OK" && res.data.code === 200) {
        const responseData = res.data.data;
        const message = responseData.account + "，欢迎回来！"
        ElNotification({
            title: "注册成功",
            message: h("i", { style: "color: teal" }, message),
            position: "bottom-right",
        });

        nowLoginStore.setlog(true);
        console.log("登录状态" + nowLoginStore.loginSession);
        localStorage.setItem("login", String(true));
        localStorage.setItem("name", String(responseData.account));
        newUserStore.setUserInfo({
            user_id: responseData.user_id,
            account: responseData.account,
            password: responseData.password,
            type: responseData.type,
            team_id: responseData.team_id,
        });
        console.log(res.data)
        if (res.data.data.team_id !== 0) {
            router.push("/team")
        }
        else {
            router.push("/choice")
        }
    }
    else if (res.data.msg === "账号不存在" && res.data.code === 401) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "用户不存在"),
        });
        return;
    }
    else if (res.data.msg === "密码错误" && res.data.code === 402) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "密码错误"),
        });
        return;
    }
}

const clear = () => {
    email.value = ""
    input.value = "";
    password.value = "";
    key.value = "";
};


const myreg = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;

const emailtest = computed(() => {
    return (myreg.test(email.value) || email.value === "");
})


const validatePass = (rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('Please input the password'))
    } else {
        if (ruleForm.checkPass !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('checkPass', () => null)
        }
        callback()
    }
}
</script>

<template>
    <div class="div">
        <el_cotainer class="title">
            <el_header>注册</el_header>
        </el_cotainer>
        <br />
        <template>
            <el-form ref="ruleFormRef" :model="RegisterForm" status-icon :rules="rules" label-width="120px"
                class="demo-ruleForm">
                <el-form-item label="Email" prop="pass">
                    <label class="label">Email&nbsp;&nbsp;&nbsp;</label>
                    <el-input v-model="RegisterForm.email" placeholder="请输入您的邮箱" autocomplete="off" />
                </el-form-item>
                <el-form-item label="UerID" prop="pass">
                    <label class="label">Account&nbsp;&nbsp;&nbsp;</label>
                    <el-input v-model="RegisterForm.input" placeholder="请输入您的用户名" autocomplete="off" />
                </el-form-item>
                <el-form-item label="Password" prop="pass">
                    <label class="label">Password&nbsp;</label>
                    <el-input v-model="RegisterForm.password" type="password" autocomplete="off" />
                </el-form-item>
                <el-form-item label="Confirm" prop="checkPass">
                    <label class="label">Confirm&nbsp;</label>
                    <el-input v-model="RegisterForm.checkPassword" type="password" autocomplete="off" />
                </el-form-item>
                <el-form-item label="Key" prop="key">
                    <label class="label">key&nbsp;&nbsp;</label>
                    <el-input v-model="RegisterForm.key" type="key" auto-complete="off" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="submitForm(ruleFormRef)">Submit</el-button>
                    <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
                </el-form-item>
            </el-form>
        </template>
        <el-tab-pane label="用户注册" name="first">
            <el-container>
                <label class="label">Email&nbsp;&nbsp;&nbsp;</label>
                <el-input v-model="email" placeholder="请输入您的邮箱" class="email" />
            </el-container>
            <p v-if="!emailtest" class="errorword">邮箱地址格式不正确</p>
            <br>
            <el-container>
                <label class="label">Account&nbsp;&nbsp;&nbsp;</label>
                <el-input v-model="input" placeholder="请输入您的用户名" class="input" />
            </el-container>
            <br />
            <el-container>
                <label class="label">Password&nbsp;</label>
                <el-input v-model="password" type="password" placeholder="请输入您的密码" show-password class="input" />
            </el-container>
            <br />
            <el-container>
                <label class="label">key&nbsp;&nbsp;</label>
                <el-input v-model="key" placeholder="请输入系统发放的密钥" show-password class="input" />
            </el-container>
        </el-tab-pane>
        <br />
        <el-button type="primary" @click="register" class="button">注册</el-button>
        <el-button @click="clear" class="button">清空</el-button>
    </div>
</template>
<style scoped>
.div {
    height: 100%;
    width: 300px;
    z-index: 0;
}

.title {
    font-size: x-large;
    font-weight: bold;
}

.errorword {
    font-size: xx-small;
    color: red;
    height: 0;
}
</style>