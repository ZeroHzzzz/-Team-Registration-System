import { defineStore } from "pinia";
import { ref, reactive } from "vue";
import { session } from "../types/user";

const loginStore = defineStore(
    "login",
    () => {
        const loginSession = ref(false);
        const setlog = (logNewSession: boolean) => {
            loginSession.value = logNewSession;
        };
        const sessionjwt = reactive({
            sessionkey: "未获取"
        });
        const SetSessionInfo = (info: session): void => {
            sessionjwt.sessionkey = info.sessionkey
        }
        return {
            sessionjwt,
            SetSessionInfo,
            loginSession,
            setlog,
        };
    },
    {
        persist: true,
    }

);
export default loginStore;
