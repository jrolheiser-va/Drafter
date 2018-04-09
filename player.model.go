package main

type Player struct {
	ID        string
	LastName  string
	FirstName string
	Position  string
	// Age          int64
	// JerseyNumber string
	// Height       string
	// Weight       string
	// BirthDate    string
	// BirthCity    string
	// BirthCountry string
	// IsRookie     bool

	TeamID           string
	TeamAbbreviation string

	GamesPlayed   int64
	Goals         int64
	Assists       int64
	Points        int64
	PointsPerGame float64
	PlusMinus     int64
	// HatTricks          int64
	// Shots              int64
	// BlockedShots       int64
	// ShotPercentage     int64
	// Penalties          int64
	// PenaltyMinutes     int64
	// PowerplayGoals     int64
	// PowerplayAssists   int64
	// PowerplayPoints    int64
	// ShorthandedGoals   int64
	// ShorthandedAssists int64
	// ShorthandedPoints  int64
	// GameWinningGoals   int64
	// GameTyingGoals     int64
	// Hits               int64
	// Faceoffs           int64
	// FaceoffWins        int64
	// FaceoffLosses      int64
	// FaceoffPercent     int64
}

func toPlayer(msfPlayer MSFPlayer) Player {
	return Player{
		ID:        msfPlayer.Description.ID,
		LastName:  msfPlayer.Description.LastName,
		FirstName: msfPlayer.Description.FirstName,
		Position:  msfPlayer.Description.Position,
		// JerseyNumber:     msfPlayer.Description.JerseyNumber,
		// Height:           msfPlayer.Description.Height,
		// Weight:           msfPlayer.Description.Weight,
		// BirthDate:        msfPlayer.Description.BirthDate,
		// Age:              msfPlayer.Description.Age,
		// BirthCity:        msfPlayer.Description.BirthCity,
		// BirthCountry:     msfPlayer.Description.BirthCountry,
		// IsRookie:         msfPlayer.Description.IsRookie,
		TeamID:           msfPlayer.Team.ID,
		TeamAbbreviation: msfPlayer.Team.Abbreviation,
		GamesPlayed:      msfPlayer.PlayerStats.GamesPlayed.Value,
		Goals:            msfPlayer.PlayerStats.Stats.Goals.Value,
		Assists:          msfPlayer.PlayerStats.Stats.Assists.Value,
		Points:           msfPlayer.PlayerStats.Stats.Points.Value,
		PlusMinus:        msfPlayer.PlayerStats.Stats.PlusMinus.Value,
		PointsPerGame:    float64(msfPlayer.PlayerStats.Stats.Points.Value) / float64(msfPlayer.PlayerStats.GamesPlayed.Value),
		// HatTricks:          msfPlayer.PlayerStats.Stats.HatTricks.Value,
		// Shots:              msfPlayer.PlayerStats.Stats.Shots.Value,
		// BlockedShots:       msfPlayer.PlayerStats.Stats.BlockedShots.Value,
		// ShotPercentage:     msfPlayer.PlayerStats.Stats.ShotPercentage.Value,
		// Penalties:          msfPlayer.PlayerStats.Stats.Penalties.Value,
		// PenaltyMinutes:     msfPlayer.PlayerStats.Stats.PenaltyMinutes.Value,
		// PowerplayGoals:     msfPlayer.PlayerStats.Stats.PowerplayGoals.Value,
		// PowerplayAssists:   msfPlayer.PlayerStats.Stats.PowerplayAssists.Value,
		// PowerplayPoints:    msfPlayer.PlayerStats.Stats.PowerplayPoints.Value,
		// ShorthandedGoals:   msfPlayer.PlayerStats.Stats.ShorthandedGoals.Value,
		// ShorthandedAssists: msfPlayer.PlayerStats.Stats.ShorthandedAssists.Value,
		// ShorthandedPoints:  msfPlayer.PlayerStats.Stats.ShorthandedPoints.Value,
		// GameWinningGoals:   msfPlayer.PlayerStats.Stats.GameWinningGoals.Value,
		// GameTyingGoals:     msfPlayer.PlayerStats.Stats.GameTyingGoals.Value,
		// Hits:               msfPlayer.PlayerStats.Stats.Hits.Value,
		// Faceoffs:           msfPlayer.PlayerStats.Stats.Faceoffs.Value,
		// FaceoffWins:        msfPlayer.PlayerStats.Stats.FaceoffWins.Value,
		// FaceoffLosses:      msfPlayer.PlayerStats.Stats.FaceoffLosses.Value,
		// FaceoffPercent:     msfPlayer.PlayerStats.Stats.FaceoffPercent.Value,
	}
}
