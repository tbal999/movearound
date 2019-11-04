package main

import (
	"bufio"
	"fmt"
	"os"
)

func showMap(i [][]int) { //Prints out the 2D slice
	for a := range i {
		fmt.Println(a, i[a])
	}
	fmt.Println("")
}

func editMap(i [][]int) [][]int { //Edits the digit at i[x][y] so it's a 2
	for index := range i {
		for index1 := range i[index] {
			if i[index][index1] == 2 {
				i[index][index1] = 0
			}
		}
	}
	i[0][0] = 2
	return i
}

func move(s string, i [][]int) { //Moves the number 2 in the slice around, up,down,left,right
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

func size(a string, i *[][]int) {
	m := *i
	switch a {
	case "gr":
		newline := make([]int, len(m[0]))
		m = append(m, newline)
		for i := range m {
			m[i] = append(m[i], 0)
		}
	case "sh":
		if len(m) == 1 {
			fmt.Println("Too Small!")
			return
		}
		m = m[:len(m)-1]
		for i := range m {
			m[i] = m[i][:len(m[i])-1]
		}
	default:
		return
	}
	*i = m
	editMap(m)
}

func main() {
	print("Start\n")
	i := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	editMap(i)
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
			size(result, &i)
		case "sh":
			size(result, &i)
		case "p":
			showMap(i)
		case "q":
			quitthegame = 1
		}
	}
}
