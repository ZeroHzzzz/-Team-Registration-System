<script setup lang="ts">
import { Search, Delete } from '@element-plus/icons-vue'
import { h, ref, nextTick, onMounted } from 'vue'
import { ElNotification } from "element-plus";
import userCondition from "../stores/userStore";
import operator from '../apis/operator/operator'
import team from '../apis/team/team'
import userService from '../apis/begin/userService'
import openPerson from "../components/personcard.vue";

const Searchinput = ref("");
let TableData2 = ref([]);
const pageSize = ref(9);
const currentPage = ref(1);
let total2 = 0;

const getTable = async () => {
    const res1 = await userService.get(Searchinput.value)

    if (res1.data.code === 200 && res1.data.msg === "ok") {
        TableData2.value = res1.data.data.userlist;
        total2 = TableData2.value.length;
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:firebrick" }, "搜索失败,未输入内容或无搜索结果"),
            type: "error",
        });
    }
}

/* let tabletype = "team"
const change = () =>{
    if (tabletype === "team") {                 
        tabletype = "member"
    }
    else {
        tabletype = "team"
    }
} */   //备用,可能会用按钮切换两张表格，待定

/* const confirm = () => {
    router.push("/Team")
} */   //待定内容,点击查询前往团队界面

//分页相关配件
const handleCurrentChange = (val: number) => {
    console.log(`current page: ${val}`)
}

//获取用户所处的团队的团队信息
const NewUserStore = userCondition();
const TableData = ref([])

onMounted(async () => {
    const TeamID = NewUserStore.userSession.TeamID

    const res = await team.get(TeamID);

    if (res.data.code === 200 && res.data.msg === "OK") {
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

//解散团队操作
const renderComponent = ref(true);
const forceRender = async () => {
    renderComponent.value = false
    await nextTick();
    renderComponent.value = true;
}

const handleDelete = async (Email: string, index: number) => {
    const res = await operator.deluser({ Email: Email, });
    if (res.data.code === 200) {
        TableData.value.splice(index, 1);
        ElNotification({
            title: "成功",
            message: h("i", { style: "color: teal" }, "解散成功！"),
            type: "success"
        });

        const TeamID = NewUserStore.userSession.TeamID
        console.log(TeamID);
        const res = await team.get(TeamID);

        if (res.data.code === 200 && res.data.msg === "OK") {
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

</script>

<template>
    <div>
        <el-card class="multi-table">
            <el-row :gutter="12">
                <el-col :span="12">
                    <el-input v-model="Searchinput" class="searchinput" placeholder="输入邮箱进行查询" :sufficon="Search">
                        <template #append>
                            <el-button :icon="Search" size="mini" @click="getTable" />
                        </template>
                    </el-input>
                </el-col>
            </el-row>
            <el-table class="team-table" :data="TableData2.slice((currentPage - 1) * pageSize, currentPage * pageSize)"
                height="480" style="width: 100%" border>
                <el-table-column prop="team_id" label="团队ID" width="200" align="center" show-overflow-tooltip />
                <el-table-column prop="team_condition" label="团队状态" width="100" align="center" show-overflow-tooltip />
                <el-table-column prop="email" label="邮箱" width="300" align="center" show-overflow-tooltip />
                <el-table-column prop="Username" label="用户名" width="200" align="center" show-overflow-tooltip />
                <el-table-column prop="type" label="用户权限" width="100" align="center" show-overflow-tooltip />
                <el-table-column label="操作" width="300" align="center">
                    <template #default="scope">
                        <el-button @click="handleDelete(scope.row.Email, scope.$index)" type="danger" :icon="Delete"
                            size="mini">注销</el-button>
                        <el-button type="info" :icon="Search" size="mini" @click="openPerson">详情</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination class="pagination" v-model:current-page="currentPage" v-model:page-size="pageSize"
                layout="prev, pager, next, jumper" :total=total2 @current-change="handleCurrentChange" />
        </el-card>
    </div>
</template>

<style scoped>
.multi-table {
    margin-top: 115px;
}

.team-table {
    margin-top: 15px;
}

.pagination {
    margin-top: 10px;
}
</style>