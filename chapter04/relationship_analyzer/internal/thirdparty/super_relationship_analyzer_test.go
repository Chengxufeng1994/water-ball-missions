package thirdparty

import (
	"log"
	"os"
	"testing"
)

var superRelationshipAnalyzer *SuperRelationshipAnalyzer

func init() {
	superRelationshipAnalyzer = NewSuperRelationshipAnalyzer()
}

func TestSuperRelationshipAnalyzer(t *testing.T) {
	script, err := os.ReadFile("super_relationship_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	superRelationshipAnalyzer.Init(string(script))

	tests := []struct {
		targetName string
		name2      string
		name3      string
		want       bool
	}{
		{
			targetName: "A",
			name2:      "B",
			name3:      "C",
			want:       true,
		},
	}

	for _, tt := range tests {
		if got := superRelationshipAnalyzer.IsMutualFriends(tt.targetName, tt.name2, tt.name3); got != tt.want {
			t.Errorf("IsMutualFriends(%s, %s, %s) = %v, want %v", tt.targetName, tt.name2, tt.name3, got, tt.want)
		}
	}
}
