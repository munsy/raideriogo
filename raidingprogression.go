package raideriogo

// ViewRaidingProgressionResponse defines the schema for tracking raiding progression.
type ViewRaidingProgressionResponse struct {
	Progression []ProgressTotalGuildsGuildsModel `json:"progression"` // (Array[ProgressTotalGuildsGuildsModel], optional)
}

// ProgressTotalGuildsGuildsModel defines the schema for tracking total progression across all guilds.
type ProgressTotalGuildsGuildsModel struct {
	Progress    int                    `json:"progress"`    // (integer): How many bosses out of all bosses in this instance ,
	TotalGuilds int                    `json:"totalGuilds"` // (integer): The number of guilds that are at this point of progression ,
	Guilds      []DefeatedAtGuildModel `json:"guilds"`      // (Array[DefeatedAtGuildModel], optional)
}

// DefeatedAtGuildModel defines the schema for tracking boss defeats by guild.
type DefeatedAtGuildModel struct {
	DefeatedAt string                               `json:"defeatedAt"` // (string),
	Guild      IDNameFactionLogoUrlRealmRegionModel `json:"guild"`      // (IdNameFactionLogoUrlRealmRegionModel, optional)
}

// IDNameFactionLogoUrlRealmRegionModel defines the schema for tracking the guild field in the DefeatedAtGuildModel struct.
type IDNameFactionLogoUrlRealmRegionModel struct {
	ID      int                      `json:"id"`      // (integer): Internal Raider.IO ID number for this guild ,
	Name    string                   `json:"name"`    // (string): Name of the guild ,
	Faction string                   `json:"faction"` // (string): Faction for this guild ,
	LogoURL string                   `json:"logoUrl"` // (string, optional): Custom logo URL for this guild ,
	Realm   NameSlugIsConnectedModel `json:"realm"`   // (NameSlugIsConnectedModel, optional),
	Region  NameShortNameSlugModel   `json:"region"`  // (NameShortNameSlugModel, optional)
}

// NameSlugIsConnectedModel defines the schema for tracking a name slug.
type NameSlugIsConnectedModel struct {
	Name        string `json:"name"`        // (string): Name of the realm ,
	Slug        string `json:"slug"`        // (string): Slug for the realm, suitable for putting in URLs ,
	IsConnected bool   `json:"isConnected"` // (boolean): Whether the realm is part of a connected realm
}

// NameShortNameSlugModel defines the schema for tracking a slug's short name.
type NameShortNameSlugModel struct {
	Name      string `json:"name"`       // (string): Name of the region ,
	ShortName string `json:"short_name"` // (string): Short name of the region ,
	Slug      string `json:"slug"`       // (string): Slug for the region, suitable for putting in URLs
}
