import request from "../request";

export default class userService {
    static async login(data: {      //登录
        Email: string,
        Password: string,
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
    }  //注册
    static async register(data: {
        Username: string,
        Password: string,
        key?: string,
        Email: string
    }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            method: "post",
            url: "/api/user/register",
            data: data,
        });
    }  //信息
    static async get(email: string) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            method: "get",
            url: "/api/user/get",
            params: {
                Email: email,
            }
        });
    }  //前端用这个把数据传递给后端
}