package raideriogo

import "testing"

func TestEndpointCharacter(t *testing.T) {
	var expectedURL = "https://raider.io/api/v1/characters/profile?region=us&realm=testRealm&name=testName"

	if EndpointCharacter("us", "testRealm", "testName") != expectedURL {
		t.Fatal("TestEndpointCharacter failed.")
	}
}
