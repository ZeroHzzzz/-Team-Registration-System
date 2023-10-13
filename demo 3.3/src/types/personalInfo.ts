export interface team {

    TeamName: string,
    TeamID: string,
    Username: number,
    Password: string,
    TeamDescription: string
    State: boolean
    Avatar: string
}
export interface member {
    Username: number,
    Type: number,
    TeamID: number
}
export interface person {
    Email: string,
    Username: string,
    Type: number,
    Password: string,
    TeamID: string,
    Sign: string,
    Description: string,
    AvataraUrl: string,
    TelePhone: number
}

