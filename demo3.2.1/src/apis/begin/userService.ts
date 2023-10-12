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
    }  //这个是不止传一个参数的写法
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