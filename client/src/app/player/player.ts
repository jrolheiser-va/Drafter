export interface Player {
    ID: string;
    LastName: string;
    FirstName: string;
    Position: string;

    TeamID: string;
    TeamAbbreviation: string;

    GamesPlayed: number;
    Goals: number;
    Assists: number;
    Points: number;
    PlusMinus: number;
}
