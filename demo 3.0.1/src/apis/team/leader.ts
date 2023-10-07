import request from "../request";
import { team } from "../..//types/personalInfo";
import { } from "../../types/personalInfo";

export default class getTeam {
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
    static async delete(data: { team_id: number, account: string }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/delete",
            method: "delete",
            data: data,
        });
    }
    static async cick(data: { user_id: number, team_id: number }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/teamCick",
            method: "delete",
            data: data,
        });
    }
}