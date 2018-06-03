package raideriogo

// APIVersion is the Raider.io version used for the REST and Websocket API.
var APIVersion = "1"

var (
	endpointRaiderIO = "https://raider.io/"
	endpointAPI      = endpointRaiderIO + "api/v" + APIVersion + "/"

	endpointCharacters = endpointAPI + "characters/"
	endpointGuilds     = endpointAPI + "guilds/"
	endpointMythicPlus = endpointAPI + "mythic-plus/"
	endpointRaiding    = endpointAPI + "raiding/"

	// EndpointCharacter defines the character endpoint.
	EndpointCharacter = func(region, realm, name, fields string) string {
		endpoint := endpointCharacters + "profile?region=" + region + "&realm=" + realm + "&name=" + name
		if fields != "" {
			endpoint += "&fields=" + fields
		}
		return endpoint
	}
	/*
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
	*/

	// EndpointGuild defines the guild endpoint.
	EndpointGuild = func(region, realm, name string) string {
		return endpointGuilds + "profile?region=" + region + "&realm=" + realm + "&name=" + name
	}

	// EndpointGuildRaidProgression defines the guild raid progression endpoint.
	EndpointGuildRaidProgression = func(region, realm, name string) string {
		return EndpointGuild(region, realm, name) + "&fields=raid_progression"
	}

	// EndpointGuildRaidRainkings defines the guild raid ranking endpoint.
	EndpointGuildRaidRainkings = func(region, realm, name string) string {
		return EndpointGuild(region, realm, name) + "&fields=raid_rankings"
	}

	// EndpointMythicPlusAffixes defines the mythic+ affixes endpoint.
	EndpointMythicPlusAffixes = func(region, locale string) string {
		return endpointMythicPlus + "affixes?region=" + region + "&locale=" + locale
	}

	// EndpointMythicPlusRuns defines the mythic+ run endpoint.
	EndpointMythicPlusRuns = func(season, region, dungeon, affixes string) string {
		endpoint := endpointMythicPlus
		if season != "" {
			endpoint += "runs?season=" + season
		}
		if region == "" && dungeon == "" && affixes == "" {
			return endpoint
		}
		if region != "" {
			endpoint += "&region=" + region
		}
		if dungeon != "" {
			endpoint += "&dungeon=" + dungeon
		}
		if affixes != "" {
			endpoint += "&affixes=" + affixes
		}
		return endpoint
	}

	// EndpointRaidingBossRankings defines the raiding endpoint.
	EndpointRaidingBossRankings = func(raid, boss, difficulty, region, realm string) string {
		endpoint := endpointRaiding + "boss-rankings?raid=" + raid + "&boss=" + boss + "&difficulty=" + difficulty + "&region=" + region
		if realm != "" {
			endpoint += "&realm=" + realm
		}
		return endpoint
	}

	// EndpointRaidingHallOfFame defines the raiding hall of fame endpoint.
	EndpointRaidingHallOfFame = func(raid, difficulty, region string) string {
		return endpointRaiding + "hall-of-fame?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region
	}

	// EndpointRaidingProgression defines the raiding progression endpoint.
	EndpointRaidingProgression = func(raid, difficulty, region string) string {
		return endpointRaiding + "progression?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region
	}

	// EndpointRaidingRaidRankings defines raid rankings endpoint.
	EndpointRaidingRaidRankings = func(raid, boss, difficulty, region, realm string) string {
		endpoint := endpointRaiding + "raid-rankings?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region
		if realm != "" {
			endpoint += "&realm=" + realm
		}
		return endpoint
	}
)
