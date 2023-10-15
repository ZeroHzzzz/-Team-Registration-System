import { defineStore } from "pinia";
import { ref } from "vue";

const loginStore = defineStore(
    "login",
    () => {
        const loginSession = ref(false);
        const setlog = (logNewSession: boolean) => {
            loginSession.value = logNewSession;
        };
        return {
            loginSession,
            setlog,
        };
    },
    {
        persist: true,
    }

);
export default loginStore;
