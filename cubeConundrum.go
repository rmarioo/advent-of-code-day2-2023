package main

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
