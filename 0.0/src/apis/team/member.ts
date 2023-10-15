import request from "../request";

export default class addTeam {
    //加入团队
    static async add(data: { Email: string, TeamID: string, TeamPassword: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/join",
            method: "post",
            data: data
        });
    }
    //退出团队
    static async out(data: { Email: string, TeamID: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/out",
            method: "delete",
            data: data,
        });
    }
}