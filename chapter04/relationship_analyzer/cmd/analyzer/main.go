package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/adpater"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/relationship_analyzer/internal/thirdparty"
)

func main() {
	scrip := "data/script.txt"
	name1 := "B"
	name2 := "C"
	analyzerV1 := adpater.NewRelationshipAnalyzerAdapter(*thirdparty.NewSuperRelationshipAnalyzer())
	analyzerV2 := adpater.NewGonumRelationshipAnalyzerAdapter()
	client := NewClient(analyzerV1, analyzerV2)
	mutualFriends := client.AnalyzeV1(scrip, name1, name2)
	fmt.Println(mutualFriends)

	hasConnection := client.AnalyzeV2(scrip, name1, name2)
	fmt.Println(hasConnection)
}

type Client struct {
	analyzerV1 domain.RelationshipAnalyzerV1
	analyzerV2 domain.RelationshipAnalyzerV2
}

func NewClient(analyzerV1 domain.RelationshipAnalyzerV1, analyzerV2 domain.RelationshipAnalyzerV2) *Client {
	return &Client{
		analyzerV1: analyzerV1,
		analyzerV2: analyzerV2,
	}
}

func (c *Client) AnalyzeV1(script, name1, name2 string) []string {
	data, err := os.ReadFile(script)
	if err != nil {
		log.Fatal(err)
	}
	c.analyzerV1.Parse(string(data))
	return c.analyzerV1.GetMutualFriends(name1, name2)
}

func (c *Client) AnalyzeV2(script, name1, name2 string) bool {
	data, err := os.ReadFile(script)
	if err != nil {
		log.Fatal(err)
	}
	graph := c.analyzerV2.Parse(string(data))
	return graph.HasConnection(name1, name2)
}
