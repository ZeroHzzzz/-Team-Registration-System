<script setup lang="ts">
import { Search, Edit, Delete } from '@element-plus/icons-vue'
import { h, ref, nextTick, reactive, onMounted } from 'vue'
import { ElNotification } from "element-plus";
import userCondition from "../stores/userStore";
import leader from '../apis/team/leader'
import team from '../apis/team/team'
import userService from '../apis/begin/userService'


const Searchinput = ref("");
const TableData2 = ref([]);
const pageSize = ref(7);
const currentPage = ref(1);
let total2 = 0;

const getTable = async() => {
    const res1 = await userService.get(Searchinput.value)
    
    if (res1.data.code === 200 && res1.data.msg === "ok") {
        TableData2.value = res1.data.data.userlist;
        total2 = TableData2.value.length;
    }
    else {
        ElNotification({
            title: "失败",
            message: h("i", { style: "color:firebrick" }, "搜索失败"),
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

//解散团队操作
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
            message: h("i", { style: "color: teal" }, "解散成功！"),
            type:"success"
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

</script>

<template>
    <div>
        <el-card class="multi-table">
            <el-row :gutter="12">
                <el-col :span="12">
                    <el-input  v-model="Searchinput" class="searchinput" placeholder="输入邮箱进行查询" :sufficon="Search">
                        <template #append>
                            <el-button :icon="Search" size="mini" @click="getTable" />
                        </template>
                    </el-input>
                </el-col>
            </el-row>
            <el-table class="team-table" :data="TableData2" height="500" style="width: 100%" border>
                <el-table-column prop="team_id" label="团队ID" width="200" align="center" show-overflow-tooltip />
                <el-table-column prop="team_condition" label="团队状态" width="100" align="center" show-overflow-tooltip />
                <el-table-column prop="email" label="邮箱" width="300" align="center" show-overflow-tooltip />
                <el-table-column prop="Username" label="用户名" width="200" align="center" show-overflow-tooltip />
                <el-table-column prop="type" label="用户权限" width="100" align="center" show-overflow-tooltip />
                    <el-table-column label="操作" width="300" align="center">
                        <template #default="scope">
                            <el-button type="primary" :icon="Edit" size="mini">编辑</el-button>
                            <el-button @click="handleDelete(scope.row.account, scope.row.TeamId, scope.$index)" type="danger" :icon="Delete" size="mini">注销</el-button>
                            <el-button type="info" :icon="Search" size="mini" @click="confirm">查询</el-button>
                        </template>
                    </el-table-column> 
            </el-table>
            <el-pagination
                class="pagination"
                v-model:current-page="currentPage"
                v-model:page-size="pageSize"
                layout="prev, pager, next, jumper"
                :total=total2
                @current-change="handleCurrentChange"
            />
        </el-card>
    </div>
</template>

<style scoped>
.multi-table{
    margin-top: 110px;
}
.team-table{
    margin-top: 15px;
}
.pagination{
    margin-top: 10px;
}
</style>