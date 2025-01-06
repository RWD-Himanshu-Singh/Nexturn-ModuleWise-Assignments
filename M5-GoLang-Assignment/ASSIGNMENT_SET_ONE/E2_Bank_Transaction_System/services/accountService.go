package services

import (
	"errors"
	"ex2_Bank/models"
	"fmt"
)

var accounts []models.Account

// AddAccount adds a new account
func AddAccount(name string, initialDeposit float64) models.Account {
	id := len(accounts) + 1
	account := models.Account{
		ID:      id,
		Name:    name,
		Balance: initialDeposit,
	}
	accounts = append(accounts, account)
	return account
}

// FindAccount searches for an account by ID
func FindAccount(accountID int) (*models.Account, error) {
	for i := range accounts {
		if accounts[i].ID == accountID {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// Deposit adds funds to an account
func Deposit(accountID int, amount float64) error {
	account, err := FindAccount(accountID)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	account.Balance += amount
	account.TransactionLog = append(account.TransactionLog, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw deducts funds from an account
func Withdraw(accountID int, amount float64) error {
	account, err := FindAccount(accountID)
	if err != nil {
		return err
	}

	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	if amount > account.Balance {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	account.TransactionLog = append(account.TransactionLog, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// ViewBalance shows the balance of an account
func ViewBalance(accountID int) (float64, error) {
	account, err := FindAccount(accountID)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

// ViewTransactionHistory shows the transaction history of an account
func ViewTransactionHistory(accountID int) ([]string, error) {
	account, err := FindAccount(accountID)
	if err != nil {
		return nil, err
	}
	return account.TransactionLog, nil
}
