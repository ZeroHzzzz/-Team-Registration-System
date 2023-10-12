import request from "../request";


export default class GetTeam {
    static async get(Account: string, TeamId: number) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team",
            methods: "get",
            params: {
                user_id: Account,
                TeamId: TeamId
            }
        });
    }
    static async show(user_id: number) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/show",
            method: "get",
            params: {
                user_id: user_id
            }
        });
    }
}