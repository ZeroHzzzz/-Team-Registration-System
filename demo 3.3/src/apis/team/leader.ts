import request from "../request";
import { team } from "../..//types/personalInfo";
import { } from "../../types/personalInfo";

export default class getTeam {
    //创建团队
    static async create(data: team) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/create",
            methond: "post",
            data: data
        });
    }
    //解散团队
    static async delete(data: { TeamID: string, Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/delete",
            method: "delete",
            data: data,
        });
    }
    //添加成员
    static async add(data: { TeamID: string, Email: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/add",
            method: "post",
            data: data,
        });
    }
    //踢出成员
    static async cick(data: { Email: string, TeamID: number }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/cick",
            method: "delete",
            data: data,
        });
    }
}