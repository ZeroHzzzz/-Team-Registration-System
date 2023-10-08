import request from "../request";

export default class userService {
    static async login(data: {
        account: string,
        password: string,
        key?: string,
    }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            method: "post",
            url: "/api/user/login",
            data: data,
        });
    }
    static async register(data: {
        account: string,
        password: string,
        key?: string
    }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            method: "post",
            url: "/api/user/register",
            data: data,
        });
    }
}