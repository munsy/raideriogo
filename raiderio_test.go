package raideriogo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-test/deep"
)

type testapi struct {
	Endpoint string
	Data     string
}

var mux *http.ServeMux
var server *httptest.Server
var client *RaiderIOClient

// /api/v1/characters/
var t1 = `{"name":"Munsy","race":"Troll","class":"Druid","active_spec_name":"Balance","active_spec_role":"DPS","gender":"female","faction":"horde","achievement_points":9110,"honorable_kills":5,"thumbnail_url":"https://render-us.worldofwarcraft.com/character/thrall/55/178454583-avatar.jpg?alt=wow/static/images/2d/avatar/8-1.jpg","region":"us","realm":"Thrall","profile_url":"https://raider.io/characters/us/thrall/Munsy"}`

// /api/v1/guilds/profile
var t2 = `{"name":"NoBelfsAllowed","faction":"horde","region":"us","realm":"Thrall","profile_url":"https://raider.io/guilds/us/thrall/NoBelfsAllowed","raid_rankings":{"the-emerald-nightmare":{"normal":{"world":0,"region":0,"realm":0},"heroic":{"world":0,"region":0,"realm":0},"mythic":{"world":0,"region":0,"realm":0}},"trial-of-valor":{"normal":{"world":0,"region":0,"realm":0},"heroic":{"world":0,"region":0,"realm":0},"mythic":{"world":0,"region":0,"realm":0}},"the-nighthold":{"normal":{"world":0,"region":0,"realm":0},"heroic":{"world":0,"region":0,"realm":0},"mythic":{"world":0,"region":0,"realm":0}},"tomb-of-sargeras":{"normal":{"world":0,"region":0,"realm":0},"heroic":{"world":0,"region":0,"realm":0},"mythic":{"world":0,"region":0,"realm":0}},"antorus-the-burning-throne":{"normal":{"world":20889,"region":8697,"realm":325},"heroic":{"world":15051,"region":6062,"realm":232},"mythic":{"world":0,"region":0,"realm":0}}},"raid_progression":{"the-emerald-nightmare":{"summary":"0/7 N","total_bosses":7,"normal_bosses_killed":0,"heroic_bosses_killed":0,"mythic_bosses_killed":0},"trial-of-valor":{"summary":"0/3 N","total_bosses":3,"normal_bosses_killed":0,"heroic_bosses_killed":0,"mythic_bosses_killed":0},"the-nighthold":{"summary":"0/10 N","total_bosses":10,"normal_bosses_killed":0,"heroic_bosses_killed":0,"mythic_bosses_killed":0},"tomb-of-sargeras":{"summary":"0/9 N","total_bosses":9,"normal_bosses_killed":0,"heroic_bosses_killed":0,"mythic_bosses_killed":0},"antorus-the-burning-throne":{"summary":"11/11 H","total_bosses":11,"normal_bosses_killed":11,"heroic_bosses_killed":11,"mythic_bosses_killed":0}}}`

// /api/v1/mythic-plus/affixes
var t3 = `{"region":"us","title":"Bolstering, Grievous, Tyrannical","leaderboard_url":"https://raider.io/mythic-plus/season-7.3.2/all/us/leaderboards-strict/affixes/bolstering-grievous-tyrannical","affix_details":[{"id":7,"name":"Bolstering","description":"When any non-boss enemy dies, its death cry empowers nearby allies, increasing their maximum health and damage by 20%.","wowhead_url":"https://wowhead.com/affix=7"},{"id":12,"name":"Grievous","description":"When injured below 90% health, players will suffer increasing damage over time until healed above 90% health.","wowhead_url":"https://wowhead.com/affix=12"},{"id":9,"name":"Tyrannical","description":"Boss enemies have 40% more health and inflict up to 15% increased damage.","wowhead_url":"https://wowhead.com/affix=9"}]}`

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = New()
}

func dispose() {
	server.Close()
}

func TestGetCharacterProfile(t *testing.T) {
	setup()
	defer dispose()

	mux.HandleFunc("/api/v1/characters/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, t1)
	})

	x, err := client.GetCharacterProfile("us", "thrall", "munsy", "")

	if nil != err {
		t.Fatal(err.Error())
	}

	y := &ViewCharacterProfileResponse{
		Name:              "Munsy",
		Race:              "Troll",
		Class:             "Druid",
		ActiveSpecName:    "Balance",
		ActiveSpecRole:    "DPS",
		Gender:            "female",
		Faction:           "horde",
		AchievementPoints: 9110,
		HonorableKills:    5,
		ThumbnailURL:      "https://render-us.worldofwarcraft.com/character/thrall/55/178454583-avatar.jpg?alt=wow/static/images/2d/avatar/8-1.jpg",
		Region:            "us",
		Realm:             "Thrall",
		ProfileURL:        "https://raider.io/characters/us/thrall/Munsy",
	}

	if diff := deep.Equal(x, y); diff != nil {
		t.Error(diff)
	}
}

