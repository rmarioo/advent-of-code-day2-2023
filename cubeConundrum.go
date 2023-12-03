package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func ConundrumPartTwo(supplier LinesSupplier, bag Bag) int {

	lines := supplier.Lines()
	sum := 0
	for _, line := range lines {
		game := ParseGameString(line)

		var bag Bag = game.CalculateMinBag()
		var power int = bag.CalculatePower()

		sum += power

	}

	return sum
}

func (game Game) CalculateMinBag() Bag {

	greenNum, blueNum, redNum := game.calculateMinColors()
	return game.bagFromColors(greenNum, blueNum, redNum)

}

func ConundrumPartOne(supplier LinesSupplier, bag Bag) int {

	lines := supplier.Lines()
	sum := 0
	for _, line := range lines {
		game := ParseGameString(line)
		if game.isPossibleFor(bag) {
			sum += game.id
		}
	}

	return sum
}

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

func (game Game) bagFromColors(greenNum int, blueNum int, redNum int) Bag {
	var cubes = []Cube{}
	if greenNum > 0 {
		cubes = append(cubes, Cube{num: greenNum, color: "green"})
	}
	if blueNum > 0 {
		cubes = append(cubes, Cube{num: blueNum, color: "blue"})
	}
	if redNum > 0 {
		cubes = append(cubes, Cube{num: redNum, color: "red"})
	}

	bag := Bag{cubes: cubes}
	return bag
}

func (game Game) calculateMinColors() (int, int, int) {
	var greenNum = 0
	var blueNum = 0
	var redNum = 0

	for _, cubeset := range game.cubeSet {

		for _, cube := range cubeset.cubes {
			if cube.color == "red" && redNum < cube.num {
				redNum = cube.num
			}
			if cube.color == "green" && greenNum < cube.num {
				greenNum = cube.num
			}
			if cube.color == "blue" && blueNum < cube.num {
				blueNum = cube.num
			}
		}
	}
	return greenNum, blueNum, redNum
}

type Bag struct {
	cubes []Cube
}

func (b Bag) CalculatePower() int {

	var power = 1
	for _, cube := range b.cubes {
		power = power * cube.num
	}
	return power
}

func canFit(bag Bag, cube Cube) bool {
	bagCanFit := false

	for _, bagCube := range bag.cubes {

		if bagCube.color == cube.color && bagCube.num >= cube.num {
			bagCanFit = true
			break
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

type LinesSupplier interface {
	Lines() []string
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
