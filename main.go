package main

import (
	"fmt"
	//"os"
)

func main() {
	fmt.Println("Rubiks Cube - Reposition last two corners")

	for maxturns := 0; maxturns < 6; maxturns++ {
		fmt.Println("Maxturns depth", maxturns+1)
		beast := NewTopLayer()
		//beast.InitializeTopLayerRightStartSwap()
		beast.InitializeTopLayerTwoCubesMisOriented()
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
			fmt.Println("....")
			// no break as there might be multiple solutions
			// break // break out of the loop
		}
	}
}

// TryMove is the type definining the four possible moves of the player
type TryMove int32

// CornerFace represents one visible faceside of a corner
// each corner has three visible of these faces
type CornerFace int32

const (
	RIGHTSTART_SWAP TryMove = 3
	LEFTSTART_SWAP  TryMove = 1
	RIGHT_TO_FRONT  TryMove = 2
	LEFT_TO_FRONT   TryMove = 0

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

var CornerFaceName = map[CornerFace]string{
	FACE_TOP:   "TOP",
	FACE_FRONT: "FRONT",
	FACE_RIGHT: "RIGHT",
	FACE_BACK:  "BACK",
	FACE_LEFT:  "LEFT",
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
// as if the cube is in a solved state
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

// InitializeTopLayerRightStartSwap is to initialize for a garbled
// up cube of top corner that is not default to correct....
func (t *TopLayer) InitializeTopLayerRightStartSwap() {
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

// InitializeTopLayerTwoCubesMisOriented to initialize top layer
// as if two cubies have their corners mis-oriented
func (t *TopLayer) InitializeTopLayerTwoCubesMisOriented() {
	t.corners = make([]Corner, 4, 4)

	// FrontRight cube - is misaligned as per example image
	t.corners[CUBE_FR].front = FACE_RIGHT
	t.corners[CUBE_FR].right = FACE_TOP
	t.corners[CUBE_FR].top = FACE_FRONT

	// FrontLeft cube
	t.corners[CUBE_FL].front = FACE_FRONT
	t.corners[CUBE_FL].left = FACE_LEFT
	t.corners[CUBE_FL].top = FACE_TOP

	// BackLeft cube
	t.corners[CUBE_BL].back = FACE_BACK
	t.corners[CUBE_BL].left = FACE_LEFT
	t.corners[CUBE_BL].top = FACE_TOP

	// BackRight cube - is misaligned as per example image
	t.corners[CUBE_BR].back = FACE_RIGHT
	t.corners[CUBE_BR].right = FACE_TOP
	t.corners[CUBE_BR].top = FACE_BACK

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
		// the corners is a slice, we have to make a deep copy of it
		beast.corners = make([]Corner, len(t.corners))
		copy(beast.corners, t.corners)
		beast.ExecuteMove(idx, turns)
		result := slayTopLayer(beast, turns+1, maxturns)
		if result.Solved() {
			// cube solved at this depth,
			// we should actually add this result
			// to some resultlist, so that we can
			// also check other possible solutions at the same depth
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
			t.corners[CUBE_FR].top = MoveRightRename(t.corners[CUBE_BR].top)
			t.corners[CUBE_FR].front = MoveRightRename(t.corners[CUBE_BR].right)
			t.corners[CUBE_FR].right = MoveRightRename(t.corners[CUBE_BR].back)
			t.corners[CUBE_BR].top = MoveRightRename(t.corners[CUBE_BL].top)
			t.corners[CUBE_BR].back = MoveRightRename(t.corners[CUBE_BL].left)
			t.corners[CUBE_BR].right = MoveRightRename(t.corners[CUBE_BL].back)
			t.corners[CUBE_BL].top = MoveRightRename(t.corners[CUBE_FL].top)
			t.corners[CUBE_BL].back = MoveRightRename(t.corners[CUBE_FL].left)
			t.corners[CUBE_BL].left = MoveRightRename(t.corners[CUBE_FL].front)
			t.corners[CUBE_FL].top = MoveRightRename(frcube.top)
			t.corners[CUBE_FL].front = MoveRightRename(frcube.right)
			t.corners[CUBE_FL].left = MoveRightRename(frcube.front)
			t.trackmoves[turn] = RIGHT_TO_FRONT
		}
	case LEFT_TO_FRONT:
		{ // turning the cube is tricky
			// because apart from asigning all the
			// faces, also the designations of the faces
			// have to be turned counter clockwise
			flcube := t.corners[CUBE_FL] // backup copy
			t.corners[CUBE_FL].top = MoveLeftRename(t.corners[CUBE_BL].top)
			t.corners[CUBE_FL].front = MoveLeftRename(t.corners[CUBE_BL].left)
			t.corners[CUBE_FL].left = MoveLeftRename(t.corners[CUBE_BL].back)
			t.corners[CUBE_BL].top = MoveLeftRename(t.corners[CUBE_BR].top)
			t.corners[CUBE_BL].back = MoveLeftRename(t.corners[CUBE_BR].right)
			t.corners[CUBE_BL].left = MoveLeftRename(t.corners[CUBE_BR].back)
			t.corners[CUBE_BR].top = MoveLeftRename(t.corners[CUBE_FR].top)
			t.corners[CUBE_BR].back = MoveLeftRename(t.corners[CUBE_FR].right)
			t.corners[CUBE_BR].right = MoveLeftRename(t.corners[CUBE_FR].front)
			t.corners[CUBE_FR].top = MoveLeftRename(flcube.top)
			t.corners[CUBE_FR].front = MoveLeftRename(flcube.left)
			t.corners[CUBE_FR].right = MoveLeftRename(flcube.front)
			t.trackmoves[turn] = LEFT_TO_FRONT

		}
	}
}

// MoveLeftRename where
// front becomes right, right becomes back, back becomes left, left becomes front
func MoveLeftRename(cf CornerFace) CornerFace {
	switch cf {
	case FACE_FRONT:
		return FACE_RIGHT
	case FACE_LEFT:
		return FACE_FRONT
	case FACE_BACK:
		return FACE_LEFT
	case FACE_RIGHT:
		return FACE_BACK
	}
	return cf // no translation for top needed
}

// MoveRightRename where
// front becomes left, right becomes front, back becomes right, left becomes back
func MoveRightRename(cf CornerFace) CornerFace {
	switch cf {
	case FACE_FRONT:
		return FACE_LEFT
	case FACE_LEFT:
		return FACE_BACK
	case FACE_BACK:
		return FACE_RIGHT
	case FACE_RIGHT:
		return FACE_FRONT
	}
	return cf // no translation for top needed
}

func (t *TopLayer) showMonster() {
	fmt.Printf("%+v", t)
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