func TestGetGuildProfile(t *testing.T) {
	setup()
	defer dispose()

	mux.HandleFunc("/api/v1/guilds/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, t2)
	})

	x, err := client.GetGuildProfile("us", "thrall", "NoBelfsAllowed", "raid_progression,raid_rankings")

	if nil != err {
		t.Fatal(err.Error())
	}

	y := &ViewGuildProfileResponse{
		Name:       "NoBelfsAllowed",
		Faction:    "horde",
		Region:     "us",
		Realm:      "Thrall",
		ProfileURL: "https://raider.io/guilds/us/thrall/NoBelfsAllowed",
		RaidRankings: RaidRankingsSchema{
			TheEmeraldNightmare: RaidDifficultyRankingsSchema{
				Normal: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Heroic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Mythic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
			},
			TrialOfValor: RaidDifficultyRankingsSchema{
				Normal: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Heroic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Mythic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
			},
			TheNighthold: RaidDifficultyRankingsSchema{
				Normal: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Heroic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Mythic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
			},
			TombOfSargeras: RaidDifficultyRankingsSchema{
				Normal: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Heroic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
				Mythic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
			},
			AntorusTheBurningThrone: RaidDifficultyRankingsSchema{
				Normal: RanksSchema{
					World:  20889,
					Region: 8697,
					Realm:  325,
				},
				Heroic: RanksSchema{
					World:  15052,
					Region: 6062,
					Realm:  232,
				},
				Mythic: RanksSchema{
					World:  0,
					Region: 0,
					Realm:  0,
				},
			},
		},
		RaidProgression: OverallRaidProgressionSchema{
			TheEmeraldNightmare: RaidProgressionSchema{
				Summary:            "0/7 N",
				TotalBosses:        7,
				NormalBossesKilled: 0,
				HeroicBossesKilled: 0,
				MythicBossesKilled: 0,
			},
			TrialOfValor: RaidProgressionSchema{
				Summary:            "0/3 N",
				TotalBosses:        3,
				NormalBossesKilled: 0,
				HeroicBossesKilled: 0,
				MythicBossesKilled: 0,
			},
			TheNighthold: RaidProgressionSchema{
				Summary:            "0/10 N",
				TotalBosses:        10,
				NormalBossesKilled: 0,
				HeroicBossesKilled: 0,
				MythicBossesKilled: 0,
			},
			TombOfSargeras: RaidProgressionSchema{
				Summary:            "0/9 N",
				TotalBosses:        9,
				NormalBossesKilled: 0,
				HeroicBossesKilled: 0,
				MythicBossesKilled: 0,
			},
			AntorusTheBurningThrone: RaidProgressionSchema{
				Summary:            "11/11 H",
				TotalBosses:        11,
				NormalBossesKilled: 11,
				HeroicBossesKilled: 11,
				MythicBossesKilled: 0,
			},
		},
	}

	if diff := deep.Equal(x, y); diff != nil {
		t.Error(diff)
	}
}

func TestGetMythicPlusAffixes(t *testing.T) {
	setup()
	defer dispose()

	mux.HandleFunc("/api/v1/guilds/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, t3)
	})

	x, err := client.GetMythicPlusAffixes("us", "en")

	if nil != err {
		t.Fatal(err.Error())
	}

	y := &ViewMythicPlusAffixesResponse{
		Region:         "us",
		Title:          "Bolstering, Grievous, Tyrannical",
		LeaderboardURL: "https://raider.io/mythic-plus/season-7.3.2/all/us/leaderboards-strict/affixes/bolstering-grievous-tyrannical",
		AffixDetails: []IDNameDescriptionWowheadModel{
			IDNameDescriptionWowheadModel{
				ID:          7,
				Name:        "Bolstering",
				Description: "When any non-boss enemy dies, its death cry empowers nearby allies, increasing their maximum health and damage by 20%.",
				WowheadURL:  "https://wowhead.com/affix=7",
			},
			IDNameDescriptionWowheadModel{
				ID:          12,
				Name:        "Grievous",
				Description: "When injured below 90% health, players will suffer increasing damage over time until healed above 90% health.",
				WowheadURL:  "https://wowhead.com/affix=12",
			},
			IDNameDescriptionWowheadModel{
				ID:          9,
				Name:        "Tyrannical",
				Description: "Boss enemies have 40% more health and inflict up to 15% increased damage.",
				WowheadURL:  "https://wowhead.com/affix=9",
			},
		},
	}

	if diff := deep.Equal(x, y); diff != nil {
		t.Error(diff)
	}
}

/* these are gonna be a pain. think of something better.
func TestGetMythicPlusRuns(t *testing.T) {

}

func TestGetRaidingBossRankings(t *testing.T) {

}

func TestGetRaidingHallOfFame(t *testing.T) {

}

func TestGetRaidingProgression(t *testing.T) {

}

func TestGetRaidingRaidRankings(t *testing.T) {

}
*/
