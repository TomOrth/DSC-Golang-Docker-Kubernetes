package main

import "fmt"

type Account struct {
	Name   string
	Amount float64
}

func (a *Account) Add(amt float64) {
	a.Amount = a.Amount + amt
}

func (a *Account) Withdraw(amt float64) bool {
	if a.Amount < amt {
		return false
	}
	a.Amount = a.Amount - amt
	return true
}

func (a Account) Print() {
	fmt.Printf("Account Name: %s, Amount: %f\n", a.Name, a.Amount)
}

func main() {
	fmt.Println("Starting balance for your checking account: ")
	var startCheck float64
	fmt.Scanf("%f", &startCheck)
	checking := &Account{
		"Checking",
		startCheck,
	}

	fmt.Println("Starting balance for your savings account: ")
	var startSaving float64
	fmt.Scanf("%f", &startSaving)
	savings := &Account{
		"Saving",
		startSaving,
	}
	run := true
	for run {
		fmt.Println("Enter option:")
		fmt.Println("1: Add amount to account")
		fmt.Println("2: Withdraw an amount from account")
		fmt.Println("3: Print account information")
		fmt.Println("4: Exit")
		var choice int
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("How much?")
			var amt float64
			fmt.Scanf("%f", &amt)
			fmt.Println("Checking or Savings?")
			var name string
			fmt.Scanf("%s", &name)
			if name == "Checking" {
				checking.Add(amt)
			} else if name == "Savings" {
				savings.Add(amt)
			} else {
				fmt.Println("Invalid option choosen! Please choose Checking or Savings")
			}
		case 2:
			fmt.Println("How much?")
			var amt float64
			fmt.Scanf("%f", &amt)
			fmt.Println("Checking or Savings?")
			var name string
			fmt.Scanf("%s", &name)
			if name == "Checking" {
				res := checking.Withdraw(amt)
				if res {
					fmt.Println("Withdraw successful!")
				} else {
					fmt.Println("Insufficient funds")
				}
			} else if name == "Savings" {
				res := savings.Withdraw(amt)
				if res {
					fmt.Println("Withdraw successful!")
				} else {
					fmt.Println("Insufficient funds")
				}
			} else {
				fmt.Println("Invalid option choosen! Please choose Checking or Savings")
			}
		case 3:
			fmt.Println("Checking or Savings?")
			var name string
			fmt.Scanf("%s", &name)
			if name == "Checking" {
				checking.Print()
			} else if name == "Savings" {
				savings.Print()
			} else {
				fmt.Println("Invalid option choosen! Please choose Checking or Savings")
			}
		case 4:
			run = false
		default:
			fmt.Println("Invalud option")
		}
	}

}
