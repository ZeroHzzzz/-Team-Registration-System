<script setup lang="ts">
import { ref, onMounted, h, nextTick } from 'vue'
import team from '../apis/team/team'
import userCondition from "../stores/userStore";
import { ElNotification } from "element-plus";
import leader from '../apis/team/leader'

const NewUserStore = userCondition();
const TableData = ref([])

onMounted(async () => {
    const TeamId = NewUserStore.userSession.team_id
    const Account = NewUserStore.userSession.account
    const res = await team.get(Account, TeamId);

    if (res.data.code === 200 && res.data.msg === "ok") {
        TableData.value = res.data.data.class_list;
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:firebrick" }, "找寻团队失败"),
            type: "error",
        });
    }
});

const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

const handleDelete = async (id: number, name: string, index: number) => {
    const res = await leader.delete({ account: name, team_id: id });
    if (res.data.code === 200) {
        TableData.value.splice(index, 1);
        ElNotification({
            title: "成功",
            message: h("i", { style: "color: teal" }, "踢出成功！"),
        });

        const account = NewUserStore.userSession.account;
        const TeamId = NewUserStore.userSession.team_id
        console.log(account);
        const res = await team.get(account, TeamId);

        if (res.data.code === 200 && res.data.msg === "ok") {
            TableData.value = res.data.data.c;
        }
        else {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color: teal" }, "网络错误！"),
            });
        }
        await forceRender();
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color: teal" }, "删除失败！"),
        });
    }
};

const AddMember = async () => {
    if (form.team_id !== 0) {
        const res = await leader.add(form);

        if (res.data.code === 200 && res.data.msg === "OK") {
            console.log(form);
            ElNotification({
                title: "成功",
                message: h("i", { style: "color:teal" }, "添加成员成功！"),
            });
            clear();
        }
        else {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color:teal" }, "添加失败，网络错误"),
            });
        }
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "加入成员，请输入正确的用户名！"),
        });
    }
};

</script>

<template>
    <div class="div">
        <el-table :data="TableData" height="500" style="width: 100%" border>
            <el-table-column prop="TeamId" label="团队" width="1000" align="center" show-overflow-tooltip />
            <el-table-column label="操作" align="center">
                <template #default="scope">
                    <el-button type="danger"
                        @click="handleDelete(scope.row.account, scope.row.TeamId, scope.$index)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-button style="width: 100%" @click="AddMember"><span>添加成员</span></el-button>
    </div>
</template>

<style scoped>
.div {
    width: 100%;
}
</style>
