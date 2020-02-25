package main

import (
	"fmt"
	//"os"
)

func main() {
	fmt.Println("Repositiop corners")

	for maxturns := 1; maxturns < 3; maxturns++ {
		fmt.Println("Maxturns depth", maxturns)
		beast := NewTopLayer()
		beast.InitializeTopLayerCustom()
		if beast.Solved() {
			beast.showMonster()
			break
		}
		//beast.showMonster()
		*beast = slayTopLayer(*beast, 0, maxturns)
		if beast.Solved() {
			// we're done :)
			fmt.Println("Yes! We slayed the monster!")
			beast.PrintMoves(maxturns)
			break // break out of the loop
		}
	}
}

// TryMove is the type definining the four possible moves of the player
type TryMove int32

// CornerFace represents one visible faceside of a corner
// each corner has three visible of these faces
type CornerFace int32

const (
	RIGHTSTART_SWAP TryMove = 0
	LEFTSTART_SWAP  TryMove = 1
	RIGHT_TO_FRONT  TryMove = 2
	LEFT_TO_FRONT   TryMove = 3

	FACE_TOP   CornerFace = 0
	FACE_RIGHT CornerFace = 1
	FACE_LEFT  CornerFace = 2
	FACE_BACK  CornerFace = 3
	FACE_FRONT CornerFace = 4

	CUBE_FR = 0 // cube front right
	CUBE_FL = 1 // cube front left
	CUBE_BL = 2 // cube back left
	CUBE_BR = 3 // cube back right
)

var Move_name = map[TryMove]string{
	LEFT_TO_FRONT:   "LEFT_TO_FRONT",
	LEFTSTART_SWAP:  "LEFTSTART_SWAP",
	RIGHT_TO_FRONT:  "RIGHT_TO_FRONT",
	RIGHTSTART_SWAP: "RIGHTSTART_SWAP",
}

// Corner represents how the cornercube looks
// and what color-face it has from the front perspective
type Corner struct {
	front CornerFace
	top   CornerFace
	left  CornerFace
	right CornerFace
	back  CornerFace
}

// TopLayer holds the info of the top corner cubes
type TopLayer struct {
	corners    []Corner  // four corners, front-right first and clockwise
	trackmoves []TryMove // slice of tried moves
}

// NewTopLayer gives you a new monster,
// initialized with 3 heads and 3 tails
func NewTopLayer() *TopLayer {
	m := TopLayer{}
	m.InitializeTopLayerDefault()
	m.trackmoves = make([]TryMove, 12)
	return &m
}

// InitializeTopLayerDefault to initialize top layer
func (t *TopLayer) InitializeTopLayerDefault() {
	t.corners = make([]Corner, 4, 4)

	// FrontRight cube
	t.corners[CUBE_FR].front = FACE_FRONT
	t.corners[CUBE_FR].right = FACE_RIGHT
	t.corners[CUBE_FR].top = FACE_TOP

	// FrontLeft cube
	t.corners[CUBE_FL].front = FACE_FRONT
	t.corners[CUBE_FL].left = FACE_LEFT
	t.corners[CUBE_FL].top = FACE_TOP

	// BackLeft cube
	t.corners[CUBE_BL].back = FACE_BACK
	t.corners[CUBE_BL].left = FACE_LEFT
	t.corners[CUBE_BL].top = FACE_TOP

	// BackRight cube
	t.corners[CUBE_BR].back = FACE_BACK
	t.corners[CUBE_BR].right = FACE_RIGHT
	t.corners[CUBE_BR].top = FACE_TOP

}

// InitializeTopLayerCustom is to initialize for a garbled
// up cube of top corner that is not default to correct....
func (t *TopLayer) InitializeTopLayerCustom() {
	t.corners = make([]Corner, 4, 4)

	// FrontRight cube
	t.corners[CUBE_FR].front = FACE_LEFT
	t.corners[CUBE_FR].right = FACE_TOP
	t.corners[CUBE_FR].top = FACE_BACK

	// FrontLeft cube
	t.corners[CUBE_FL].front = FACE_FRONT
	t.corners[CUBE_FL].left = FACE_LEFT
	t.corners[CUBE_FL].top = FACE_TOP

	// BackLeft cube
	t.corners[CUBE_BL].back = FACE_BACK
	t.corners[CUBE_BL].left = FACE_TOP
	t.corners[CUBE_BL].top = FACE_RIGHT

	// BackRight cube
	t.corners[CUBE_BR].back = FACE_TOP
	t.corners[CUBE_BR].right = FACE_RIGHT
	t.corners[CUBE_BR].top = FACE_FRONT

}

func slayTopLayer(t TopLayer, turns int, maxturns int) TopLayer {
	//fmt.Print("Turn:", turns)
	//t.showMonster()
	if turns > maxturns || t.Solved() {
		// we either did not succeed in the amount of time allotted,
		// or we lost, or we won. Either way we are at the end of our tries:
		//fmt.Print(".")
		return t
	}
	for idx := range Move_name {
		beast := t
		beast.ExecuteMove(idx, turns)
		result := slayTopLayer(beast, turns+1, maxturns)
		if result.Solved() {
			return result
		}
	}

	return t
}

