package raideriogo

// ViewRaidingBossRankingsResponse defines the schema for tracking raid boss rankings.
type ViewRaidingBossRankingsResponse struct {
	BossRankings []BossRanking `json:"bossRankings"`
}

// Guild defines the schema for a guild.
type Guild struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Faction string `json:"faction"`
	LogoURL string `json:"logoUrl"`
	Realm   Realm  `json:"realm"`
	Region  Region `json:"region"`
}

// BossRanking defines the schema for a boss ranking.
type BossRanking struct {
	Rank               int                `json:"rank"`
	RegionRank         int                `json:"regionRank"`
	Guild              Guild              `json:"guild"`
	EncountersDefeated EncountersDefeated `json:"encountersDefeated"`
}
