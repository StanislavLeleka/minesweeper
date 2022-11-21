package game

import "fmt"

// A function to play the game.
func Play() {
	// Get board size and black hole count from user.
	size, blackHoles := getGameSettings()
	// Start the game and get a game object
	// or an error if the user input is not valid.
	game, err := startGame(size, blackHoles)
	// If we have an error, go back to entering the size of the board
	// and the number of black holes.
	if err != nil {
		fmt.Println("Failed to start game. Error:", err.Error())
		fmt.Println("Try again.")
		// Call Play() again to enter board size
		// and black hole count again.
		Play()
	}

	// User stays in the game until a black hole is found
	// or the user wins.
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
	// Show the current game board to the user
	// and reveal all the black holes.
	game.printBoard(true)
	// Notify the user whether the game is won or lost.
	if game.GameStatus == Lost {
		fmt.Println("You lost!")
	} else if game.GameStatus == Won {
		fmt.Println("You win!")
	}
}

// A function to get board size and black hole count from user.
// Returns the entered board size and number of black holes.
func getGameSettings() (int, int) {
	var size int       // holds board size
	var blackHoles int // holds number of black holes

	fmt.Print("Enter board size: ")
	// Read user input into the size variable.
	_, err := fmt.Scanln(&size)
	// If input is invalid (not a number) panic is called.
	if err != nil {
		panic("Invalid input!")
	}

	fmt.Print("Enter black holes count: ")
	// Read user input into the blackHoles variable.
	_, err = fmt.Scanln(&blackHoles)
	// If input is invalid (not a number) panic is called.
	if err != nil {
		panic("Invalid input!")
	}

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
	_, err := fmt.Scanf("%d %d \n", &row, &col)
	// If the input is invalid, print an error and return incorrect coordinates.
	if err != nil {
		fmt.Println(err)
		fmt.Println("Invalid input! Try again.")
		return -1, -1
	}

	return row, col
}

// A function to start the game.
// Creates a new instance of the game and returns it,
// or returns an error if there was one when the game was initialized.
func startGame(size int, blackHoles int) (*MinesweeperGame, error) {
	// Initialize the game.
	game, err := newGame(size, blackHoles)
	// If there is an error, return it.
	if err != nil {
		return nil, err
	}
	// Place the black holes randomly on the game board.
	game.populateBlackHoles()

	// Return the created game instance.
	return game, nil
}
