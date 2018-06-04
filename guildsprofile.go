package raideriogo

// ViewGuildProfileResponse defines the schema for tracking a guild's generic profile info.
type ViewGuildProfileResponse struct {
	Name            string                       `json:"name"`             // Name of guild
	Faction         string                       `json:"faction"`          // Guild's faction
	Region          string                       `json:"region"`           // Region guild belongs to
	Realm           string                       `json:"realm"`            // Realm guild resides on
	ProfileURL      string                       `json:"profile_url"`      // URL for Guild's profile on Raider.IO
	RaidRankings    RaidRankingsSchema           `json:"raid_rankings"`    // Optional
	RaidProgression OverallRaidProgressionSchema `json:"raid_progression"` // Optional
}

// RaidRankingsSchema defines the schema for tracking raid rankings.
type RaidRankingsSchema struct {
	AntorusTheBurningThrone RaidDifficultyRankingsSchema `json:"antorus-the-burning-throne"` // optional
	TombOfSargeras          RaidDifficultyRankingsSchema `json:"tomb-of-sargeras"`           // optional
	TheNighthold            RaidDifficultyRankingsSchema `json:"the-nighthold"`              // optional
	TheEmeraldNightmare     RaidDifficultyRankingsSchema `json:"the-emerald-nightmare"`      // optional
	TrialOfValor            RaidDifficultyRankingsSchema `json:"trial-of-valor"`             // optional
}

// RaidDifficultyRankingsSchema defines the schema for tracking rankings by difficulty.
type RaidDifficultyRankingsSchema struct {
	Normal RanksSchema `json:"normal"` // optional
	Heroic RanksSchema `json:"heroic"` // optional
	Mythic RanksSchema `json:"mythic"` // optional
}
