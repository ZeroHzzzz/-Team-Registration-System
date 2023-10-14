<script setup lang="ts">
import { ElMessage, ElMessageBox } from "element-plus";
import { h, reactive, ref } from "vue";
import userCondition from '../stores/userStore'
const NewUserStore = userCondition()
const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign,
    AvataraUrl: NewUserStore.userSession.AvataraUrl,
})
//弹窗卡片
const openPerson = () => {
    ElMessageBox({
        title: '个人信息',
        message: h('p', null, [
            h('img', { attrs: { src: PersonCard.AvataraUrl } }),
            h('h1',  null, PersonCard.UserName),
            h('p', null, PersonCard.Sign),
            h('h3', null, PersonCard.Description)
        ]),
        confirmButtonText: '确认',
        showCancelButton: true,
        cancelButtonText: '返回'
    })
}
//修改表单
const visible = ref(false)

const editPerson = () => {
    
}
</script>

<template>
    <el-button type="primary" @click="openPerson">我测尼玛</el-button>
    <br />
    <el-button type="warning" @click="visible=true">测不出来</el-button>
    <el-dialog v-model="visible" title="编辑信息">
        <el-form>
            <el-form-item label="用户名">
                <el-input v-model="PersonCard.UserName" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button type="primary" @click="editPerson">修改</el-button>
            <el-button @click="visible=false">取消</el-button>
        </template>
    </el-dialog>
</template>

<style scoped>
</style>