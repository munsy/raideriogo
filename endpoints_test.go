package raideriogo

import "testing"

func TestEndpointCharacter(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/characters/profile?region=testRegion&realm=testRealm&name=testName",
		"https://raider.io/api/v1/characters/profile?region=testRegion&realm=testRealm&name=testName&fields=gear,guild,mythic_plus_scores",
	}

	if EndpointCharacter("testRegion", "testRealm", "testName", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointCharacter("testRegion", "testRealm", "testName", "gear,guild,mythic_plus_scores") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointGuild(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/guilds/profile?region=testRegion&realm=testRealm&name=testName",
		"https://raider.io/api/v1/guilds/profile?region=testRegion&realm=testRealm&name=testName&fields=raid_progression,raid_rankings",
	}

	if EndpointGuild("testRegion", "testRealm", "testName", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointGuild("testRegion", "testRealm", "testName", "raid_progression,raid_rankings") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointMythicPlusAffixes(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/mythic-plus/affixes?region=testRegion",
		"https://raider.io/api/v1/mythic-plus/affixes?region=testRegion&locale=us",
	}

	if EndpointMythicPlusAffixes("testRegion", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointMythicPlusAffixes("testRegion", "us") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointMythicPlusRuns(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/mythic-plus/runs",
		"https://raider.io/api/v1/mythic-plus/runs?season=season-7.3.0",
		"https://raider.io/api/v1/mythic-plus/runs?season=season-7.3.0&region=us",
		"https://raider.io/api/v1/mythic-plus/runs?season=season-7.3.0&region=us&dungeon=eye-of-azshara",
		"https://raider.io/api/v1/mythic-plus/runs?season=season-7.3.0&region=us&dungeon=eye-of-azshara&affixes=bolstering",
	}

	if EndpointMythicPlusRuns("", "", "", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointMythicPlusRuns("season-7.3.0", "", "", "") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
	if EndpointMythicPlusRuns("season-7.3.0", "us", "", "") != expectedURL[2] {
		t.Fatal("String mismatch.")
	}
	if EndpointMythicPlusRuns("season-7.3.0", "us", "eye-of-azshara", "") != expectedURL[3] {
		t.Fatal("String mismatch.")
	}
	if EndpointMythicPlusRuns("season-7.3.0", "us", "eye-of-azshara", "bolstering") != expectedURL[4] {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointRaidingBossRankings(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/raiding/boss-rankings?raid=tomb-of-sargeras&boss=goroth&difficulty=mythic&region=us",
		"https://raider.io/api/v1/raiding/boss-rankings?raid=tomb-of-sargeras&boss=goroth&difficulty=mythic&region=us&realm=earthen-ring",
		"https://raider.io/api/v1/raiding/boss-rankings?raid=tomb-of-sargeras&boss=goroth&difficulty=mythic&region=us&realm=connected-earthen-ring",
	}

	if EndpointRaidingBossRankings("tomb-of-sargeras", "goroth", "mythic", "us", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointRaidingBossRankings("tomb-of-sargeras", "goroth", "mythic", "us", "earthen-ring") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
	if EndpointRaidingBossRankings("tomb-of-sargeras", "goroth", "mythic", "us", "connected-earthen-ring") != expectedURL[2] {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointRaidingHallOfFame(t *testing.T) {
	var expectedURL = "https://raider.io/api/v1/raiding/hall-of-fame?raid=tomb-of-sargeras&difficulty=mythic&region=us"

	if EndpointRaidingHallOfFame("tomb-of-sargeras", "mythic", "us") != expectedURL {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointRaidingProgression(t *testing.T) {
	var expectedURL = "https://raider.io/api/v1/raiding/progression?raid=tomb-of-sargeras&difficulty=mythic&region=us"

	if EndpointRaidingProgression("tomb-of-sargeras", "mythic", "us") != expectedURL {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointRaidingRaidRankings(t *testing.T) {
	expectedURL := [...]string{
		"https://raider.io/api/v1/raiding/raid-rankings?raid=tomb-of-sargeras&difficulty=mythic&region=us",
		"https://raider.io/api/v1/raiding/raid-rankings?raid=tomb-of-sargeras&difficulty=mythic&region=us&realm=earthen-ring",
		"https://raider.io/api/v1/raiding/raid-rankings?raid=tomb-of-sargeras&difficulty=mythic&region=us&realm=connected-earthen-ring",
	}

	if EndpointRaidingRaidRankings("tomb-of-sargeras", "mythic", "us", "") != expectedURL[0] {
		t.Fatal("String mismatch.")
	}
	if EndpointRaidingRaidRankings("tomb-of-sargeras", "mythic", "us", "earthen-ring") != expectedURL[1] {
		t.Fatal("String mismatch.")
	}
	if EndpointRaidingRaidRankings("tomb-of-sargeras", "mythic", "us", "connected-earthen-ring") != expectedURL[2] {
		t.Fatal("String mismatch.")
	}
}
