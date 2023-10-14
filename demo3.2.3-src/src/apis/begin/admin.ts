import request from "../request";

export default class admin {
    static async delteam(data: {
        team_id:number,
    }) {
        return request({
            "headers": {
                "Content-Type": "application/json"
            },
            method:"delete",
            url:"/api/admin/delteam",
            data:data,
        });
    }
}