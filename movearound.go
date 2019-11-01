package main

import (
	"bufio"
	"fmt"
	"os"
)

func showMap(i [][]int) {
	for a := range i {
		fmt.Println(i[a])
	}
	fmt.Println("")
}

func editMap(x, y int, i [][]int) {
	i[x][y] = 2
}

func move(s string, i [][]int) {
	if s == "w" {
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

		}
	} else if s == "s" {
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

		}
	} else if s == "a" {
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
		}
	} else if s == "d" {
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
		}
	} else if s == "sh" {
		// SHRINK
		fmt.Println("Shrinking the map")
		for in := range i {
			hi := len(i[in]) - 1
			for a := range i[in] {
				wi := len(i[a]) - 1
				i := make([][]int, wi)
				for x := range i {
					i[x] = make([]int, hi)
				}
				editMap(0, 0, i)
				showMap(i)
				return
			}
		}

	} else if s == "gr" {
		// GROW
		fmt.Println("Growing the map")
		i := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
		for in := range i {
			hi := len(i[in]) + 1
			for a := range i[in] {
				wi := len(i[a]) + 1
				i := make([][]int, wi)
				for x := range i {
					i[x] = make([]int, hi)
				}
				editMap(0, 0, i)
				showMap(i)
				return
			}
		}

	}
}

func main() {
	print("Start/n")
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
		if result == "w" {
			move(result, i)
		}
		if result == "s" {
			move(result, i)
		}
		if result == "a" {
			move(result, i)
		}
		if result == "d" {
			move(result, i)
		}
		if result == "gr" {
			move(result, i)
		}
		if result == "sh" {
			move(result, i)
		}
		if result == "q" {
			quitthegame = 1
		}
	}
}
