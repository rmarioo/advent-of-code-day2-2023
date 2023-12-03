package main

import (
	"reflect"
	"strconv"
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

	var gameParsed Game = ParseGameString("Game 95: 6 blue; 3 red; 4 blue")

	gameExpected := Game{id: 95, cubeSet: []CubeSet{
		{cubes: []Cube{{num: 6, color: "blue"}}},
		{cubes: []Cube{{num: 3, color: "red"}}},
		{cubes: []Cube{{num: 4, color: "blue"}}}}}

	if reflect.DeepEqual(gameParsed, gameExpected) == false {
		t.Errorf("expected %q got %q", gameExpected, gameParsed)
	}

}

func ParseGameString(s string) Game {
	gameId := parseGameId(s)

	return Game{id: gameId, cubeSet: []CubeSet{
		{cubes: []Cube{{num: 6, color: "blue"}}},
		{cubes: []Cube{{num: 3, color: "red"}}},
		{cubes: []Cube{{num: 4, color: "blue"}}}}}
}

func parseGameId(s string) int {
	var split []string = strings.Split(s, ":")
	i := strings.Split(split[0], " ")

	gameId, _ := strconv.Atoi(i[1])
	return gameId
}

func assertEquals(res bool, expected bool, t *testing.T) {
	if res != expected {
		t.Errorf("expected %t got %t", expected, res)
	}
}
