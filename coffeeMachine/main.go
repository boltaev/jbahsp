package main

import (
	"fmt"
)

// printing remains of products and balance
func remaining(w, m, b, c, moneyHas int) {
	fmt.Println("\nThe coffee machine has:")
	fmt.Println(w, " of water")
	fmt.Println(m, " of milk")
	fmt.Println(b, " of coffee beans")
	fmt.Println(c, " of disposable cups")
	fmt.Println("$", moneyHas, " of money")
	mainAns(w, m, b, c, moneyHas)
}

// asking ml of each product for filling
// then calculating resources
func fill(w, m, b, c, moneyHas int) {
	//scanning info about resources
	var x, y, z, q int
	fmt.Println("\nWrite how many ml of water do you want to add:")
	fmt.Scan(&x)
	w += x

	fmt.Println("Write how many ml of milk do you want to add:")
	fmt.Scan(&y)
	m += y

	fmt.Println("Write how many grams of coffee beans do you want to add:")
	fmt.Scan(&z)
	b += z

	fmt.Println("Write how many disposable cups of coffee do you want to add:")
	fmt.Scan(&q)
	c += q

	mainAns(w, m, b, c, moneyHas)
}

// using if else for checking resources and
// then calculate resources
func enoughRes(w, m, b, c, moneyHas, waterNeed, milkNeed, beansNeed, price int) {
	if w < waterNeed || m < milkNeed || b < beansNeed || c < 1 {
		fmt.Println("Sorry, not enough water!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		w -= waterNeed
		m -= milkNeed
		b -= beansNeed
		moneyHas += price
		c--
	}
	mainAns(w, m, b, c, moneyHas)
}

// clarify type of coffee
// then call enoughRes fuc for checking resources
func buy(w, m, b, c, moneyHas int) {
	var coffeeType string
	var waterNeed, milkNeed, beansNeed, price int

	fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&coffeeType)
	switch coffeeType {
	case "1":
		waterNeed = 250
		beansNeed = 16
		milkNeed = 0
		price = 4
		enoughRes(w, m, b, c, moneyHas, waterNeed, milkNeed, beansNeed, price)

	case "2":
		waterNeed = 350
		milkNeed = 75
		beansNeed = 20
		price = 7
		enoughRes(w, m, b, c, moneyHas, waterNeed, milkNeed, beansNeed, price)

	case "3":
		waterNeed = 200
		milkNeed = 100
		beansNeed = 12
		price = 6
		enoughRes(w, m, b, c, moneyHas, waterNeed, milkNeed, beansNeed, price)

	case "back":
		mainAns(w, m, b, c, moneyHas)

	default:
		fmt.Println("You entered unknown command, is it accurate")
		mainAns(w, m, b, c, moneyHas)
	}

}

// printing total amount of balance
// then reset balance (zeroing)
func take(w, m, b, c, moneyHas int) {
	fmt.Printf("\nI gave you %d $, now you're a rich", moneyHas)
	moneyHas = 0
	mainAns(w, m, b, c, moneyHas)
}

// mainAns is a main menu part function, "all roads lead to the mainAns(rome)"
// asking a type of action, scanning answer, compare with switch and calling this function
func mainAns(waterHas, milkHas, beansHas, cupsHas, moneyHas int) {
	var mainAns string
	fmt.Println("\nWrite action (buy, fill, take, remaining, exit):")
	fmt.Scan(&mainAns)

	switch mainAns {
	case "buy":
		buy(waterHas, milkHas, beansHas, cupsHas, moneyHas)
	case "fill":
		fill(waterHas, milkHas, beansHas, cupsHas, moneyHas)
	case "take":
		take(waterHas, milkHas, beansHas, cupsHas, moneyHas)
	case "remaining":
		remaining(waterHas, milkHas, beansHas, cupsHas, moneyHas)
	case "exit":
		fmt.Println("Bye!")
	default:
		fmt.Println("You entered wrong command(try to restart program)")
	}
}

func main() {
	waterHas := 400
	milkHas := 540
	beansHas := 120
	cupsHas := 9
	moneyHas := 550
	mainAns(waterHas, milkHas, beansHas, cupsHas, moneyHas)
}

//comment part of program: old functions, future features that won't implement.

/*
after choosing the type of coffee:
1 - message about how many cups of this coffee able to make
2 - asking how many cups coffee you need
3 - scanning it
4 - checking one more time and message success or unsuccessful operation and remind about 1 line
*/

/*func ableCups() {
	//calculating how many cups of coffee, machine able to make: N - is amount able cups of coffee
	var cups, N, waterHas, water, milkHas, milk, beansHas, beans int
	N = waterHas / water
	if N < milkHas/milk || N < beansHas/beans {
		if N < milkHas/milk {
			N = milkHas / milk
		}
		if N < beansHas/beans {
			N = beansHas / beans
		}
	}

	//scanning info about demand
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)

	if N == cups {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if N > cups {
		fmt.Println("Yes, I can make that amount of coffee (and even ", cups-N, " more than that)")
	} else {
		fmt.Println("No, I can make only ", N, " cups of coffee")
	}
}
*/
