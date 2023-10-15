<script setup lang="ts">
import GetTeam from '../apis/team/team'
import { computed, onMounted, reactive } from "vue";
import { useRoute } from "vue-router";
import userCondition from "../stores/userStore";
import teamCondition from '../stores/teamStore';

const newTeamStore = teamCondition()
const NewUserStore = userCondition()
const route = useRoute()
const IsUserPage = computed(() => {
    return !(route.name === "Login" || route.name === "Register" || route.name === "Team")
})

const IsTeamPage = computed(() => {
    return (route.name === "Team")
})

const PersonCard = reactive({
    UserName: NewUserStore.userSession.Username,
    Sign: NewUserStore.userSession.Sign,
    Description: NewUserStore.userSession.Sign
})
const TeamCard = reactive({
    TeamName: newTeamStore.teamSession.TeamName,
    TeamDescription: newTeamStore.teamSession.TeamDescription
})
onMounted(async () => {
    const TeamID = NewUserStore.userSession.TeamID
    const res = await GetTeam.show(TeamID);

    if (res.data.code === 200 && res.data.msg === "OK") {
        TeamCard.TeamName = res.data.data
    }
})

</script>

<template>
    <div class="header">
        <div v-if="IsUserPage" class="team">
            <header>
                <div class="ava">

                </div>
                <div class="person_text">
                    <h1>{{ PersonCard.UserName }}</h1>
                    <h2>{{ PersonCard.Sign }}</h2>
                </div>
            </header>
        </div>
        <div v-if="IsTeamPage" class="person">
            <header>
                <div class="ava">

                </div>
                <div class="team_text">
                    <h1>{{ TeamCard.TeamName }}</h1>
                    <h2>{{ TeamCard.TeamDescription }}</h2>
                </div>
            </header>
        </div>
    </div>
</template>

<style scoped>
.header {
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    background-color: #394467;
    z-index: 50;
    height: 197px;
}


.title {
    color: #ededed;
}
</style>