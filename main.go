package main

import (
	"fmt"
	"math/rand"
	"time"
)

func checkingInRow(rowValues []int, num int) bool {
	isDuplicate := false
	for _, eachVal := range rowValues {
		if num == eachVal {
			isDuplicate = true
			break
		}
	}
	return isDuplicate
}
func checkingInColumn(allRows [][]int, rowValues []int, num int) bool {
	isDuplicate := false
	//fmt.Println(allRows) //to test
	//fmt.Println(rowValues) //to test
	numIndex := len(rowValues)
	for _, echSlice := range allRows {
		if num == echSlice[numIndex] {
			isDuplicate = true
			//fmt.Printf("%d is not valid in equal index.\n", num) //to test
			break
		}
	}
	//if isDuplicate == false { //to test
	//	//fmt.Printf("%d is valid in equal index.\n", num)
	//}
	return isDuplicate
}
func checkingInSquare(allRows [][]int, rowValues []int, num int) bool {
	isDuplicate := false
	//fmt.Println(allRows, num) //to test
	y := len(allRows) + 1
	x := len(rowValues) + 1
	startOfY := y - ((y - 3*((y-1)/3)) - 1)
	startOfX := x - ((x - 3*((x-1)/3)) - 1)
	for j := startOfY; j < startOfY+((y-3*((y-1)/3))-1); j++ {
		for i := startOfX; i < startOfX+3; i++ {
			//fmt.Println("y is:", y, "x is:", x)       //(to test)
			//fmt.Println(num, allRows[j-1][i-1], j, i) //(to test)
			if num == allRows[j-1][i-1] {
				isDuplicate = true
				//fmt.Printf("%d is not valid in square.\n", num) //to test
				break
			}
		}
		if isDuplicate == true {
			break
		}
	}
	//if isDuplicate == false { //to test
	//	fmt.Printf("%d is valid in square.\n", num)
	//}
	return isDuplicate
}

func randomList(input int) [][]int {
	rand.Seed(time.Now().Unix())
	var allRows [][]int
	for len(allRows) < input {
		var rowValues []int
		for len(rowValues) < input {
			//make an empty list
			var randList []int
			//Create a loop to iterate until the list is input long
			var check int
			for len(randList) < input {
				//Take the random number 1 to Input
				num := rand.Intn(input) + 1
				//Using the checkingInRow function to check Num is not repeated
				if checkingInRow(randList, num) == false {
					//If num is not repeated, put it in the list
					randList = append(randList, num)
				} else {
					continue
				}
				//Checking for non-repetition of num in the column
				if checkingInColumn(allRows, rowValues, num) == false {
					if checkingInSquare(allRows, rowValues, num) == false {
						rowValues = append(rowValues, num)
					} else {
						//num is not valid,remove it from the list
						randList = randList[:len(randList)-1]
						//Checking the previous num and the new num
						if check == num {
							//Empty the not valid lists
							rowValues = rowValues[:0]
							randList = randList[:0]
							break
						}
						//add previous Num in variable
						check = num
					}
				} else {
					//num is not valid,remove it from the list
					randList = randList[:len(randList)-1]
					//Checking the previous num and the new num
					if check == num {
						//Empty the not valid lists
						rowValues = rowValues[:0]
						randList = randList[:0]
						break
					}
					//add previous Num in variable
					check = num
				}
			}
		}
		allRows = append(allRows, rowValues)
	}
	fmt.Println(allRows)
	return allRows
}

func makeTable(input int, tableValues [][]int) {
	fmt.Print("┌")
	for i := 0; i < input-1; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("-")
		}
		fmt.Print("┬")
	}
	for j := 0; j < 3; j++ {
		fmt.Print("-")
	}
	fmt.Print("┐\n")
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print("|")
			fmt.Print(" ")
			fmt.Print(tableValues[i][j])
			fmt.Print(" ")
		}
		fmt.Print("|\n")
		fmt.Print("└")
		for j := 0; j < (input)-1; j++ {
			for k := 0; k < 3; k++ {
				fmt.Print("-")
			}
			fmt.Print("┴")
		}
		for k := 0; k < 3; k++ {
			fmt.Print("-")
		}
		fmt.Print("┘")
		fmt.Println()
	}
}
func main() {
	fmt.Println("Let's make a sudoku table!!")
	fmt.Println("Enter the number of rows and columns:")
	//The program is coded for number 9.
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	makeTable(input, randomList(input))
}

//Note

//برای y = 3 و x = 3
//x = 2 , y = 3 لازم نیس
//x = 1 , y = 3 لازم نیس
//x = 3 , y = 2 چک میکنه
//x = 2 , y = 2 چک میکنه
//x = 1 , y = 2 چک میکنه
//x = 3 , y = 1 چک میکنه
//x = 2 , y = 1 چک میکنه
//x = 1 , y = 1 چک میکنه

//برای y = 9 و x = 9
//x = 8 , y = 9 لازم نیس
//x = 7 , y = 9 لازم نیس
//x = 9 , y = 8 چک میکنه
//x = 8 , y = 8 چک میکنه
//x = 7 , y = 8 چک میکنه
//x = 9 , y = 7 چک میکنه
//x = 8 , y = 7 چک میکنه
//x = 7 , y = 7 چک میکنه

//برای y = 2 و x = 8
//x = 7 , y = 2 لازم نیس
//x = 9 , y = 1 چک میکنه
//x = 8 , y = 1 چک میکنه
//x = 7 , y = 1 چک میکنه

//برای y = 4 و x = 5
//x = 4 , y = 4 لازم نیس

//برای y = 8 و x = 3
//x = 2 , y = 8 لازم نیس
//x = 1 , y = 8 لازم نیس
//x = 3 , y = 7 چک میکنه
//x = 2 , y = 7 چک میکنه
//x = 1 , y = 7 چک میکنه

//1 - 3 * 0
//2 - 3 * 0
//3 - 3 * 0
//4 - 3 * 1
//5 - 3 * 1
//6 - 3 * 1
//7 - 3 * 2
//8 - 3 * 2
//9 - 3 * 2
//y - 3 *(y-1)/3
// len(y) = y - 3 * (y-1)/3

//1 + 2 -> x +(1*3 - x)
//2 + 1 -> x +(1*3 - x)
//3 + 0 -> x +(1*3 - x)
//4 + 2 -> x +(2*3 - x)
//5 + 1 -> x +(2*3 - x)
//6 + 0 -> x +(2*3 - x)
//7 + 2 -> x +(3*3 - x)
//8 + 1 -> x +(3*3 - x)
//9 + 0 -> x +(3*3 - x)
//x+((x+2)/3)*3 - x)

//1 - (1-1)0 = 1
//2 - (2-1)1 = 1
//3 - (3-1)2 = 1
//4 - ((x - 3 * (x-1)/3)-1)0 = 4
//5 - ((x - 3 * (x-1)/3)-1)1 = 4
//6 - 2 = 4
//7 - 0 = 7
//8 - 1 = 7
//9 - 2 = 7
