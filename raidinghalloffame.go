package raideriogo

// ViewRaidingHallOfFameResponse defines the schema for tracking the hall of fame.
type ViewRaidingHallOfFameResponse struct {
	HallOfFame HallOfFame `json:"hallOfFame"`
}

type HallOfFame struct {
	BossKills     BossKills     `json:"bossKills"`
	WinningGuilds WinningGuilds `json:"winningGuilds"`
}

type BossSummary struct {
	EncounterID int    `json:"encounterId"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Ordinal     int    `json:"ordinal"`
	WingID      int    `json:"wingId"`
	IconURL     string `json:"iconUrl"`
}

type DefeatedBy struct {
	TotalCount int    `json:"totalCount"`
	Guilds     Guilds `json:"guilds"`
}

// Realm defines the schema for a realm.
type Realm struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	AltSlug     string `json:"altSlug"`
	Locale      string `json:"locale"`
	IsConnected bool   `json:"isConnected"`
}

type GuildDefeat struct {
	Guild      Guild  `json:"guild"`
	DefeatedAt string `json:"defeatedAt"`
}

type WinningGuild struct {
	Rank               int           `json:"rank"`
	Guild              Guild         `json:"guild"`
	EncountersDefeated []interface{} `json:"encountersDefeated"`
}

type BossKill struct {
	Boss        string      `json:"boss"`
	BossSummary BossSummary `json:"bossSummary"`
	DefeatedBy  DefeatedBy  `json:"defeatedBy"`
}

type Guilds []GuildDefeat

type WinningGuilds []WinningGuild

type BossKills []BossKill
