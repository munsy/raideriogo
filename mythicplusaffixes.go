package raideriogo

// ViewMythicPlusAffixesResponse defines the schema for tracking mythic+ affixes.
type ViewMythicPlusAffixesResponse struct {
	Region         string                          `json:"region"`          // Region that affixes are for
	Title          string                          `json:"title"`           // Human readable title for the current affixes
	LeaderboardURL string                          `json:"leaderboard_url"` // URL for the leaderboard
	RecentRun      KeystoneRun                     `json:"recent_run"`      // KeystoneRun, optional
	AffixDetails   []IDNameDescriptionWowheadModel `json:"affix_details"`   // List of details for each active affix
}

// IdNameDescriptionWowheadModel defines the schema for mythic+ affixes via Wowhead data pulled into Raider.io.
type IDNameDescriptionWowheadModel struct {
	ID          int    `json:"id"`          // Blizzard's ID for this affix
	Name        string `json:"name"`        // Friendly name for this affix
	Description string `json:"description"` // Description of the effects of this affix
	WowheadURL  string `json:"wowhead_url"` // URL to open details on this affix on Wowhead
}
