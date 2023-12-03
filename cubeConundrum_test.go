package main

import (
	"github.com/samber/lo"
	"reflect"
	. "strconv"
	"strings"
	"testing"
)

func TestSingleCubeSameColorGame(t *testing.T) {

	cubeSet := CubeSet{cubes: []Cube{Cube{num: 3, color: "red"}}}
	game := Game{id: 1, cubeSet: []CubeSet{cubeSet}}

	cube := Cube{num: 3, color: "red"}
	bag := Bag{cubeSet: []CubeSet{CubeSet{cubes: []Cube{cube}}}}

	assertEquals(game.isPossibleFor(bag), true, t)

}

func TestSingleCubeDifferentColor(t *testing.T) {

	cubeSet := CubeSet{cubes: []Cube{Cube{num: 3, color: "red"}}}
	game := Game{id: 1, cubeSet: []CubeSet{cubeSet}}

	cube := Cube{num: 3, color: "green"}
	bag := Bag{cubeSet: []CubeSet{CubeSet{cubes: []Cube{cube}}}}

	assertEquals(game.isPossibleFor(bag), false, t)

}

func TestMultiCubesSameColorGame(t *testing.T) {

	game := Game{id: 1, cubeSet: []CubeSet{CubeSet{cubes: []Cube{
		{num: 3, color: "red"},
		{num: 2, color: "green"}}}}}

	bag := Bag{cubeSet: []CubeSet{CubeSet{cubes: []Cube{
		{num: 3, color: "red"},
		{num: 3, color: "green"}}}}}

	assertEquals(game.isPossibleFor(bag), true, t)

}

func TestParseGameString(t *testing.T) {

	var gameParsed Game = ParseGameString("Game 95: 6 blue, 1 green; 3 red, 11 green; 4 blue")

	gameExpected := Game{id: 95, cubeSet: []CubeSet{
		{cubes: []Cube{{num: 6, color: "blue"}, {num: 1, color: "green"}}},
		{cubes: []Cube{{num: 3, color: "red"}, {num: 11, color: "green"}}},
		{cubes: []Cube{{num: 4, color: "blue"}}}}}

	if reflect.DeepEqual(gameParsed, gameExpected) == false {
		t.Errorf("expected %q got %q", gameExpected, gameParsed)
	}

}

func ParseGameString(s string) Game {
	var fragments = strings.Split(s, ":")
	gameId := parseGameId(fragments[0])
	cubeSets := parseCubeSets(fragments[1])
	return Game{id: gameId, cubeSet: cubeSets}
}

func parseCubeSets(s string) []CubeSet {
	cubesetsFragments := strings.Split(s, ";")

	var cubeSets []CubeSet

	for _, fragment := range cubesetsFragments {
		cubes := parseCubes(strings.Split(fragment, ","))
		cubeSets = append(cubeSets, CubeSet{cubes: cubes})
	}

	return cubeSets
}

func parseCubes(cubeFragments []string) []Cube {
	var cubes = []Cube{}
	for _, cubeFragment := range cubeFragments {
		cube := parseCube(cubeFragment)
		cubes = append(cubes, cube)
	}
	return cubes
}

func parseCube(s string) Cube {
	numAndColor := strings.Split(s, " ")
	return Cube{lo.Must(Atoi(numAndColor[1])), numAndColor[2]}
}

func parseGameId(s string) int {
	i := strings.Split(s, " ")

	gameId, _ := Atoi(i[1])
	return gameId
}

func assertEquals(res bool, expected bool, t *testing.T) {
	if res != expected {
		t.Errorf("expected %t got %t", expected, res)
	}
}
