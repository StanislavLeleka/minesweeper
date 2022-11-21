# minesweeper
Classic puzzle game Minesweeper.

To run the game, open the project folder and run the following command:
```go
go run main.go
```
After that you will be asked to enter the size of the game board and the number of black holes.
```
Enter board size: 8
Enter black holes count: 8
```
After entering the size of the board and the number of black holes, the created board will be printed in this format:
```
  0 1 2 3 4 5 6 7
0 * * * * * * * *
1 * * * * * * * *
2 * * * * * * * *
3 * * * * * * * *
4 * * * * * * * *
5 * * * * * * * *
6 * * * * * * * *
7 * * * * * * * *
```
Then you can enter the coordinates of the cell you want to open (enter two numbers separated by a space).
```
Your move (row, col): 3 5
```
Updated board will be printed after opening the cell.
```
  0 1 2 3 4 5 6 7 
0 0 1 * * * * * * 
1 0 1 * * * * * * 
2 0 1 2 2 2 1 2 1 
3 0 0 0 0 0 0 0 0
4 0 1 1 2 1 1 0 0
5 0 1 * * * 1 0 0
6 0 1 2 * * 2 1 1
7 0 0 1 * * * * *
```
Keep entering coordinates until you win or open a black hole.