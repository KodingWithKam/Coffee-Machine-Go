package main

import (
	"fmt"
	"strconv"
)

func main() {
	// init selectedOption var
	var selectedOption string
	// setup initial values for machine
	cashAvailable := 550
	waterAvailable := 400
	milkAvailable := 540
	beansAvailable := 120
	cupsAvailable := 9

	// Start process
	for selectedOption != "exit" {
		selectedOption = getUserSelection()
		// handle selection
		switch selectedOption {
		case "buy":
			waterAvailable,
				milkAvailable,
				beansAvailable,
				cupsAvailable,
				cashAvailable = buyCoffeeFromMachine(waterAvailable, milkAvailable, beansAvailable, cupsAvailable, cashAvailable)
		case "fill":
			waterAvailable, milkAvailable, beansAvailable, cupsAvailable = fillMachine(waterAvailable, milkAvailable, beansAvailable, cupsAvailable)
		case "take":
			cashAvailable = withdrawCashFromMachine(cashAvailable)
		case "remaining":
			fmt.Println(getMachineStatus(
				waterAvailable,
				milkAvailable,
				beansAvailable,
				cupsAvailable,
				cashAvailable,
			))
		case "exit":
			break
		default:
			fmt.Println("Entered option is invalid")
		}
	}
}

// function to begin process
func getUserSelection() string {
	// Start process
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	// setup variables
	var selectedOption string
	fmt.Scan(&selectedOption)

	return selectedOption
}

// function to get user coffee order
func getCoffeeOrder() string {
	// get user input for coffee order
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - mocha, back - to main menu:")
	var userSelection string
	fmt.Scan(&userSelection)

	return userSelection
}

// function to buy coffee
func buyCoffeeFromMachine(
	water int,
	milk int,
	beans int,
	cups int,
	cash int) (int, int, int, int, int) {
	// get user order
	var userSelection = getCoffeeOrder()
	if userSelection != "back" {
		// check to see if machine has needed things to process order
		if hasMinimumIngredientsForOrder(userSelection, water, milk, beans, cups) {
			// switch statement
			switch userSelection {
			case "1": // espresso
				// subtract ingredients
				water -= 250
				beans -= 16
				cups -= 1
				// add to order cost
				cash += 4
			case "2": // latte
				// subtract ingredients
				water -= 350
				milk -= 75
				beans -= 20
				cups -= 1
				// add to order cost
				cash += 7
			case "3": // cappuccino
				// subtract ingredients
				water -= 200
				milk -= 100
				beans -= 12
				cups -= 1
				// add to order cost
				cash += 6
			case "4": // mocha
				// subtract ingredients
				water -= 300
				milk -= 250
				beans -= 30
				cups -= 1
				// add to order cost
				cash += 10
			default:
				// invalid
			}
			fmt.Println("I have enough resources, making you a coffee!")
		} else {
			printMissingIngredients(userSelection, water, milk, beans, cups)
		}
	}
	// return updated values
	return water, milk, beans, cups, cash
}

// function to fill coffee machine
func fillMachine(water int, milk int, beans int, cups int) (int, int, int, int) {
	// water
	fmt.Println("Write how many ml of water you want to add:")
	var coffeeMachineWater int
	fmt.Scan(&coffeeMachineWater)

	// milk
	fmt.Println("Write how many ml of milk you want to add:")
	var coffeeMachineMilk int
	fmt.Scan(&coffeeMachineMilk)

	// beans
	fmt.Println("Write how many grams of coffee beans you want to add:")
	var coffeeMachineBeans int
	fmt.Scan(&coffeeMachineBeans)

	// cups
	fmt.Println("Write how many disposable cups you want to add:")
	var coffeeMachineCups int
	fmt.Scan(&coffeeMachineCups)

	// add current state with new items filled
	water += coffeeMachineWater
	milk += coffeeMachineMilk
	beans += coffeeMachineBeans
	cups += coffeeMachineCups

	return water, milk, beans, cups
}

// function to withdraw money from machine
func withdrawCashFromMachine(cash int) int {
	fmt.Println("I gave you $" + strconv.Itoa(cash))
	fmt.Println()
	withdrawAmount := cash * 0
	return withdrawAmount
}

// function to print status of machine
func getMachineStatus(water int, milk int, beans int, cups int, cash int) string {
	return "The coffee machine has:\n" +
		strconv.Itoa(water) + " ml of water\n" +
		strconv.Itoa(milk) + " ml of milk\n" +
		strconv.Itoa(beans) + " g of coffee beans\n" +
		strconv.Itoa(cups) + " disposable cups \n" +
		"$" + strconv.Itoa(cash) + " of money \n"
}

// function to print errors for lack of ingredients
func printMissingIngredients(orderOption string, water int, milk int, beans int, cups int) {
	switch orderOption {
	case "1":
		// espresso
		if water < 250 {
			fmt.Println("Sorry, not enough water!")
		}
		if beans < 16 {
			fmt.Println("Sorry, not enough beans!")
		}
	case "2":
		// latte
		if water < 350 {
			fmt.Println("Sorry, not enough water!")
		}
		if milk < 75 {
			fmt.Println("Sorry, not enough milk!")
		}
		if beans < 20 {
			fmt.Println("Sorry, not enough beans!")
		}
	case "3":
		// cappuccino
		if water < 200 {
			fmt.Println("Sorry, not enough water!")
		}
		if milk < 100 {
			fmt.Println("Sorry, not enough milk!")
		}
		if beans < 20 {
			fmt.Println("Sorry, not enough beans!")
		}
	default:
	}
	if cups < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}

// function to check if min ingredients are available
func hasMinimumIngredientsForOrder(orderOption string, water int, milk int, beans int, cups int) bool {
	switch orderOption {
	case "1":
		// espresso
		return water >= 250 && beans >= 16 && cups >= 1
	case "2":
		// latte
		return water >= 350 && milk >= 75 && beans >= 20 && cups >= 1
	case "3":
		// cappuccino
		return water >= 200 && milk >= 100 && beans >= 20 && cups >= 1
	case "4":
		// mocha
		return water >= 300 && milk >= 250 && beans >= 30 && cups >= 1
	default:
		return false
	}
}
