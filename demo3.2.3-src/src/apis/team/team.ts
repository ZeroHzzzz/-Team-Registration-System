import request from "../request";


export default class GetTeam {
    //获得团队成员信息
    static async get(Email: string, TeamID: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/get",
            methods: "get",
            params: {
                Email: Email,
                TeamID: TeamID
            }
        });
    }
    //获得团队名片信息
    static async show(Username: string, TeamID: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/show",
            method: "get",
            params: {
                Username: Username,
                TeamID: TeamID
            }
        });
    }
}