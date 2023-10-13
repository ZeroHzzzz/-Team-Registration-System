import request from '../request'


export default class ava {
    static async get(Email: string,) {
        return request({
            "headers": {
                "Content-Type": "multipart/form-data",
            },
            url: "/api/user/avaget",
            method: "get",
            params: {
                Email: Email
            }
        });
    }
    static async add(data: FormData) {
        return request({
            "headers": {
                "Content-Type": "multipart/form-data",
            },
            url: "/api/user/avaadd",
            method: "post",
            data: data
        });
    }
}