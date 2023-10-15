import { defineStore } from "pinia";
import { reactive } from "vue";
import { person } from "../types/personalInfo";

const userCondition = defineStore(
    "user",
    () => {
        const userSession = reactive({
            Email: "",
            Username: "未登录",
            Type: -1,
            TeamID: "",
            Sign: "哼啊啊啊啊啊啊啊啊",
            Description: "这个用户很懒，没有写简介",
            AvataraUrl: "",
            TelePhone: 10000000000
        });

        const setUserInfo = (info: person): void => {
            userSession.Email = info.Email;
            userSession.Username = info.Username
            userSession.Type = info.Type
            userSession.TeamID = info.TeamID
            userSession.Sign = info.Sign
            userSession.Description = info.Description
            userSession.AvataraUrl = info.AvataraUrl
            userSession.TelePhone = info.TelePhone
        };

        const setUserAva = (): void => {

        }
        return {
            userSession,
            setUserInfo,
            setUserAva
        };
    },
    {
        persist: true
    }
);

export default userCondition;