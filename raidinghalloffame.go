package raideriogo

// ViewRaidingHallOfFameResponse defines the schema for a complete hall of fame response.
type ViewRaidingHallOfFameResponse struct {
	HallOfFame HallOfFame `json:"hallOfFame"`
}

// HallOfFame defines the schema for the hall of fame.
type HallOfFame struct {
	BossKills     BossKills     `json:"bossKills"`
	WinningGuilds WinningGuilds `json:"winningGuilds"`
}

// BossSummary defines the schema for a boss summary.
type BossSummary struct {
	EncounterID int    `json:"encounterId"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Ordinal     int    `json:"ordinal"`
	WingID      int    `json:"wingId"`
	IconURL     string `json:"iconUrl"`
}

// DefeatedBy defines the schema for the guilds that have defeated a given boss.
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

// GuildDefeat defines the schema for a boss defeat by a guild.
type GuildDefeat struct {
	Guild      Guild  `json:"guild"`
	DefeatedAt string `json:"defeatedAt"`
}

// WinningGuild defines the schema for a winning guild.
type WinningGuild struct {
	Rank               int           `json:"rank"`
	Guild              Guild         `json:"guild"`
	EncountersDefeated []interface{} `json:"encountersDefeated"`
}

// BossKill defines the schema for a boss kill.
type BossKill struct {
	Boss        string      `json:"boss"`
	BossSummary BossSummary `json:"bossSummary"`
	DefeatedBy  DefeatedBy  `json:"defeatedBy"`
}

// Guilds is a type definition for an array of GuildDefeat structs.
type Guilds []GuildDefeat

// WinningGuilds is a type definition for an array of WinningGuild structs.
type WinningGuilds []WinningGuild

// BossKills is a type definition for an array of BossKill structs.
type BossKills []BossKill
