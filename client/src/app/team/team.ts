export interface Team {
    ID: string;
    Name: string;
}

export const PLAYOFF_TEAMS: Team[] = [
    {ID: 'NSH', Name: 'Nashville Predators'},
    {ID: 'COL', Name: 'Colorado Avalanche'},

    {ID: 'WPJ', Name: 'Winnipeg Jets'},
    {ID: 'MIN', Name: 'Minnesota Wild'},

    {ID: 'VGK', Name: 'Vegas Golden Knights'},
    {ID: 'LAK', Name: 'Los Angeles Kings'},

    {ID: 'ANA', Name: 'Anaheim Ducks'},
    {ID: 'SJS', Name: 'San Jose Sharks'},


    {ID: 'TBL', Name: 'Tampa Bay Lightning'},
    {ID: 'NJD', Name: 'New Jersey Devils'},

    {ID: 'BOS', Name: 'Boston Bruins'},
    {ID: 'TOR', Name: 'Toronto Maple Leafs'},

    {ID: 'WSH', Name: 'Washington Capitals'},
    {ID: 'CBJ', Name: 'Columbus Blue Jackets'},

    {ID: 'PIT', Name: 'Pittsburgh Penguins'},
    {ID: 'PHI', Name: 'Philadelphia Flyers'},
];

export const PLAYOFF_TEAM_IDS: string[] = PLAYOFF_TEAMS.map(
    team => team.ID
);
















