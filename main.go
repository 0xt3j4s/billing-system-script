package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string,r *bufio.Reader) string {
	fmt.Println(prompt)
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)
	
	option := getInput("Choose from the options given below:\n	 A \t : add item,\n	 T \t : add tip,\n	 S \t : show bill,\n	 UP \t : update price of the item,\n	 UT \t : update the tip of the bill,\n	 D \t : delete the item,\n	 E \t : exit \n>", reader)


	switch option {
	case "A":
		fmt.Println("Add an item")
		name := getInput("Item name: ", reader)
		price := getInput("Item price: ", reader)
		price = strings.Trim(price, "$")

		pre, _ := strconv.ParseFloat(price, 64)

		b.addItem(name, pre)
		fmt.Println("Item added - ", name, price)
	case "T":
		fmt.Println("Add a tip")
		tip := getInput("Enter tip amount ($): ", reader)
		tip = strings.Trim(tip, "$")

		pre, _ := strconv.ParseFloat(tip, 64)

		b.addTip(pre)
	case "S":
		fmt.Println("Show bill")
		fmt.Println(b.format())
	case "UP":
		fmt.Println("Update price")
	case "UT":
		fmt.Println("Update tip")
	case "D":
		fmt.Println("Delete item")
	case "E":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid option, please try again...")
		promptOptions(b)
	}
	promptOptions(b)
}


func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name := getInput("Create a new bill name: ", reader)
	fmt.Println("Created the bill - ", name)

	myBill := newBill(name)

	return myBill
}

func main() {
	fmt.Println("Billing System")
	
	newBill := createBill()

	// newBill.addItem("onion rings", 5.55)
	// newBill.addItem("mushroom rings", 6.55)
	// newBill.addItem("fries", 4.55)
	// newBill.addItem("cheese fries", 5.55)

	// newBill.addTip(5.00)
	fmt.Println(newBill.format())

	promptOptions(&newBill)
	
}

