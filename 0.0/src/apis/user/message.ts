import request from "../request";



export default class PopMessage {
    static async get(Email: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/msg/getmsg",
            methond: "get",
            params: {
                Email: Email
            }
        })

    }
    static async update(data: { Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/msg/updateunread",
            methond: "post",
            data: data,
        })

    }
}