// ExecuteMove executes possible moves and transforms the faces of the cubes with move TryMove and
// ensures the moves is tracked in the trackmoves list at place 'turn'
func (t *TopLayer) ExecuteMove(move TryMove, turn int) {
	// fmt.Println(Move_name[move])
	switch move {
	case LEFTSTART_SWAP:
		{
			// FrontLeft cube
			flcube := t.corners[CUBE_FL] // backup copy
			t.corners[CUBE_FL].front = t.corners[CUBE_BR].right
			t.corners[CUBE_FL].left = t.corners[CUBE_BR].top
			t.corners[CUBE_FL].top = t.corners[CUBE_BR].back

			t.corners[CUBE_BR].back = t.corners[CUBE_BL].back
			t.corners[CUBE_BR].right = t.corners[CUBE_BL].top
			t.corners[CUBE_BR].top = t.corners[CUBE_BL].left

			t.corners[CUBE_BL].back = flcube.top
			t.corners[CUBE_BL].left = flcube.left
			t.corners[CUBE_BL].top = flcube.front
			t.trackmoves[turn] = LEFTSTART_SWAP
		}
	case RIGHTSTART_SWAP:
		{
			// FrontRight cube
			frcube := t.corners[CUBE_FR] // backup copy
			t.corners[CUBE_FR].front = t.corners[CUBE_BL].left
			t.corners[CUBE_FR].right = t.corners[CUBE_BL].top
			t.corners[CUBE_FR].top = t.corners[CUBE_BL].back

			t.corners[CUBE_BL].back = t.corners[CUBE_BR].back
			t.corners[CUBE_BL].left = t.corners[CUBE_BR].top
			t.corners[CUBE_BL].top = t.corners[CUBE_BR].right

			t.corners[CUBE_BR].back = frcube.top
			t.corners[CUBE_BR].right = frcube.right
			t.corners[CUBE_BR].top = frcube.front
			t.trackmoves[turn] = RIGHTSTART_SWAP
		}
	case RIGHT_TO_FRONT:
		{
			frcube := t.corners[CUBE_FR] // backup copy
			t.corners[CUBE_FR].front = t.corners[CUBE_BR].right
			t.corners[CUBE_FR].right = t.corners[CUBE_BR].back
			t.corners[CUBE_BR].back = t.corners[CUBE_BL].left
			t.corners[CUBE_BR].right = t.corners[CUBE_BL].back
			t.corners[CUBE_BL].back = t.corners[CUBE_FL].left
			t.corners[CUBE_BL].left = t.corners[CUBE_FL].front
			t.corners[CUBE_FL].front = frcube.right
			t.corners[CUBE_FL].left = frcube.front
			t.trackmoves[turn] = RIGHT_TO_FRONT
		}
	case LEFT_TO_FRONT:
		{
			flcube := t.corners[CUBE_FL] // backup copy
			t.corners[CUBE_FL].front = t.corners[CUBE_BL].left
			t.corners[CUBE_FL].left = t.corners[CUBE_BL].back
			t.corners[CUBE_BL].back = t.corners[CUBE_BR].right
			t.corners[CUBE_BL].left = t.corners[CUBE_BR].back
			t.corners[CUBE_BR].back = t.corners[CUBE_FR].right
			t.corners[CUBE_BR].right = t.corners[CUBE_FR].front
			t.corners[CUBE_FR].front = flcube.left
			t.corners[CUBE_FR].right = flcube.front
			t.trackmoves[turn] = LEFT_TO_FRONT
		}
	}
}

func (m *TopLayer) showMonster() {
	fmt.Printf("%+v", m)
}

// Solved tells if all the cubes are properly positioned
func (t *TopLayer) Solved() bool {

	if // FrontRight cube
	t.corners[CUBE_FR].front == FACE_FRONT &&
		t.corners[CUBE_FR].right == FACE_RIGHT &&
		t.corners[CUBE_FR].top == FACE_TOP &&

		// FrontLeft cube
		t.corners[CUBE_FL].front == FACE_FRONT &&
		t.corners[CUBE_FL].left == FACE_LEFT &&
		t.corners[CUBE_FL].top == FACE_TOP &&

		// BackLeft cube
		t.corners[CUBE_BL].back == FACE_BACK &&
		t.corners[CUBE_BL].left == FACE_LEFT &&
		t.corners[CUBE_BL].top == FACE_TOP &&

		// BackRight cube
		t.corners[CUBE_BR].back == FACE_BACK &&
		t.corners[CUBE_BR].right == FACE_RIGHT &&
		t.corners[CUBE_BR].top == FACE_TOP {
		return true
	}
	return false

}

// PrintMoves is to be called when monster is slain to show what moves let to this
// the program stores earlier moves in the trackmoves list
func (t *TopLayer) PrintMoves(turns int) {
	printTopLayer := NewTopLayer()
	for i := 0; i <= turns; i++ {
		fmt.Printf("%+v", *printTopLayer)
		fmt.Println(Move_name[TryMove(t.trackmoves[i])])
		printTopLayer.ExecuteMove(t.trackmoves[i], i)

	}
}
