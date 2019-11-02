package main

import (
	"bufio"
	"fmt"
	"os"
)

func showMap(i [][]int) {
	for a := range i {
		fmt.Println(a, i[a])
	}
	fmt.Println("")
}

func editMap(x, y int, i [][]int) {
	i[x][y] = 2
}

func move(s string, i [][]int) {
	switch s {
	case "w":
		// MOVE UP
		fmt.Println("Moving Up")
		for a := range i {
			if a == 0 {
				for b := range i[a] {
					if i[a+1][b] == 2 {
						i[a+1][b] = 0
						i[a][b] = 2
						showMap(i)
						return
					}
					if i[a][b] == 2 {
						i[a][b] = 0
						i[len(i[a])-1][b] = 2
						showMap(i)
						return
					}
				}
			}
			if a != 0 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a-1][b] = 2
						i[a][b] = 0
						showMap(i)
						return
					}
				}
			}

		} // END MOVE UP
	case "s":
		// MOVE DOWN
		fmt.Println("Moving Down")
		for a := range i {
			if a != len(i[a])-1 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a+1][b] = 2
						showMap(i)
						return
					}
				}
			}
			if a == len(i[a])-1 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[0][b] = 2
						showMap(i)
						return
					}
				}
			}

		} // END MOVE DOWN
	case "a":
		// MOVE LEFT
		fmt.Println("Moving Left")
		for a := range i {
			for b := range i[a] {
				if b == 0 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][len(i[b])-1] = 2
						showMap(i)
						return
					}
				}
				if b != 0 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b-1] = 2
						showMap(i)
						return
					}
				}
			}
		} //END MOVE LEFT
	case "d":
		// MOVE RIGHT
		fmt.Println("Moving Right")
		for a := range i {
			for b := range i[a] {
				if b == len(i[b])-1 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][0] = 2
						showMap(i)
						return
					}
				}
				if b != len(i[b])-1 {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a][b+1] = 2
						showMap(i)
						return
					}
				}
			}
		} //END MOVE RIGHT
	} //END CASES
} // END FUNCTION

//ISSUE IS HERE ======================================================================================================================
func size(s string, i [][]int) [][]int { // This function doesn't work.
	switch s {
	case "gr":
		fmt.Println("Grow")
		//y := len(i) + 1
		//x := len(i[0]) + 1
		for a := range i {
			if a == 1 {
				i = append(i, i[a])
				break
			}
		}
		for a := range i {
			i[a] = append(i[a], 0)
		}
		editMap(0, 0, i)
		showMap(i)
		return i
	case "sh":
		fmt.Println("Shrink")
		for a := range i {
			if a == 1 {
				if len(i[a]) == 1 {
					fmt.Println("Too small")
					return i
				}
			}
		}
		//y := len(i) + 1
		//x := len(i[0]) + 1
		i := i[:len(i)-1]
		for a := range i {
			i[a] = i[a][:len(i[a])-1]
		}
		editMap(0, 0, i)
		showMap(i)
		return i
	}
	editMap(0, 0, i)
	showMap(i)
	return i
}

//ISSUE IS HERE ======================================================================================================================

func main() {
	print("Start\n")
	i := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	editMap(0, 0, i)
	showMap(i)
	scanner := bufio.NewScanner(os.Stdin)
	quitthegame := 0
	fmt.Println("Press w, s, a, d to move the number 2 around the map.")
	fmt.Println("Type sh or gr to shrink / grow the map.")
	fmt.Println("Alternatively, press q to quit.")
	for quitthegame == 0 {
		fmt.Println("Type below: ")
		fmt.Println("(press q to quit)")
		scanner.Scan()
		result := scanner.Text()
		switch result {
		case "w":
			move(result, i)
		case "s":
			move(result, i)
		case "a":
			move(result, i)
		case "d":
			move(result, i)
		case "gr":
			size(result, i)
		case "sh":
			size(result, i)
		case "p":
			showMap(i)
		case "q":
			quitthegame = 1
		}
	}
}
