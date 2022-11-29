package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A function to play the game.
func Play() {
	var size int       // holds board size
	var blackHoles int // holds number of black holes
	// Get board size and black hole count from user.
	size, blackHoles, err := getGameSettings()
	// Validate user input.
	// If the entered values are invalid or there are errors during the input, start the input again.
	for err != nil || !validateInput(size, blackHoles) {
		fmt.Println("Game settings are invalid. Please try again.")
		size, blackHoles, err = getGameSettings()
	}
	// Initialize the game.
	game := newGame(size, blackHoles)
	// Place the black holes randomly on the game board.
	game.populateBlackHoles()

	// User stays in the game until a black hole is found or the user wins.
	// So we keep waiting for the game status to change.
	for game.GameStatus == Playing {
		fmt.Println("Current board:")
		// Show the user the current game board.
		game.printBoard(false)

		// Get the coordinates of the cell to open from the user.
		row, col := makeMove()
		// If the coordinates are valid, open the cell.
		// If the coordinates are invalid, go back to entering new coordinates.
		if (row >= 0 && row < size) && (col >= 0 && col < size) {
			game.openCell(row, col)
		} else {
			continue
		}
	}

	// If the game loop is over, it means the game is over.
	// Show the current game board to the user and reveal all the black holes.
	game.printBoard(true)
	// Notify the user whether the game is won or lost.
	if game.GameStatus == Lost {
		fmt.Println("You lost!")
	} else if game.GameStatus == Won {
		fmt.Println("You win!")
	}
}

// A function to get board size and black hole count from user.
// Returns the entered board size and number of black holes or error if input is invalid.
func getGameSettings() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter board size: ")
	// Read user input.
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	size, err := strconv.Atoi(strings.TrimSpace(input)) // get board size
	// If input is invalid return error.
	if err != nil {
		return 0, 0, err
	}

	fmt.Print("Enter black holes count: ")
	// Read user input.
	input, err = reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	blackHoles, err := strconv.Atoi(strings.TrimSpace(input)) // number of black holes
	// If input is invalid return error.
	if err != nil {
		return 0, 0, err
	}

	return size, blackHoles, nil
}

// A function to get the coordinates of the cell to open from the user.
// Returns the entered coordinates (row and column).
func makeMove() (int, int) {
	fmt.Print("Your move (row, col): ")

	var row int // holds row value
	var col int // holds column value
	// Read user input into the row and col variables.
	// The user must enter two numbers separated by a space.
	_, err := fmt.Scanf("%d %d\n", &row, &col)
	// If the input is invalid, print an error and return incorrect coordinates.
	if err != nil {
		fmt.Println(err)
		fmt.Println("Invalid input! Try again.")
		return -1, -1
	}

	return row, col
}

// A function to validate board size and black holes count
// Returns false if the input parameters are invalid.
func validateInput(size int, blackHoles int) bool {
	// Board size validation.
	if size < 1 {
		return false
	}

	// Number of black holes validation.
	if blackHoles < 1 || blackHoles > size*size {
		return false
	}

	return true
}
