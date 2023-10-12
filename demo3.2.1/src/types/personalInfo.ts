export interface admin {
    Email: string,
    Username: string,
    Password: string,
    type: number,
    team_id: number,
    team_condition: number,
}
export interface team {
    teamName: string,
    team_id: number,
    user_id: number,
    teamPassword: string,
    teamCondition: boolean
}
export interface member {
    Account: string,
    user_id: number,
    type: number,
    team_id: number
}
export interface person {
    user_id: number,
    account: string,
    type: number,
    password: string,
    team_id: number
}
export interface space {
    user_id: number,
    account: string,

}
