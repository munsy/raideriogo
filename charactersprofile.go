package raideriogo

// ViewCharacterProfileResponse defines the schema for tracking a character's generic profile info.
type ViewCharacterProfileResponse struct {
	Name                     string                       `json:"name"`
	Race                     string                       `json:"race"`
	Class                    string                       `json:"class"`
	ActiveSpecName           string                       `json:"active_spec_name"`
	ActiveSpecRole           string                       `json:"active_spec_role"`
	Gender                   string                       `json:"gender"`
	Faction                  string                       `json:"faction"`
	AchievementPoints        int                          `json:"achievement_points"`
	HonorableKills           int                          `json:"honorable_kills"`
	ThumbnailURL             string                       `json:"thumbnail_url"`
	Region                   string                       `json:"region"`
	Realm                    string                       `json:"realm"`
	ProfileURL               string                       `json:"profile_url"`
	Gear                     CharacterGearSchema          `json:"gear"`                        // optional
	RaidProgression          OverallRaidProgressionSchema `json:"raid_progression"`            // optional
	MythicPlusRanks          MythicPlusRoleRanksSchema    `json:"mythic_plus_ranks"`           // optional
	MythicPlusScore          MythicPlusScoresSchema       `json:"mythic_plus_scores"`          // optional
	PreviousMythicPlusRanks  MythicPlusRoleRanksSchema    `json:"previous_mythic_plus_ranks"`  // optional
	PreviousMythicPlusScores MythicPlusScoresSchema       `json:"previous_mythic_plus_scores"` // optional
	MythicPlusRecentRuns     []KeystoneRun                `json:"mythic_plus_recent_runs"`     // optional Character's three most recent Mythic+ runs from current season
	MythicPlusBestRuns       []KeystoneRun                `json:"mythic_plus_best_runs"`       // optional Character's three highest scoring Mythic+ runs from current season
}

// CharacterGearSchema defines the schema for tracking a character's item level and artifact info.
type CharacterGearSchema struct {
	ItemLevelEquipped int `json:"item_level_equipped"` // Character's equipped item level
	ItemLevelTotal    int `json:"item_level_total"`    // Character's total item level in bags
	ArtifactTraits    int `json:"artifact_traits"`     // Character's artifact traits on equipped artifact weapon
}

// OverallRaidProgressionSchema defines the schema for tracking a character's overall raid progression.
type OverallRaidProgressionSchema struct {
	AntorusTheBurningThrone RaidProgressionSchema `json:"antorus-the-burning-throne"` // optional
	TombOfSargeras          RaidProgressionSchema `json:"tomb-of-sargeras"`           // optional
	TheNighthold            RaidProgressionSchema `json:"the-nighthold"`              // optional
	TheEmeraldNightmare     RaidProgressionSchema `json:"the-emerald-nightmare"`      // optional
	TrialOfValor            RaidProgressionSchema `json:"trial-of-valor"`             // optional
}

// MythicPlusRoleRanksSchema defines the schema for tracking a character's mythic+ rankings by role.
type MythicPlusRoleRanksSchema struct {
	Overall     RanksSchema `json:"overall"`
	Tank        RanksSchema `json:"tank"`
	Healer      RanksSchema `json:"healer"`
	DPS         RanksSchema `json:"dps"`
	Class       RanksSchema `json:"class"`
	ClassTank   RanksSchema `json:"class_tank"`
	ClassHealer RanksSchema `json:"class_healer"`
	ClassDPS    RanksSchema `json:"class_dps"`
}

// MythicPlusScoresSchema defines the schema for tracking a character's mythic+ scores by role.
type MythicPlusScoresSchema struct {
	All    int `json:"all"`    // Player's best score across all possible roles
	DPS    int `json:"dps"`    // Player's score from DPS roles
	Healer int `json:"healer"` // Player's score from healer roles
	Tank   int `json:"tank"`   // Player's score from tank roles
}

// KeystoneRun defines the schema for tracking a character's keystone run history.
type KeystoneRun struct {
	Dungeon             string `json:"dungeon"`               // Friendly name of Dungeon
	ShortName           string `json:"short_name"`            // Shortened or abbreviated name of Dungeon
	MythicLevel         int    `json:"mythic_level"`          // Mythic+ Level of the run
	CompletedAt         string `json:"completed_at"`          // When this run was completed
	ClearTimeMs         int    `json:"clear_time_ms"`         // How long it took to complete the dungeon, in milliseconds
	NumKeystoneUpgrades int    `json:"num_keystone_upgrades"` // How many times the keystone used for this dungeon would have been upgraded after completion
	Score               int    `json:"score"`                 // How many points this run was worth
	URL                 string `json:"url"`                   // URL to this specific run to view roster details
}

// RaidProgressionSchema defines the schema for tracking a character's raid progression.
type RaidProgressionSchema struct {
	Summary            string `json:"summary"`              // Human readable summary of progression
	TotalBosses        int    `json:"total_bosses"`         // Number of bosses in instance
	NormalBossesKilled int    `json:"normal_bosses_killed"` // Number of normal mode bosses killed
	HeroicBossesKilled int    `json:"heroic_bosses_killed"` // Number of heroic mode bosses killed
	MythicBossesKilled int    `json:"mythic_bosses_killed"` // Number of mythic mode bosses killed
}

// RanksSchema defines the schema for tracking a character's world, region, and realm rankings.
type RanksSchema struct {
	World  int `json:"world"`  // World rank
	Region int `json:"region"` // Region rank
	Realm  int `json:"realm"`  // Realm rank
}
