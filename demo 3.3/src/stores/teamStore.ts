import { defineStore } from "pinia";
import { reactive } from "vue";
import { team } from "../types/personalInfo";

const teamCondition = defineStore(
    "team",
    () => {
        const teamSession = reactive({
            TeamName: "",
            State: false,
            TeamDescription: "",
            TeamID: "",
            Teamava: ""
        })



        const setTeamInfo = (info: team): void => {
            teamSession.TeamName = info.TeamName;
            teamSession.State = info.State;
            teamSession.TeamDescription = info.TeamDescription;
            teamSession.TeamID = info.TeamID;
            teamSession.Teamava = info.Avatar;
            ;
        }

        return {
            teamSession,
            setTeamInfo

        }
    }
)

export default teamCondition