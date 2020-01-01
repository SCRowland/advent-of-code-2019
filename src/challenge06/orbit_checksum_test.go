package challenge06

import "testing"

func TestExampleChecksum(t *testing.T) {
	orbitMap := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

	checksum := OrbitCountChecksum(orbitMap)
	if checksum != 42 {
		t.Errorf("GetChecksum(orbitMap) = %d, not 42", checksum)
	}

}

var orbitMaps = []struct {
	orbitMap         string
	expectedCheckSum int
}{
	{
		"COM)B",
		1,
	},
	{
		`COM)B
B)C
C)D`,
		6,
	},
	{
		exampleInput,
		42,
	},
	{
		puzzleInput,
		130681,
	},
}

func TestDirectOrbits(t *testing.T) {
	for _, testItem := range orbitMaps {
		got := OrbitCountChecksum(testItem.orbitMap)
		if got != testItem.expectedCheckSum {
			t.Errorf("GetChecksum(orbitMap) = %d, not %d", got, testItem.expectedCheckSum)
		}
	}
}

var orbitalTransferMaps = []struct {
	orbitMap  string
	transfers int
}{
	{
		exampleOrbitalTransferMap,
		exampleOrbitalTransferDistance,
	},
	{
		puzzleInput,
		313,
	},
}

func TestOrbitalTransfer(t *testing.T) {
	for _, testItem := range orbitalTransferMaps {
		got := MinimumOrbitalTransferCount(testItem.orbitMap)
		if got != testItem.transfers {
			t.Errorf("MinimumOrbitalTransferCount() = %d not %d", got, testItem.transfers)
		}
	}
}
