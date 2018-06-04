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

// Realm defines the schema for a realm.
type Realm struct {
	ID          int    `json:""`
	Name        string `json:""`
	Slug        string `json:""`
	AltSlug     string `json:""`
	Locale      string `json:""`
	IsConnected bool   `json:""`
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

/*
public class Dungeon
{
    public int id { get; set; }
    public string name { get; set; }
    public string short_name { get; set; }
    public string slug { get; set; }
    public int keystone_timer_ms { get; set; }
}

public class WeeklyModifier
{
    public int id { get; set; }
    public string icon { get; set; }
    public string name { get; set; }
    public string description { get; set; }
}

public class Class
{
    public int id { get; set; }
    public string name { get; set; }
    public string slug { get; set; }
}

public class Race
{
    public int id { get; set; }
    public string name { get; set; }
    public string slug { get; set; }
}

public class Spec
{
    public string name { get; set; }
    public string slug { get; set; }
}

public class Realm
{
    public int id { get; set; }
    public string name { get; set; }
    public string slug { get; set; }
    public string altSlug { get; set; }
    public string locale { get; set; }
    public bool isConnected { get; set; }
}

public class Region
{
    public string name { get; set; }
    public string slug { get; set; }
    public string short_name { get; set; }
}

public class Character
{
    public int id { get; set; }
    public int persona_id { get; set; }
    public string name { get; set; }
    public Class @class { get; set; }
    public Race race { get; set; }
    public string faction { get; set; }
    public Spec spec { get; set; }
    public string path { get; set; }
    public Realm realm { get; set; }
    public Region region { get; set; }
}

public class Roster
{
    public Character character { get; set; }
    public object oldCharacter { get; set; }
    public bool isTransfer { get; set; }
    public string role { get; set; }
}

public class Run
{
    public string season { get; set; }
    public Dungeon dungeon { get; set; }
    public int keystone_run_id { get; set; }
    public int keystone_team_id { get; set; }
    public object keystone_platoon_id { get; set; }
    public int mythic_level { get; set; }
    public int clear_time_ms { get; set; }
    public int keystone_time_ms { get; set; }
    public DateTime completed_at { get; set; }
    public int num_chests { get; set; }
    public int time_remaining_ms { get; set; }
    public string faction { get; set; }
    public List<WeeklyModifier> weekly_modifiers { get; set; }
    public int num_modifiers_active { get; set; }
    public List<Roster> roster { get; set; }
    public object platoon { get; set; }
}

public class Ranking
{
    public int rank { get; set; }
    public double score { get; set; }
    public Run run { get; set; }
}

public class Params
{
    public string season { get; set; }
    public string region { get; set; }
    public string dungeon { get; set; }
}

public class RootObject
{
    public List<Ranking> rankings { get; set; }
    public string leaderboard_url { get; set; }
    public Params @params { get; set; }
}
*/
