package main

import (
	"ex2_Bank/services"
	"ex2_Bank/utils"
	"fmt"
	"os"
)

func displayMenu() {
	fmt.Println("\nBank Transaction System")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. View Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func main() {
	services.AddAccount("Rakesh", 2000)
	services.AddAccount("Ganesh", 600)

	var choice int
	for {
		displayMenu()
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		var accountID int
		var amount float64

		switch choice {
		case utils.DepositOption:
			fmt.Print("Enter Account ID: ")
			_, err := fmt.Scan(&accountID)
			if err != nil {
				fmt.Println("Invalid account ID")
				continue
			}
			fmt.Print("Enter deposit amount: ")
			_, err = fmt.Scan(&amount)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid deposit amount")
				continue
			}
			err = services.Deposit(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case utils.WithdrawOption:
			fmt.Print("Enter Account ID: ")
			_, err := fmt.Scan(&accountID)
			if err != nil {
				fmt.Println("Invalid account ID")
				continue
			}
			fmt.Print("Enter withdrawal amount: ")
			_, err = fmt.Scan(&amount)
			if err != nil || amount <= 0 {
				fmt.Println("Invalid withdrawal amount")
				continue
			}

			err = services.Withdraw(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case utils.ViewBalanceOption:
			fmt.Print("Enter Account ID: ")
			_, err := fmt.Scan(&accountID)
			if err != nil {
				fmt.Println("Invalid account ID")
				continue
			}

			balance, err := services.ViewBalance(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account balance: %.2f\n", balance)
			}

		case utils.ViewHistoryOption:
			fmt.Print("Enter Account ID: ")
			_, err := fmt.Scan(&accountID)
			if err != nil {
				fmt.Println("Invalid account ID")
				continue
			}

			history, err := services.ViewTransactionHistory(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, transaction := range history {
					fmt.Println(transaction)
				}
			}

		case utils.ExitOption:
			fmt.Println("Exiting program...")
			os.Exit(0)

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
