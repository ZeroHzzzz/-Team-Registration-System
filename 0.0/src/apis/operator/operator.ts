import request from "../request";

export default class admin {

    //解散团队
    static async delteam(data: { TeamID: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/admin/delteam",
            method: "delete",
            data: data,
        });
    }
    static async deluser(data: { Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/admin/deluser",
            method: "delete",
            data: data,
        });
    }
    static async getuser(data: { Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/admin/getalluser",
            method: "get",
            data: data,
        });
    }
    static async getteam(data: { Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/admin/getallteam",
            method: "get",
            data: data,
        });
    }
}