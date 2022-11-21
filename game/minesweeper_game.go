package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// A Cell represents single cell on the game board.
type Cell struct {
	hasBlackHole       bool // contains a black hole
	isOpen             bool // is opened by player
	adjacentBlackHoles int  // the number of adjacent black holes
}

type GameStatus int

// Enum is used as an indicator of the current state of the game
const (
	Playing GameStatus = iota // game in progress
	Won                       // player won the game
	Lost                      // player lost the game
)

const (
	invalidBoardMsg      = "invalid board size"
	invalidBlackHolesMsg = "invalid number of black holes"
)

// A MinesweeperGame is a struct that represents the entire game
// and tracks its progress.
type MinesweeperGame struct {
	// size is the game board size (size x size).
	// The user sets the value of this field.
	size int
	// blackHoles is the number of black holes on the board
	// The user sets the value of this field.
	blackHoles int
	movesLeft  int // number of moves left to finish the game

	// board is the game board that contains free cells
	// and cells with black holes.
	board          [][]*Cell
	blackHoleCells [][]int // location of black holes

	GameStatus GameStatus // status of the game
}

// A Function to initialize the game.
// Creates a new instance of the game with the given board size
// and number of black holes.
// Returns a game object or an error if the board size is invalid (< 1)
// or the number of black holes is too large (> board size * board size) or too small (< 1)
func newGame(size int, blackHoles int) (*MinesweeperGame, error) {
	// Initialization of the game object.
	// Default game status is Playing.
	game := &MinesweeperGame{
		size:       size,
		blackHoles: blackHoles,
		GameStatus: Playing,
	}

	// Board size validation.
	// Returns an error if the size is invalid.
	if size < 1 {
		return nil, errors.New(invalidBoardMsg)
	}
	// Number of black holes validation.
	// Returns an error if the number of black holes is invalid.
	if blackHoles < 1 || blackHoles > size*size {
		return nil, errors.New(invalidBlackHolesMsg)
	}

	// Calculates the number of moves to finish the game.
	game.movesLeft = size*size - blackHoles
	game.initializeBoard() // initialization of a new game board

	return game, nil
}

// A function to place the black holes randomly on the game board.
func (s *MinesweeperGame) populateBlackHoles() {
	// Location of the black holes is empty by default
	s.blackHoleCells = make([][]int, s.blackHoles)
	// Marker to check where the black hole was placed on generated position.
	// Prevents a black hole from being assigned to the position
	// where one already exists.
	placedBlackHoles := make([]bool, s.size*s.size)

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	// Continue until all the black holes have been created.
	for i := 0; i < s.blackHoles; {
		randomNum := rnd.Intn(s.size * s.size)
		row := randomNum % s.size
		col := randomNum / s.size

		// Add black holes if there are no black holes
		// on the board at that location.
		if !placedBlackHoles[randomNum] {
			// Mark that black hole was placed on that position.
			placedBlackHoles[randomNum] = true
			// Place black hole on the board.
			s.board[row][col].hasBlackHole = true
			// Add the black hole position to the collection of black hole locations.
			s.blackHoleCells[i] = []int{row, col}

			// Iterate through adjacent cells to the current black hole.
			for _, nextCell := range s.getSurroundingCells(row, col) {
				// Increment the number of adjacent black holes
				// for the adjacent cell with the current black hole.
				s.board[nextCell[0]][nextCell[1]].adjacentBlackHoles++
			}

			i++
		}
	}
}

