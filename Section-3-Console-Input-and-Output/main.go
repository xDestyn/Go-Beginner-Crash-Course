package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	// Check to see if we cannot open the keyboard package.
	if err != nil {
		log.Fatal(err)
	}

	// This will execute when the main function completes.
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	char := ' '

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for char != 'q' && char != 'Q' {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))

		if _, ok := coffees[i]; ok {
			fmt.Println(fmt.Printf("You chose %s", coffees[i]))
		}

		// if key != 0 {
		// 	fmt.Println("You pressed: ", char, key)
		// } else {
		// 	fmt.Println("You pressed: ", char)
		// }

		// if key == keyboard.KeyEsc {
		// 	break
		// }
	}

	fmt.Println("Program exiting...")
}
