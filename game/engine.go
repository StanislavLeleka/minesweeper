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
	// Get the size of the game board and the number of black holes.
	size, blackHoles := getGameSettings()
	// Initialize the game.
	game := newGame(size, blackHoles)
	// Place the black holes randomly on the game board.
	game.populateBlackHoles()

	// User stays in the game until a black hole is found or the user wins.
	// So we keep waiting for the game status to change.
	for game.gameStatus == Playing {
		fmt.Println("Current board:")
		// Show the user the current game board.
		game.printBoard(false)

		// Get the coordinates of the cell to open from the user.
		row, col := makeMove()
		// If the coordinates are valid, open the cell.
		// If the coordinates are invalid, go back to entering new coordinates.
		if positionIsValid(row, col, size) {
			game.openCell(row, col)
		}
	}

	// If the game loop is over, it means the game is over.
	// Show the current game board to the user and reveal all the black holes.
	game.printBoard(true)
	// Notify the user whether the game is won or lost.
	if game.gameStatus == Lost {
		fmt.Println("You lost!")
	} else if game.gameStatus == Won {
		fmt.Println("You win!")
	}
}

// A function to get board size and black hole count from user.
// Returns board size and number of black holes.
func getGameSettings() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	getUserInput := func(msg string, isValid func(int) bool) int {
		for {
			fmt.Print(msg)
			// Read user input.
			input, err := reader.ReadString('\n')
			// If there is an error, start over.
			if err != nil {
				continue
			}
			val, err := strconv.Atoi(strings.TrimSpace(input)) // convert user input to int
			// If there is an error, start over.
			if err != nil || !isValid(val) {
				continue
			}
			return val // return valid input
		}
	}
	// Get board size.
	size := getUserInput("Enter board size: ", isBoardSizeValid)
	// Get the number of black holes.
	blackHoles := getUserInput("Enter black holes count: ",
		func(bc int) bool { return isBlackHolesCountValid(bc, size) })

	return size, blackHoles
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

// A function to validate game board size.
func isBoardSizeValid(size int) bool {
	return size > 0
}

// A function to validate number of black holes.
func isBlackHolesCountValid(blackHoles int, size int) bool {
	return blackHoles > 1 && blackHoles < size*size
}

// A function of validate the entered coordinates by the user.
func positionIsValid(x int, y int, max int) bool {
	return (x >= 0 && x < max) && (y >= 0 && y < max)
}