// A function to open a cell selected by the user on the game board.
func (mg *MinesweeperGame) openCell(row int, col int) {
	// User has opened a black hole.
	// Game over and user lost.
	if mg.board[row][col].hasBlackHole {
		// Reveal all the black holes because user lost.
		mg.revalAllBlackHoles()
		mg.GameStatus = Lost // game status is lost
		return
	}

	mg.board[row][col].isOpen = true // selected cell is now open
	mg.movesLeft--                   // decrease the number of moves to finish the game

	// If a cell has zero adjacent black holes
	// the game needs to automatically make the surrounding cells open.
	// (A breadth-first search (BFS) algorithm is used
	// to traverse all possible surrounding cells)
	if mg.board[row][col].adjacentBlackHoles == 0 {
		// Get all adjacent cells with the current one
		// and create a queue for BFS.
		surroundingCells := mg.getSurroundingCells(row, col)
		for len(surroundingCells) > 0 {
			// Dequeue a cell from adjacent cells queue
			row, col := surroundingCells[0][0], surroundingCells[0][1]
			cell := mg.board[row][col]
			surroundingCells = surroundingCells[1:]

			// If a cell is already open, skip it
			// and go to the next one in the queue.
			if cell.isOpen {
				continue
			}

			cell.isOpen = true // selected cell is now open
			mg.movesLeft--     // decrease the number of moves to finish the game

			// If the cell has no adjacent black holes,
			// take all adjacent cells with it and add them to the queue.
			if cell.adjacentBlackHoles == 0 {
				surroundingCells = append(surroundingCells, mg.getSurroundingCells(row, col)...)
			}
		}
	}

	// If the number of moves is zero,
	// the game is considered over and user has won.
	if mg.movesLeft == 0 {
		mg.GameStatus = Won // game status is won
	}
}

// A function to print the current game board.
// Example of output:
//
//   0 1 2
// 0 * * *
// 1 * * *
// 2 * * *
//
// You can set the "debug" parameter to true to reval where all the black holes are.
// Example of output:
//
//   0 1 2
// 0 0 1 1
// 1 1 3 H
// 2 1 H H
//
// By default, the "debug" parameter should be false.
func (s *MinesweeperGame) printBoard(debug bool) {
	fmt.Println()
	fmt.Print("  ")

	for i := 0; i < s.size; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for i := 0; i < s.size; i++ {
		fmt.Printf("%d ", i)

		for j := 0; j < s.size; j++ {
			if debug {
				if s.board[i][j].hasBlackHole {
					fmt.Printf("%c ", 'H')
				} else {
					fmt.Printf("%d ", s.board[i][j].adjacentBlackHoles)
				}
			} else {
				if s.board[i][j].isOpen {
					fmt.Printf("%d ", s.board[i][j].adjacentBlackHoles)
				} else {
					fmt.Printf("* ")
				}
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

// A function to initialize a new game board.
func (s *MinesweeperGame) initializeBoard() {
	// Game board of is empty by default
	s.board = make([][]*Cell, s.size)

	// Assign all the cells as black hole free.
	for i := 0; i < s.size; i++ {
		s.board[i] = make([]*Cell, s.size)
		for j := 0; j < s.size; j++ {
			s.board[i][j] = &Cell{
				isOpen:             false, // closed by default
				hasBlackHole:       false, // black hole free
				adjacentBlackHoles: 0,     // no adjacent black holes
			}
		}
	}
}

// A function to get all adjacent cells.
// Returns all the 8 adjacent cells.
//     (r-1,c-1) (r-1,c)  (r-1,c+1)
//             \    |    /
//			    \   |   /
//               \  |  /
//     (r,c-1)----(r,c)----(r,c+1)
//                / | \
//               /  |  \
//              /   |   \
//     (r+1,c-1) (r+1,c) (r+1,c+1)
func (s *MinesweeperGame) getSurroundingCells(row int, col int) [][]int {
	surrounding := make([][]int, 0) // collection of surrounding cells
	rows := []int{row - 1, row, row + 1}
	cols := []int{col - 1, col, col + 1}

	for _, nextRow := range rows {
		// Only process a cell if it's valid.
		if nextRow < 0 || nextRow >= s.size {
			continue
		}
		for _, nextCol := range cols {
			// Only process a cell if it's valid.
			if nextCol < 0 || nextCol >= s.size {
				continue
			}
			// Add a valid cell to the collection.
			surrounding = append(surrounding, []int{nextRow, nextCol})
		}
	}

	return surrounding
}

// A function to reveal all the black holes if user has lost.
func (mg *MinesweeperGame) revalAllBlackHoles() {
	// Iterate over all locations of black holes.
	for _, cell := range mg.blackHoleCells {
		mg.board[cell[0]][cell[1]].isOpen = true // open the cell
	}
}
