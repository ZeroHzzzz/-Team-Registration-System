<template>
    <div class="trigger" v-if="IsUserPage" @mouseenter="isCollapse = false" @mouseleave="isCollapse = true">
        <el-menu default-active="2" class="el-menu" :collapse="isCollapse" @open="handleOpen" @close="handleClose">
            <template #title><el-icon>
                    <House />
                </el-icon>
                <span>个人及团队</span>
            </template>
            <el-menu-item-group>
                <template #title><span>团队</span></template>
                <el-menu-item index="1-1" @click="pushToAdd"><el-icon>
                        <Plus />
                    </el-icon>
                    <span v-if="hasCollapse">创建团队</span>
                </el-menu-item>
                <el-menu-item index="1-2" @click="pushToCreate"><el-icon>
                        <CirclePlus />
                    </el-icon>
                    <span v-if="hasCollapse">加入团队</span>
                </el-menu-item>
                <el-menu-item index="1-3" @click="pushToTeam"><el-icon>
                        <HomeFilled />
                    </el-icon>
                    <span v-if="hasCollapse">团队首页</span>
                </el-menu-item>
            </el-menu-item-group>
            <el-menu-item-group>
                <template #title><span>个人</span></template>
                <el-menu-item index="1-3" @click="pushToPerson"><el-icon>
                        <UserFilled />
                    </el-icon>
                    <span v-if="hasCollapse">名片</span>
                </el-menu-item>
            </el-menu-item-group>
            <el-menu-item index="2" @click="table = true"><el-icon>
                    <Message />
                </el-icon>
                <template #title><span>消息</span></template>

            </el-menu-item>
            <el-menu-item index="3" @click="onOperator" :disabled="isDisabled" class="op"><el-icon>
                    <Avatar />
                </el-icon>
                <template #title><span>管理员</span></template>
            </el-menu-item>
            <el-menu-item index="4" @click="pushOut"><el-icon>
                    <RemoveFilled />
                </el-icon>
                <template #title><span>退出</span></template>
            </el-menu-item>
        </el-menu>
        <el-drawer v-model="table" title="下面是你收到的消息哦,少女祈祷中ing" direction="rtl" size="50%">
            <el-table :data="TableData">
                <el-table-column label="Member"></el-table-column>
                <el-table-column label="Date"></el-table-column>
                <el-table-column label="Message"></el-table-column>
            </el-table>
        </el-drawer>
    </div>
</template>

<script setup lang="ts">
import message from '../apis/user/message'
import { ref, computed, onMounted, h } from "vue";
import { useRoute } from "vue-router";
import userCondition from '../stores/userStore'
import loginStore from '../stores/loginStore'
import { storeToRefs } from "pinia";
import { ElNotification } from "element-plus";
import router from "../routers";

const route = useRoute()
const IsUserPage = computed(() => {
    return !(route.name === "Login" || route.name === "Register")
})
const table = ref(false)
const TableData = ref([])
const newUserStore = userCondition();
const newLoginStore = loginStore();
const { loginSession } = storeToRefs(newLoginStore);
const { userSession } = storeToRefs(newUserStore);
let isCollapse = ref(true)
const hasCollapse = computed(() => !isCollapse.value)
const isDisabled = ref(newUserStore.userSession.type === 4);
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}

const pushToAdd = () => {
    router.push("/Add")
}
const pushToCreate = () => {
    router.push("/Create")
}
const pushToPerson = () => {
    router.push("/Space")
}
const onOperator = () => {
    router.push("/Operator")
}
const pushToTeam = () => {
    router.push("/Team")
}
const pushOut = () => {
    localStorage.clear();
    sessionStorage.clear();
    loginSession.value = false;
    userSession.value = {
        id: -1,
        account: "未登录",
        type: 0,
        team_id: 0
    };
    router.push("/Login");
};

onMounted(async () => {
    const account = newUserStore.userSession.account;
    console.log(account);
    const res = await message.get(account);

    if (res.data.code === 200 && res.data.msg === "OK") {
        TableData.value = res.data.data.class_list;
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "获取消息失败"),
        });
    }
})
</script>

<style scoped>
.trigger {
    z-index: 75;
    left: 0;
    top: 0;
    margin-top: 250px;
    position: fixed;
    place-items: start;
}

.el-menu {
    height: max-content;
}

.op {
    bottom: 0%;
}
</style>