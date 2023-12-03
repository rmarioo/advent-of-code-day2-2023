package main

import (
	"reflect"
	"testing"
)

func TestSingleCubeSameColorGame(t *testing.T) {

	cubeSet := CubeSet{cubes: []Cube{Cube{num: 3, color: "red"}}}
	game := Game{id: 1, cubeSet: []CubeSet{cubeSet}}

	bag := Bag{cubes: []Cube{{num: 3, color: "red"}}}

	assertEquals(game.isPossibleFor(bag), true, t)

}

func TestSingleCubeDifferentColor(t *testing.T) {

	cubeSet := CubeSet{cubes: []Cube{Cube{num: 3, color: "red"}}}
	game := Game{id: 1, cubeSet: []CubeSet{cubeSet}}

	bag := Bag{cubes: []Cube{{num: 3, color: "green"}}}

	assertEquals(game.isPossibleFor(bag), false, t)

}

func TestMultiCubesSameColorGame(t *testing.T) {

	game := Game{id: 1, cubeSet: []CubeSet{CubeSet{cubes: []Cube{
		{num: 3, color: "red"},
		{num: 2, color: "green"}}}}}

	bag := Bag{cubes: []Cube{
		{num: 3, color: "red"},
		{num: 3, color: "green"}}}

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

/*func TestConundrum(t *testing.T) {

	var supplier LinesSupplier = StubLinesSupplier{lines: []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}}

	var res int = Conundrum(supplier)

	assertEqualsInt(res, 2, t)
}

func Conundrum(supplier LinesSupplier) int {

	//lines := supplier.Lines()

	return 1
}*/

type StubLinesSupplier struct {
	lines []string
}

func (f StubLinesSupplier) Lines() []string {
	return f.lines
}

func assertEquals(res bool, expected bool, t *testing.T) {
	if res != expected {
		t.Errorf("expected %t got %t", expected, res)
	}
}

func assertEqualsInt(res int, expected int, t *testing.T) {
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}
