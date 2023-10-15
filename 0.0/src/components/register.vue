<template>
    <div class="div">
        <div class="card">
            <el_cotainer class="title">
                <el_header>注册</el_header>
            </el_cotainer>
            <br />
            <br />
            <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="120px">
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
                    <el-button @click="Pushback">返回</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { reactive, ref, h } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { ElNotification } from "element-plus";
import loginStore from '../stores/loginStore';
import userService from '../apis/begin/userService'
import router from '../routers'

const nowLoginStore = loginStore()

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

    if (res.data.msg === "Accepted" && res.data.code === 200) {
        ElNotification({
            title: "注册成功",
            message: h("i", { style: "color: teal" },),
            position: "bottom-right",
        });
        console.log("登录状态" + nowLoginStore.loginSession);
        console.log(res.data)
        router.push("/Login")
    }
    else if (res.data.msg === "邮箱格式错误" && res.data.code === 401) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "邮箱格式错误"),
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
    else if (res.data.msg === "参数错误" && res.data.code === 400) {
        ElNotification({
            title: "登陆失败",
            message: h("i", { style: "color:teal" }, "参数错误"),
        });
        return;
    }
}
</script>

<style scoped>
.card {
    width: 120%;
    height: 400px;
    border: 1px solid transparent;
    border-radius: 15px;
    padding-right: 30px;
    background-color: rgba(255, 255, 255, 0.5);
    box-shadow: 5px 5px 0 0 rgba(0, 0, 0, 0.2);
    position: relative;
    padding: 1em;
    right: 15px;
}

.div {
    color: #313855;
    font-size: 16px;
    font-weight: bold;
}

.title {
    margin-left: 65px;
    font-size: 35px;
    font-weight: bold;
    color: #313855;
}
</style>
