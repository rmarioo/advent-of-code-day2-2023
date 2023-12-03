package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func (game Game) isPossibleFor(bag Bag) bool {

	same := true
	for _, cubeset := range game.cubeSet {

		for _, cube := range cubeset.cubes {

			if canFit(bag, cube) == false {
				same = false
				break
			}
		}

	}
	return same
}

type Bag struct {
	cubeSet []CubeSet
}

func canFit(bag Bag, cube Cube) bool {
	bagCanFit := false
	for _, cubeset := range bag.cubeSet {

		for _, bagCube := range cubeset.cubes {

			if bagCube.color == cube.color && bagCube.num >= cube.num {
				bagCanFit = true
				break
			}
		}

	}
	return bagCanFit
}

type Cube struct {
	num   int
	color string
}

type CubeSet struct {
	cubes []Cube
}

type Game struct {
	id      int
	cubeSet []CubeSet
}

type LinesSupplier interface {
	Lines() []string
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
	return Cube{lo.Must(strconv.Atoi(numAndColor[1])), numAndColor[2]}
}

func parseGameId(s string) int {
	i := strings.Split(s, " ")

	gameId, _ := strconv.Atoi(i[1])
	return gameId
}

type FromFileLinesSupplier struct {
	fileName string
}

func (f FromFileLinesSupplier) Lines() []string {
	return ReadFileLines(f.fileName)
}
func ReadFileLines(fileName string) []string {

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
