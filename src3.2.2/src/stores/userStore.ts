import { defineStore } from "pinia";
import { reactive } from "vue";
import { person } from "../types/personalInfo";

const userCondition = defineStore(
    "user",
    () => {
        const userSession = reactive({
            id: -1,
            account: "未登录",
            type: 0,
            team_id: 0,
        });

        const setUserInfo = (info: person): void => {
            userSession.id = info.user_id;
            userSession.account = info.account
            userSession.type = info.type
            userSession.team_id = info.team_id
        };
        return {
            userSession,
            setUserInfo,
        };
    },
    {
        persist: true
    }
);

export default userCondition;