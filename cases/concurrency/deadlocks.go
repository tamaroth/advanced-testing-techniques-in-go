package concurrency

import (
	"time"

	"github.com/sasha-s/go-deadlock"
)

type BankAccount struct {
	balance int
	mu      deadlock.Mutex
}

func (acc *BankAccount) Transfer(to *BankAccount, amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	time.Sleep(time.Millisecond) // Force deadlock.
	to.mu.Lock()
	defer to.mu.Unlock()
	acc.balance -= amount
	to.balance += amount
}

// mtx     deadlock.Mutex

// Deposit adds money to the account
func (acc *BankAccount) Deposit(amount int) {
	acc.balance += amount // Race Condition if not locked
}

// Withdraw subtracts money from the account
func (acc *BankAccount) Withdraw(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance -= amount
}

// Balance returns the current balance
func (acc *BankAccount) Balance() int {
	return acc.balance // Race Condition if not locked
}
