package raideriogo

import "testing"

func TestEndpointCharacterNoFields(t *testing.T) {
	var expectedURL = "https://raider.io/api/v1/characters/profile?region=testRegion&realm=testRealm&name=testName"

	if EndpointCharacter("testRegion", "testRealm", "testName", "") != expectedURL {
		t.Fatal("String mismatch.")
	}
}

func TestEndpointCharacterWithFields(t *testing.T) {
	var expectedURL = "https://raider.io/api/v1/characters/profile?region=testRegion&realm=testRealm&name=testName&fields=gear%2Cguild%2Cmythic_plus_scores"

	if EndpointCharacter("testRegion", "testRealm", "testName", "gear,guild,mythic_plus_scores") != expectedURL {
		t.Fatal("String mismatch.")
	}
}
