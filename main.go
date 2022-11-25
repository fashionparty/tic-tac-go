package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var arrayOfFields [3][3]Field
var isFieldBlank bool
var isFieldCorrect bool
var continueGame = true

// main Gameplay scenario code
func main() {
	fmt.Println("Welcome to Tic-Tac-Go.\nYou play as Cross and first move is yours.")
	fmt.Println("How to play: place cross on the board by passing x/y coordinates of the field -> two numbers (1-3) without spaces (e.g. `11` `13 `32` `22`)")
	for continueGame {
		fmt.Println()
		for i := 0; true; i++ {
			if checkResult() == None {
				drawBoard()
				askForCoordinates()
				if isFieldBlank && isFieldCorrect {
					computerMove()
				}
				fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")
				fmt.Println()
				if !isFieldBlank {
					fmt.Println("Given field is not blank! Please select different coordinates.")
					fmt.Println()
				}
				if !isFieldCorrect {
					fmt.Println("Given coordinates are incorrect! Please try again.")
					fmt.Println()
				}
			} else {
				break
			}
		}
		drawBoard()
		fmt.Println("Game Over")
		var gameWinner string
		if checkResult() == Player {
			gameWinner = "Player"
		} else {
			gameWinner = "Computer"
		}
		fmt.Println(gameWinner + " won the game!")
		fmt.Print("\nStart again? y/n:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "y" {
			continueGame = true
			clearBoard()
		} else if input == "n" {
			continueGame = false
			fmt.Println()
			fmt.Println("Thanks for playing!")
		} else {
			fmt.Println("Wrong input.")
		}
	}
}

func clearBoard() {
	for i := 0; i < len(arrayOfFields); i++ {
		for k := 0; k < len(arrayOfFields); k++ {
			arrayOfFields[i][k].Symbol = Empty
		}
	}
}

// checkResult Checks if someone won
func checkResult() winner {
	gameWinner := winner(None)
	if arrayOfFields[0][0].Symbol == Cross && arrayOfFields[0][1].Symbol == Cross && arrayOfFields[0][2].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[1][0].Symbol == Cross && arrayOfFields[1][1].Symbol == Cross && arrayOfFields[1][2].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[2][0].Symbol == Cross && arrayOfFields[2][1].Symbol == Cross && arrayOfFields[2][2].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][0].Symbol == Cross && arrayOfFields[1][0].Symbol == Cross && arrayOfFields[2][0].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][1].Symbol == Cross && arrayOfFields[1][1].Symbol == Cross && arrayOfFields[2][1].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][2].Symbol == Cross && arrayOfFields[1][2].Symbol == Cross && arrayOfFields[2][2].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][0].Symbol == Cross && arrayOfFields[1][1].Symbol == Cross && arrayOfFields[2][2].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][2].Symbol == Cross && arrayOfFields[1][1].Symbol == Cross && arrayOfFields[2][0].Symbol == Cross {
		gameWinner = Player
	} else if arrayOfFields[0][0].Symbol == Circle && arrayOfFields[0][1].Symbol == Circle && arrayOfFields[0][2].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[1][0].Symbol == Circle && arrayOfFields[1][1].Symbol == Circle && arrayOfFields[1][2].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[2][0].Symbol == Circle && arrayOfFields[2][1].Symbol == Circle && arrayOfFields[2][2].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[0][0].Symbol == Circle && arrayOfFields[1][0].Symbol == Circle && arrayOfFields[2][0].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[0][1].Symbol == Circle && arrayOfFields[1][1].Symbol == Circle && arrayOfFields[2][1].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[0][2].Symbol == Circle && arrayOfFields[1][2].Symbol == Circle && arrayOfFields[2][2].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[0][0].Symbol == Circle && arrayOfFields[1][1].Symbol == Circle && arrayOfFields[2][2].Symbol == Circle {
		gameWinner = Computer
	} else if arrayOfFields[0][2].Symbol == Circle && arrayOfFields[2][2].Symbol == Circle && arrayOfFields[2][0].Symbol == Circle {
		gameWinner = Computer
	}
	return gameWinner
}

// computerMove Checks which of the fields are still blank and places circle randomly on one of them
func computerMove() {
	var blankCoordinates []string
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if arrayOfFields[k][i].Symbol == Empty {
				firstCoordinate := strconv.Itoa(k + 1)
				secondCoordinate := strconv.Itoa(i + 1)
				blankCoordinates = append(blankCoordinates, firstCoordinate+secondCoordinate)
			}
		}
	}
	if len(blankCoordinates) > 0 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		circleIndex := r1.Intn(len(blankCoordinates))
		selectedBlankCoordinate := blankCoordinates[circleIndex]
		firstSelectdValue, _ := strconv.ParseInt(selectedBlankCoordinate[:1], 32, 64)
		secondSelectdValue, _ := strconv.ParseInt(selectedBlankCoordinate[1:2], 32, 64)
		arrayOfFields[firstSelectdValue-1][secondSelectdValue-1].Symbol = Circle
	}
}

// askForCoordinates Reads user input and calls validation functions
func askForCoordinates() {
	isFieldCorrect = true
	isFieldBlank = true
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Enter coordiantes: ")
	input, _ := reader.ReadString('\n')
	fmt.Println()
	if validateUserInput(input[:2]) {
		firstValueString := string(input[0])
		firstValueInt, _ := strconv.ParseInt(firstValueString, 32, 64)
		secondValueString := string(input[1])
		secondValueInt, _ := strconv.ParseInt(secondValueString, 32, 64)
		if validateIfGivenFieldIsBlank(arrayOfFields[firstValueInt-1][secondValueInt-1]) {
			isFieldBlank = true
			arrayOfFields[firstValueInt-1][secondValueInt-1].Symbol = Cross
		} else {
			isFieldBlank = false
		}
	} else {
		isFieldCorrect = false
	}
}

// validateUserInput Checks if user is selecting correct coordinates
func validateUserInput(input string) bool {
	allowedInput := [9]string{"11", "12", "13", "21", "22", "23", "31", "32", "33"}
	isValid := true
	for _, value := range allowedInput {
		if value != input {
			isValid = false
		} else {
			isValid = true
			break
		}
	}
	return isValid
}

// validateIfGivenFieldIsBlank Returns true if given field is empty
func validateIfGivenFieldIsBlank(f Field) bool {
	if f.Symbol == Empty {
		return true
	} else {
		return false
	}
}

// drawBoard Draws board and one of three symbols in each field
func drawBoard() {
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			var drawedSymbol string
			if arrayOfFields[k][i].Symbol == Empty {
				drawedSymbol = "*"
			} else if arrayOfFields[k][i].Symbol == Cross {
				drawedSymbol = "X"
			} else {
				drawedSymbol = "O"
			}
			fmt.Print("[" + drawedSymbol + "]")
		}
		fmt.Println()
	}
}
