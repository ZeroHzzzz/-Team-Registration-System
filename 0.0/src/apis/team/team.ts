import request from "../request";


export default class GetTeam {
    //获得团队成员信息
    static async get(TeamID: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/geteammember",
            method: "get",
            params: {
                TeamID: TeamID
            }
        });
    }
    //获得团队名片信息

    static async show(TeamID: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/get",
            method: "get",
            params: {
                TeamID: TeamID
            }
        });
    }

}