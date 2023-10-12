import { createRouter, createWebHistory } from "vue-router";
import { h } from "vue";
import { ElNotification } from "element-plus";
import Login from '../pages/Login.vue';
import Team from '../pages/Team.vue';
import Add from "../pages/Add.vue";
import Choice from "../pages/Choice.vue";
import Create from "../pages/Create.vue";
import register from "../pages/Register.vue";
import operator from '../pages/Operator.vue'

const routes = [
    {
        path: "/",
        redirect: "/Login",
    },
    {
        path: "/Login",
        name: "Login",
        component: Login
    },
    {
        path: "/Team",
        name: "Team",
        component: Team
    },
    {
        path: "/Add",
        name: "Add",
        component: Add
    },
    {
        path: "/Choice",
        name: "Choice",
        component: Choice
    },
    {
        path: "/Create",
        name: "Create",
        component: Create
    },
    {
        path: "/Register",
        name: "Register",
        component: register
    },
    {
        path: "/Operator",
        name: "Operator",
        component: operator
    }

]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

router.beforeEach((to, _, next) => {
    if (localStorage.getItem("login") === "false") {
        if (to.path === "/Login") {
            next();
        }
        else {
            ElNotification({
                title: "失败",
                message: h("i", { style: "color: teal" }, "请您先登录！"),
            });
            next("/Login");
        }
    }
    else {
        next();
    }
});

export default router;