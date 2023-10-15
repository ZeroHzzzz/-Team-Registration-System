import request from "../request";
import { team } from "../..//types/personalInfo";


export default class getTeam {
    //创建团队
    static async create(data: { TeamName: string, Password: string, TeamDescription: string, LeaderID: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/create",
            method: "post",
            data: data
        });
    }
    //解散团队
    static async delete(data: { TeamID: string, Email: string, TeamPassword: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/del",
            method: "delete",
            data: data,
        });
    }
    //添加成员
    /*
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
    */
    //踢出成员
    static async quit(data: { Email: string, TeamID: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/quit",
            method: "delete",
            data: data,
        });
    }
    static async updateteam(data: team) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/updateprofile",
            method: "post",
            data: data,
        });
    }
    static async submit(data: { TeamID: string, }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/submit",
            method: "post",
            data: data,
        });
    }
    static async cancel(data: team) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/cancel",
            method: "post",
            data: data,
        });
    }
}