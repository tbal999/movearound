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
			if a != 2 {
				for b := range i[a] {
					if i[a][b] == 2 {
						i[a][b] = 0
						i[a+1][b] = 2
						showMap(i)
						return
					}
				}
			}
			if a == 2 {
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
	}
}

func main() {
	i := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	editMap(1, 1, i)
	showMap(i)
	scanner := bufio.NewScanner(os.Stdin)
	x := 0
	fmt.Println("Press w, s, a, d to move the number 2 around the map")
	fmt.Println("Alternatively, press q to quit.")
	for x == 0 {
		fmt.Println("Type below: ")
		fmt.Println("(press q to quit)")
		scanner.Scan()
		result := scanner.Text()
		if result == "w" {
			move("w", i)
		}
		if result == "s" {
			move("s", i)
		}
		if result == "a" {
			move("a", i)
		}
		if result == "d" {
			move("d", i)
		}
		if result == "q" {
			x = 1
		}
	}
}
