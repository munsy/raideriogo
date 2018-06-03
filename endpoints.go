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

	// EndpointGuild defines the guild endpoint.
	EndpointGuild = func(region, realm, name, fields string) string {
		endpoint := endpointGuilds + "profile?region=" + region + "&realm=" + realm + "&name=" + name
		if fields != "" {
			endpoint += "&fields=" + fields
		}
		return endpoint
	}

	// EndpointMythicPlusAffixes defines the mythic+ affixes endpoint.
	EndpointMythicPlusAffixes = func(region, locale string) string {
		endpoint := endpointMythicPlus + "affixes?region=" + region
		if locale != "" {
			endpoint += "&locale=" + locale
		}
		return endpoint
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
	EndpointRaidingRaidRankings = func(raid, difficulty, region, realm string) string {
		endpoint := endpointRaiding + "raid-rankings?raid=" + raid + "&difficulty=" + difficulty + "&region=" + region
		if realm != "" {
			endpoint += "&realm=" + realm
		}
		return endpoint
	}
)
