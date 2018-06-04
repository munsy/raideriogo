package raideriogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RaiderIOClient struct {
}

func New() *RaiderIOClient {
	return &RaiderIOClient{}
}

func (rio *RaiderIOClient) GetCharacterProfile(region, realm, name, fields string) error {
	var characterProfile *ViewCharacterProfileResponse

	return get(EndpointCharacter(region, realm, name, fields), characterProfile)
}

/*
func (rio *RaiderIOClient) GetGuildProfile(region, realm, name, fields string) (ViewGuildProfileResponse, error) {
	guildProfile, err := get(EndpointGuild(region, realm, name, fields))

	if nil != err {
		return nil, err
	}

	return guildProfile, nil
}

func (rio *RaiderIOClient) GetMythicPlusAffixes(region, locale string) (ViewMythicPlusAffixesResponse, error) {
	affixes, err := get(EndpointMythicPlusAffixes(region, locale))

	if nil != err {
		return nil, err
	}

	return affixes, nil
}

func (rio *RaiderIOClient) GetMythicPlusRuns(season, region, dungeon, affixes string) (ViewMythicPlusRunsResponse, error) {
	runs, err := get(EndpointMythicPlusRuns(season, region, dungeon, affixes))

	if nil != err {
		return nil, err
	}

	return runs, nil
}

func (rio *RaiderIOClient) GetRaidingBossRankings(raid, boss, difficulty, region, realm string) (ViewRaidingBossRankingsResponse, error) {
	bossRankings, err := get(EndpointRaidingBossRankings(raid, boss, difficulty, region, realm))

	if nil != err {
		return nil, err
	}

	return bossRankings, nil
}

func (rio *RaiderIOClient) GetRaidingHallOfFame(raid, difficulty, region string) (ViewRaidingHallOfFameResponse, error) {
	hallOfFame, err := get(EndpointRaidingHallOfFame(raid, difficulty, region))

	if nil != err {
		return nil, err
	}

	return hallOfFame, nil
}

func (rio *RaiderIOClient) GetRaidingProgression(raid, difficulty, region string) (ViewRaidingProgressionResponse, error) {
	progression, err := get(EndpointRaidingProgression(raid, difficulty, region))

	if nil != err {
		return nil, err
	}

	return progression, nil
}

func (rio *RaiderIOClient) GetRaidingRaidRankings(raid, difficulty, region, realm string) (ViewRaidingRaidRankingsResponse, error) {
	rankings, err := get(EndpointRaidingRaidRankings(raid, difficulty, region, realm))

	if nil != err {
		return nil, err
	}

	return rankings, nil
}
*/

// Returns an interface corresponding to the given endpoint.
func get(endpoint string, v interface{}) error {
	response, err := http.Get(endpoint)
	if nil != err {
		return err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return err
	}

	return json.Unmarshal(data, v)
}
