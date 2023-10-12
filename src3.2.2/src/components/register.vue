<template>
    <div class="div">
        <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm">
            <h2 class="title">注册</h2>
            <el-form-item label="用户名" prop="UserName">
                <el-input v-model="ruleForm.UserName" type="UserName" autocomplete="off" />
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
                <el-input v-model="ruleForm.checkPass" type="password" autocomplete="off" />
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
                <el-input v-model="ruleForm.email" />
            </el-form-item>
            <el-form-item label="密钥" prop="key">
                <el-input v-model="ruleForm.key" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm(ruleFormRef)">注册</el-button>
                <el-button @click="resetForm(ruleFormRef)">清空</el-button>
            </el-form-item>
            <el-button id="back" @click="Pushback">返回</el-button>
        </el-form>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref, h } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { ElNotification } from "element-plus";
import loginStore from '../stores/loginStore';
import userCondition from '../stores/userStore'
import userService from '../apis/begin/userService'
import router from '../routers'

const nowLoginStore = loginStore()
const newUserStore = userCondition();

const ruleForm = reactive({
    UserName: '',
    pass: '',
    checkPass: '',
    email: '',
    key: ''
})

const myreg = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;

const ruleFormRef = ref<FormInstance>()

const Pushback = () => {
    router.push("/Login")
}
const checkEmail = (_rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('Please input the Email'))
    } else if (!myreg.test(ruleForm.email)) {
        callback(new Error("The email format is incorrect"))
    } else {
        if (ruleForm.email !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('checkEmail', () => null)
        }
        callback()
    }
}

const validatePass = (_rule: any, value: any, callback: any) => {
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
const validatePass2 = (_rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('Please input the password again'))
    } else if (value !== ruleForm.pass) {
        callback(new Error("Two inputs don't match!"))
    } else {
        callback()
    }
}

const validatePass3 = (_rule: any, value: any, callback: any) => {
    if (value === '') {
        callback(new Error('Please input the name'))
    } else {
        if (ruleForm.UserName !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('UserName', () => null)
        }
        callback()
    }
}

const rules = reactive<FormRules<typeof ruleForm>>({
    pass: [{ validator: validatePass, trigger: 'blur' }],
    checkPass: [{ validator: validatePass2, trigger: 'blur' }],
    email: [{ validator: checkEmail, trigger: 'blur' }],
    UserName: [{ validator: validatePass3, trigger: 'blur' }]
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            console.log('submit!')
            register()
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}

const register = async () => {
    if (ruleForm.UserName === "" || ruleForm.checkPass === "" || ruleForm.email === "" || ruleForm.pass === "") {
        ElNotification({
            title: "注册失败",
            message: h("i", { style: "color: teal" }, "你的账号或密码为空"),
            position: "bottom-right",
        });
        return;
    }
    const loginInfo = ref({
        Username: ruleForm.UserName,
        Password: ruleForm.checkPass,
        Email: ruleForm.email,
        key: ruleForm.key
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
</script>

<style scoped>
.div{
    width: 100%;
    border:1px solid  transparent;
    border-radius: 15px;
    padding-right: 30px;
    background-color: rgba(255,255,255,0.5);
    /* 设置box-shadow使其有立体感 */
    box-shadow: 5px 5px 0 0  rgba(0,0,0,0.2);
    position: relative;
    left: -15px;
    }
</style>
