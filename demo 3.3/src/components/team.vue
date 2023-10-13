<script setup lang="ts">
import { ref, onMounted, h, nextTick, reactive } from "vue"
import team from '../apis/team/team'
import userCondition from "../stores/userStore";
import { ElNotification, ElMessageBox } from "element-plus";
import leader from '../apis/team/leader'

const NewUserStore = userCondition();
const TableData = ref([]);
const form = reactive({
    Email: NewUserStore.userSession.Email,
    TeamID: NewUserStore.userSession.TeamID,
    Username: NewUserStore.userSession.Username,
})

onMounted(async () => {
    const TeamID = NewUserStore.userSession.TeamID
    const Username = NewUserStore.userSession.Username
    const res = await team.get(Username, TeamID);

    if (res.data.code === 200 && res.data.msg === "ok") {
        TableData.value = res.data.data.Team;
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "找寻团队失败"),
        });
    }
});

const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

const handleDelete = async (id: string, name: string, index: number) => {
    const res = await leader.delete({ Email: name, TeamID: id });
    if (res.data.code === 200) {
        TableData.value.splice(index, 1);
        ElNotification({
            title: "成功",
            message: h("i", { style: "color: teal" }, "踢出成功！"),
        });

        const account = NewUserStore.userSession.Username;
        const TeamID = NewUserStore.userSession.TeamID
        console.log(account);
        const res = await team.get(account, TeamID);

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
    if (form.TeamID !== "") {
        const res = await leader.add(form);

        if (res.data.code === 200 && res.data.msg === "OK") {
            console.log(form);
            ElNotification({
                title: "成功",
                message: h("i", { style: "color:teal" }, "添加成员成功！"),
            });
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

const open = () => {
    ElMessageBox.prompt('添加团队成员，请输入用户名', 'Tip', {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
    })
        .then(({ value }) => {
            form.Username = value
            AddMember()
        })
}
</script>

<template>
    <div class="div">
        <el-table :data="TableData" height="400" style="width: 100%" border>
            <el-table-column prop="TeamId" label="团队" width="1000" align="center" show-overflow-tooltip />
            <el-table-column label="操作" align="center">
                <template #default="scope">
                    <el-button type="danger"
                        @click="handleDelete(scope.row.account, scope.row.TeamId, scope.$index)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-button style="width: 100%" @click="open"><span>添加成员</span></el-button>
    </div>
</template>

<style scoped>
.div {
    width: 100%;
}
</style>
