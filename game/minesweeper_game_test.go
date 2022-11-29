package game

import (
	"testing"
)

// Valid input parameters to initialize the game.
const (
	boardSize  = 8
	blackHoles = 8
)

// A test to validate user input data.
func TestInputValidation(t *testing.T) {
	validationErrMsg := "input must be invalid"

	// Invalid board size.
	valid := isBoardSizeValid(0)
	if valid {
		t.Errorf(validationErrMsg)
	}

	// Invalid number of black holes.
	// (must not be less than 1)
	valid = isBlackHolesCountValid(0, boardSize)
	if valid {
		t.Errorf(validationErrMsg)
	}

	// Invalid number of black holes.
	// (must not be greater than size*size)
	valid = isBlackHolesCountValid((boardSize*boardSize)+1, boardSize)
	if valid {
		t.Errorf(validationErrMsg)
	}

	// Valid board size.
	valid = isBoardSizeValid(boardSize)
	if !valid {
		t.Errorf("input validation failed")
	}

	// Valid number of black holes.
	valid = isBlackHolesCountValid(blackHoles, boardSize)
	if !valid {
		t.Errorf("input validation failed")
	}
}

// Tests the process of initializing a new game.
func TestGameInitialization(t *testing.T) {
	// Initialize the game.
	game := newGame(boardSize, blackHoles)

	// Check that the number of movesLeft to finish the game is correct.
	movesLeft := boardSize*boardSize - blackHoles
	if game.movesLeft != movesLeft {
		t.Errorf("got %d, wanted %d", game.movesLeft, movesLeft)
	}

	// Check that the generated board is not empty and contains no black holes.
	if len(game.board) == 0 {
		t.Errorf("board has not been initialized")
	}
	for _, cells := range game.board {
		for _, cell := range cells {
			if cell.hasBlackHole {
				t.Errorf("the board must not have black holes")
			}
		}
	}
}

// Tests that black holes have been placed on the game board.
func TestBlackHolesGeneration(t *testing.T) {
	// Initialize a valid game.
	game := newGame(boardSize, blackHoles)

	//Make sure the list of black holes is empty by default.
	if len(game.blackHoleCells) > 0 {
		t.Errorf("the list of black holes must be empty")
	}

	// Populate black holes.
	game.populateBlackHoles()

	// Verify the number of black holes is correct.
	if len(game.blackHoleCells) != blackHoles {
		t.Errorf("got %d, wanted %d", len(game.blackHoleCells), blackHoles)
	}

	// Ensure that the game board contains all generated black holes.
	boardBlackHoles := 0
	for _, cells := range game.board {
		for _, cell := range cells {
			if cell.hasBlackHole {
				boardBlackHoles++
			}
		}
	}
	if boardBlackHoles != blackHoles {
		t.Errorf("got %d, wanted %d", len(game.blackHoleCells), blackHoles)
	}
}

// Tests cell opening.
func TestOpenCell(t *testing.T) {
	// Initialize a valid game.
	game := newGame(boardSize, blackHoles)
	// Populate black holes.
	game.populateBlackHoles()

	// Get cells without black holes.
	cells := getCellsWithoutBlackHoles(game)
	movesLeft := game.movesLeft

	// Open cell.
	game.openCell(cells[0][0], cells[0][1])
	// Make sure the cell is open now.
	if !game.board[cells[0][0]][cells[0][1]].isOpen {
		t.Errorf("cell must be open")
	}
	// Make sure that the number of moves has decreased.
	if movesLeft == game.movesLeft {
		t.Errorf("number of moves should be less")
	}
}

// Tests opening the cell with the black hole.
func TestOpenCellWithBlackHole(t *testing.T) {
	// Initialize a valid game.
	game := newGame(boardSize, blackHoles)
	// Populate black holes.
	game.populateBlackHoles()

	// Open cell with black hole.
	row, col := game.blackHoleCells[0][0], game.blackHoleCells[0][1]
	game.openCell(row, col)

	// Verify that the game status is Lost now.
	if game.gameStatus != Lost {
		t.Errorf("got %d, wanted %d", game.gameStatus, Lost)
	}

	// Verify that all black holes are revealed.
	for _, cell := range game.blackHoleCells {
		if !game.board[cell[0]][cell[1]].isOpen {
			t.Errorf("cell must be open")
		}
	}
}

// Test that after opening all cells without black holes, the game is considered won.
func TestOpenAllCellsAndWin(t *testing.T) {
	// Initialize a valid game.
	game := newGame(boardSize, blackHoles)
	// Populate black holes.
	game.populateBlackHoles()

	// Get cells without black holes.
	cells := getCellsWithoutBlackHoles(game)
	for _, cell := range cells {
		game.openCell(cell[0], cell[1])
	}

	// Verify that the status of the game is Won after opening all the cells.
	if game.gameStatus != Won {
		t.Errorf("got %d, wanted %d", game.gameStatus, Won)
	}
}

// Tests that all adjacent cells have been obtained.
func TestGetSurroundingCells(t *testing.T) {
	// Initialize a valid game.
	game := newGame(boardSize, blackHoles)
	// Coordinates of the cell to check.
	row, col := 2, 3
	// Get all 8 adjacent cells.
	cells := game.getSurroundingCells(row, col)

	if cells[0][0] != row-1 || cells[0][1] != col-1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[0][0], cells[0][1], row-1, col-1)
	}
	if cells[1][0] != row-1 || cells[1][1] != col {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[1][0], cells[1][1], row-1, col)
	}
	if cells[2][0] != row-1 || cells[2][1] != col+1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[2][0], cells[2][1], row-1, col+1)
	}
	if cells[3][0] != row || cells[3][1] != col-1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[3][0], cells[3][1], row, col-1)
	}
	if cells[5][0] != row || cells[5][1] != col+1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[5][0], cells[5][1], row, col+1)
	}
	if cells[6][0] != row+1 || cells[6][1] != col-1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[6][0], cells[6][1], row+1, col-1)
	}
	if cells[7][0] != row+1 || cells[7][1] != col {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[7][0], cells[7][1], row+1, col)
	}
	if cells[8][0] != row+1 || cells[8][1] != col+1 {
		t.Errorf("got (%d, %d), wanted (%d, %d)", cells[8][0], cells[8][1], row+1, col+1)
	}
}

// Helper function to get all cells without black holes.
func getCellsWithoutBlackHoles(game *MinesweeperGame) [][]int {
	cells := make([][]int, 0)
	for i := 0; i < game.size; i++ {
		for j := 0; j < game.size; j++ {
			if !game.board[i][j].hasBlackHole {
				cells = append(cells, []int{i, j})
			}
		}
	}
	return cells
}
