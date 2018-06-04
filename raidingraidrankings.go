package raideriogo

import "time"

// ViewRaidingRaidRankingsResponse defines the schema for tracking raid rankings.
type ViewRaidingRaidRankingsResponse struct {
	RaidRankings RaidRankings `json:"raidRankings"`
}

// EncounterDefeated defines the schema for encounters defeated.
type EncounterDefeated struct {
	Slug          string    `json:"slug"`
	LastDefeated  time.Time `json:"lastDefeated"`
	FirstDefeated time.Time `json:"firstDefeated"`
}

// EncountersDefeated defines the schema for encounters defeated.
type EncountersDefeated []EncounterDefeated

// RaidRanking defines the schema for a raid ranking.
type RaidRanking struct {
	Rank               int                `json:"rank"`
	RegionRank         int                `json:"regionRank"`
	Guild              Guild              `json:"guild"`
	EncountersDefeated EncountersDefeated `json:"encountersDefeated"`
}

// RaidRankings defines the schema for raid rankings.
type RaidRankings []RaidRanking
