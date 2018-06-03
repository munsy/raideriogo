package raideriogo

// APIVersion is the Raider.io version used for the REST and Websocket API.
var APIVersion = "1"

var (
	EndpointRaiderIO = "https://raider.io/"
	EndpointAPI = EndpointRaiderIO + "api/v" + APIVersion + "/"

	EndpointCharacters = EndpointAPI + "characters/"
	EndpointGuilds = EndpointAPI + "guilds/"
	EndpointMythicPlus = EndpointAPI + "mythic-plus/"
	EndpointRaiding = EndpointAPI + "raiding/"

	// Character Endpoints
	EndpointCharacter = func(region, realm, name string) string { return EndpointCharacters + "profile?region=" + region + "&realm=" + realm + "&name=" + name }
	EndpointCharacterGear = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=gear" }
	EndpointCharacterGuild = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=guild" }
	EndpointCharacterRaidProgression = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=raid_progression" }
	EndpointCharacterMythicPlusScores = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_scores" }
	EndpointCharacterMythicPlusRanks = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_ranks" }
	EndpointCharacterMythicPlusRecentRuns = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_recent_runs" }
	EndpointCharacterMythicPlusBestRuns = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_best_runs" }
	EndpointCharacterMythicPlusHighestLevelRuns = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_highest_level_runs" }
	EndpointCharacterMythicPlusWeeklyHighestLevelRuns = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=mythic_plus_weekly_highest_level_runs" }
	EndpointCharacterPreviousMythicPlusScores = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=previous_mythic_plus_scores" }
	EndpointCharacterPreviousMythicPlusRanks = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=previous_mythic_plus_ranks" }
	EndpointCharacterRaidAchievementMeta = func(region, realm, name string) string { return EndpointCharacter(region, realm, name) + "&fields=raid_achievement_meta" }

	// Guild Endpoints
	EndpointGuild = func(region, realm, name string) string { return EndpointGuilds + "profile?region=" + region + "&realm=" + realm + "&name=" + name }
	EndpointGuildRaidProgression = func(region, realm, name string) string { return EndpointGuild(region, realm, name) + "&fields=raid_progression" }
	EndpointGuildRaidRainkings = func(region, realm, name string) string { return EndpointGuild(region, realm, name) + "&fields=raid_rankings" }

	// Mythic Plus Endpoints
	EndpointMythicPlusAffixes = func(region, locale string) string { return EndpointMythicPlus + "affixes?region=" + region + "&locale=" + locale }
	EndpointMythicPlusRuns = func(season, region, dungeon, affixes string) string {
		endpoint := EndpointMythicPlus
		if season != "" {
			endpoint += "runs?season=" + season
		}
		if region == "" && dungeon == "" && affixes == "" {
			return endpoint
		} else {
			if region != "" {
				endpoint += "&region=" + region
			}
			if dungeon != "" {
				endpoint += "&dungeon=" + dungeon
			}
			if affixes != "" {
				endpoint += "&affixes=" + affixes
			}
		}
		return endpoint
	}

	// Raiding Endpoints
	EndpointRaidingBossRankings = func(raid, boss, difficulty, region, realm string) string {
		endpoint := EndpointRaiding + "boss-rankings?raid=" + raid + "&boss=" + boss + "&difficulty=" + difficulty + "&region=" + region
		if realm != "" {
			endpoint += "&realm=" + realm
		}
		return endpoint
	}
	EndpointRaidingHallOfFame = func(raid, difficulty, region string) string { return EndpointRaiding + "hall-of-fame?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region }
	EndpointRaidingProgression = func(raid, difficulty, region string) string { return EndpointRaiding + "progression?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region }
	EndpointRaidingRaidRankings = func(raid, boss, difficulty, region, realm string) string {
                endpoint := EndpointRaiding + "raid-rankings?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region
                if realm != "" {
                        endpoint += "&realm=" + realm
                }
                return endpoint
        }
)















