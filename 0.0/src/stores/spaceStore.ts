import { defineStore } from "pinia";
import { reactive } from "vue";
import { person } from "../types/personalInfo";

const visitCondition = defineStore(
    "user",
    () => {
        const visitSession = reactive({
            Email: "",
            Username: "未登录",
            Type: -1,
            TeamID: "",
            Sign: "哼啊啊啊啊啊啊啊啊",
            Description: "这个用户很懒，没有写简介",
            AvataraUrl: ""
        });

        const setUserInfo = (info: person): void => {
            visitSession.Email = info.Email;
            visitSession.Username = info.Username
            visitSession.Type = info.Type
            visitSession.TeamID = info.TeamID
            visitSession.Sign = info.Sign
            visitSession.Description = info.Description
            visitSession.AvataraUrl = info.AvataraUrl
        };

        const setUserAva = (): void => {

        }
        return {
            visitSession,
            setUserInfo,
            setUserAva
        };
    },
    {
        persist: true
    }
);

export default visitCondition;