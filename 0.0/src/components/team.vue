<script setup lang="ts">
import { ref, onMounted, h, nextTick, reactive } from "vue"
import team from '../apis/team/team'
import userCondition from "../stores/userStore";
import { ElNotification } from "element-plus";
import leader from '../apis/team/leader'
import teamCondition from '../stores/teamStore'


const newTeamStore = teamCondition();
const NewUserStore = userCondition();
const TableData = ref([]);

onMounted(async () => {
    const TeamID = NewUserStore.userSession.TeamID
    const res = await team.get(TeamID);
    if (res.data.code === 200 && res.data.msg === "Accepted") {
        TableData.value = res.data.data;
        const responseData = res.data.data;
        newTeamStore.setTeamInfo({
            TeamName: responseData.TeamName,
            State: responseData.State,
            TeamDescription: responseData.TeamDescription,
            TeamID: responseData.TeamID,
            Avatar: responseData.Avatar,
            TeamPassword: responseData.TeamPassword
        })
        console.log(NewUserStore.userSession)
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
    const res = await leader.quit({ Email: name, TeamID: id });
    if (res.data.code === 200) {
        TableData.value.splice(index, 1);
        ElNotification({
            title: "成功",
            message: h("i", { style: "color: teal" }, "踢出成功！"),
        });
        const TeamID = NewUserStore.userSession.TeamID
        const res = await team.get(TeamID);

        if (res.data.code === 200 && res.data.msg === "Accept") {
            TableData.value = res.data.data;
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
const form = reactive({
    TeamID: newTeamStore.teamSession.TeamID
})
const handleSignup = async () => {
    const res = await leader.submit(form);
    if (res.data.code === 200 && res.data.msg === "Accepted") {
        ElNotification({
            title: "成功",
            message: h("i", { style: "color:teal" }, "报名成功！"),
        });
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:teal" }, "报名失败，网络错误"),
        });
    }
};
/*
const DeleteTeam = async (index: number) => {
    const res = await leader.delete({
        TeamID: NewUserStore.userSession.TeamID,
        Email: NewUserStore.userSession.Email,
    });
    if (res.data.code === 200) {
        TableData.value.splice(index,);
        ElNotification({
            title: "成功",
            message: h("i", { style: "color: teal" }, "踢出成功！"),
        });
        const TeamID = NewUserStore.userSession.TeamID
        const res = await team.get(TeamID);

        if (res.data.code === 200 && res.data.msg === "Accept") {
            TableData.value = res.data.data;
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
*/
const isDisabled = ref(NewUserStore.userSession.Type !== 1);
const isSignup = (newTeamStore.teamSession.State === false)
/*
const AddMember = async () => {
    if (TeamCard.TeamID !== "") {
        const res = await leader.add(TeamCard);

        if (res.data.code === 200 && res.data.msg === "OK") {
            console.log(TeamCard);
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
            TeamCard.Username = value
            AddMember()
        })
}
*/
</script>

<template>
    <div class="div">
        <el-table :data="TableData" height="400" style="width:200%" border>
            <el-table-column prop="TeamID" label="团队号" width="200" align="center" show-overflow-tooltip />
            <el-table-column prop="Username" label="用户名" width="200" align="center" show-overflow-tooltip />
            <el-table-column prop="TelePhone" label="手机号" width="200" align="center" show-overflow-tooltip />
            <el-table-column prop="Email" label="邮箱" width="200" align="center" show-overflow-tooltip />
            <el-table-column label="操作" width="100" align="center">
                <template #default="scope">
                    <el-button type="danger" :disabled="isDisabled"
                        @click="handleDelete(scope.row.TeamID, scope.row.Email, scope.$index)">删除</el-button>
                    <br>
                    <el-button type="primary" :disabled="isDisabled" @click="">详情</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-button type="primary" :disabled="isDisabled" @click="handleSignup" v-if="isSignup">报名</el-button>
        <el-button type="primary" :disabled="isDisabled" @click="" v-if="!isSignup">取消报名</el-button>
        <el-button type="danger" :disabled="isDisabled" @click="DeleteTeam">解散团队</el-button>
    </div>
</template>

<style scoped>
.div {
    width: 100%;
}
</style>
