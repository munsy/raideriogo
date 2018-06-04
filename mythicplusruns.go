package raideriogo

// ViewMythicPlusRunsResponse defines the schema for mythic+ runs.
type ViewMythicPlusRunsResponse struct {
	Rankings       []Ranking `json:""`
	LeaderboardURL string    `json:""`
	Params         Params    `json:""`
}

// Dungeon defines the schema for a dungeon.
type Dungeon struct {
	ID              int    `json:""`
	Name            string `json:""`
	ShortName       string `json:""`
	Slug            string `json:""`
	KeystoneTimerMS int    `json:""`
}

// WeeklyModifier defines the schema for the weekly mythic+ modifiers.
type WeeklyModifier struct {
	ID          int    `json:""`
	Icon        string `json:""`
	Name        string `json:""`
	Description string `json:""`
}

// Class defines the schema for a class.
type Class struct {
	ID   int    `json:""`
	Name string `json:""`
	Slug string `json:""`
}

// Race defines the schema for a race.
type Race struct {
	ID   int    `json:""`
	Name string `json:""`
	Slug string `json:""`
}

// Spec defines the schema for a spec.
type Spec struct {
	Name string `json:""`
	Slug string `json:""`
}

// Region defines the schema for a region.
type Region struct {
	Name      string `json:""`
	Slug      string `json:""`
	ShortName string `json:""`
}

// Character defines the schema for a character.
type Character struct {
	ID        int    `json:""`
	PersonaID int    `json:""`
	Name      string `json:""`
	Class     Class  `json:""`
	Race      Race   `json:""`
	Faction   string `json:""`
	Spec      Spec   `json:""`
	Path      string `json:""`
	Realm     Realm  `json:""`
	Region    Region `json:""`
}

// Roster defines the schema for a roster.
type Roster struct {
	Character    Character `json:""`
	OldCharacter Character `json:""`
	IsTransfer   bool      `json:""`
	Role         string    `json:""`
}

// Run defines the schema for a run.
type Run struct {
	Season             string           `json:""`
	Dungeon            Dungeon          `json:""`
	KeystoneRunID      int              `json:""`
	KeystoneTeamID     int              `json:""`
	KeystonePlatoonID  int              `json:""`
	MythicLevel        int              `json:""`
	ClearTimeMS        int              `json:""`
	CompletedAt        string           `json:""`
	NumChests          int              `json:""`
	TimeRemainingMS    int              `json:""`
	Faction            string           `json:""`
	WeeklyModifier     []WeeklyModifier `json:""`
	NumModifiersActive int              `json:""`
	Roster             []Roster         `json:""`
	Platoon            string           `json:""`
}

// Ranking defines the schema for a  ranking.
type Ranking struct {
	Rank  int     `json:""`
	Score float64 `json:""`
	Run   Run     `json:""`
}

// Params defines the schema for this query.
type Params struct {
	Season  string `json:""`
	Region  string `json:""`
	Dungeon string `json:""`
}
