import request from "../request";



export default class PopMessage {
    static async get(account: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/user/getmessage",
            methond: "get",
            params: {
                account: account
            }
        })

    }
}