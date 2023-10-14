import { createRouter, createWebHistory } from "vue-router";
import { h } from "vue";
import { ElNotification } from "element-plus";
import Login from '../pages/Login.vue';
import Team from '../pages/Team.vue';
import Add from "../pages/Add.vue";
import Choice from "../pages/Choice.vue";
import Create from "../pages/Create.vue";
import Register from "../pages/Register.vue";
import pinia from "../stores";
import loginStore from "../stores/loginStore";
import Person from "../pages/Person.vue";
import Space from "../pages/Space.vue";
import Operator from "../pages/Operator.vue";


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
        component: Register
    },
    {
        path: "/Person",
        name: "Person",
        component: Person
    },
    {
        path: "/Space",
        name: "Space",
        component: Space
    },
    {
        path: "/Operator",
        name: "Operator",
        component: Operator
    }


]

const router = createRouter({
    history: createWebHistory(),
    routes: routes,
})

const login = loginStore(pinia)

router.beforeEach((to, _, next) => {
    if (!login.loginSession.valueOf()) {
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
    }
});

export default router;