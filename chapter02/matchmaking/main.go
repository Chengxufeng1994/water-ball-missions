package main

import "fmt"

func main() {
	john, _ := NewIndividual(1, "John", Male, 28, "I am John", "basketball,soccer,video game", Coord{0, 0})
	jane, _ := NewIndividual(2, "Jane", Female, 24, "I am Jane", "basketball,soccer,video game", Coord{10, 10})
	jenny, _ := NewIndividual(3, "Jenny", Female, 24, "I am Jenny", "swimming,booking,coffee", Coord{4, 8})
	kakarot, _ := NewIndividual(4, "Kakarot", Male, 30, "I am Kakarot", "basketball,booking,coffee", Coord{10, 10})
	allIndividuals := []Individual{john, jane, jenny, kakarot}

	matchMakingSystem := NewMatchmakingSystem(NewDistanceBasedStrategy())
	result := matchMakingSystem.Matching(john, allIndividuals)
	fmt.Println(result)

	matchMakingSystem = NewMatchmakingSystem(NewReverseStrategy(NewDistanceBasedStrategy()))
	result = matchMakingSystem.Matching(john, allIndividuals)
	fmt.Println(result)
}
