package raideriogo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var mux *http.ServeMux
var server *httptest.Server
var client *RaiderIOClient

var characterDetailsFullProfile = `{
  "name": "Munsy",
  "race": "Troll",
  "class": "Druid",
  "active_spec_name": "Balance",
  "active_spec_role": "DPS",
  "gender": "female",
  "faction": "horde",
  "achievement_points": 9110,
  "honorable_kills": 5,
  "thumbnail_url": "https://render-us.worldofwarcraft.com/character/thrall/55/178454583-avatar.jpg?alt=wow/static/images/2d/avatar/8-1.jpg",
  "region": "us",
  "realm": "Thrall",
  "profile_url": "https://raider.io/characters/us/thrall/Munsy"
}`

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
		fmt.Fprint(w, characterDetailsFullProfile)
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

	if !reflect.DeepEqual(x, y) {
		t.Fatal("DeepEqual() failed.")
	}
}

func TestGetGuildProfile(t *testing.T) {

}

func TestGetMythicPlusAffixes(t *testing.T) {

}

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
