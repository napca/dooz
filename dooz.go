package main

import "fmt"

func checker(table [9]string) bool {
	winningCombinations := [][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combination := range winningCombinations {
		a, b, c := combination[0], combination[1], combination[2]
		if table[a] != "" && table[a] == table[b] && table[a] == table[c] {
			return true
		}
	}

	return false
}
func draw(table [9]string) {
	for i := 0; i < 7; i++ {
		if i%3 == 0 {
			fmt.Printf("%-1s%s%-1s%s%-1s\n", table[i], "|", table[i+1], "|", table[i+2])
		}
	}

	return
}

func main() {
	table := [9]string{}
	player := false
	var turn string
	var play func()
	play = func() {
		if player == false {
			turn = `X`
		} else {
			turn = `O`
		}
		draw(table)
		fmt.Println(turn, "'s turn, pick from 1-9")
		var choice int
		fmt.Scanln(&choice)
		if choice > 0 && choice < 10 && table[choice-1] == "" {
			table[choice-1] = turn

			if checker(table) == true {
				draw(table)
				fmt.Println(turn, " won")
				return
			} else {
				player = !player
				var t bool
				for i := 0; i < 9; i++ {
					if table[i] == "" {
						break
					}
					if table[i] != "" && i == 8 {
						t = true
					}
					if t {
						fmt.Println("It's a Tie!")
						return
					}
				}
				play()
			}
		} else {
			fmt.Println("wrong choice, retry.")
			play()
		}

	}
	play()

}
