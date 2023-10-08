import request from "../request";
import { team } from "../../types/personalInfo";

export default class addTeam {
    static async add(data: team) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/add",
            method: "post",
            data: data
        });
    }
    static async out(data: { user_id: number }) {
        return request({
            "headers": {
                "Content-Type": "application/json",
            },
            url: "/api/team/out",
            method: "delete",
            data: data,
        });
    }
}