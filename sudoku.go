package main

import (
	"fmt"
	"os"
)

// sudoku tahtasının boyutu
const N = 9

func parseBoard(args []string) [][]int {
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
		for j, char := range args[i] {
			if char == '.' {
				board[i][j] = 0
			} else {
				board[i][j] = int(char - '0')
			}
		}
	}
	return board
}

func printBoard(board [][]int) {
	for i := 0; i < N; i++ { // satırları ve her satırdaki sürunları dolaşır
		for j := 0; j < N; j++ {
			fmt.Print(board[i][j], " ") // hücreler arası boşluk bırakır
		}
		fmt.Println() // satır sonu yeni satıra geçer
	}
}

func solveSudoku(board [][]int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= N; num++ { // 1-9 arası sayıları dener
					if isSafe(board, row, col, num) {
						board[row][col] = num   // sayı doğruysa sayıyı yerleştir
						if solveSudoku(board) { //backtracking yaparak tahtanın kalan kısmı çözülmeye çalışılır
							return true // çözüm bulunursa true döner
						}
						board[row][col] = 0 // çözüm bulunmazsa geri al
					}
				}
				return false // denenen hiçbir sayı uymazsa false döner
			}
		}
	}
	return true // tüm hücreler doluysa çözüm bulunmuştur
}

func isSafe(board [][]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if board[row][x] == num { // aynı satırdaki tüm hücreler dolaşılır
			return false // num zaten bu satırda varsa false döner
		}
	}

	for y := 0; y < N; y++ {
		if board[y][col] == num { // aynı sütundaki tüm hücreler dolaşılır
			return false // num zaten bu sütunda varsa false döner
		}
	}

	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true // bu üç kontrolden hiçbiri false döndürmediyse, bu hücreye num yerleştirilir
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	board := parseBoard(os.Args[1:])
	if solveSudoku(board) {
		printBoard(board) // çözülen tahtayı yazdırmak için
	} else {
		fmt.Println("\nError") // çözüm bulunamazsa
	}
}
