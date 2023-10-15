import request from '../request'


export default class ava {
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